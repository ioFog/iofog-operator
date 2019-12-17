package kog

import (
	"context"
	"fmt"
	"strings"

	iofogv1 "github.com/eclipse-iofog/iofog-operator/pkg/apis/iofog/v1"
	k8sclient "github.com/eclipse-iofog/iofog-operator/pkg/controller/client"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *ReconcileKog) reconcileIofogConnectors(kog *iofogv1.Kog) error {

	// Find the current state to compare against requested state
	depList := &appsv1.DeploymentList{}
	if err := r.client.List(context.Background(), &client.ListOptions{}, depList); err != nil {
		return err
	}
	// Determine which connectors to create and delete
	createConnectors := make(map[string]bool)
	deleteConnectors := make(map[string]bool)
	for _, connector := range kog.Spec.Connectors.Instances {
		name := prefixConnectorName(connector.Name)
		createConnectors[name] = true
		deleteConnectors[name] = false
	}
	for _, dep := range depList.Items {
		if strings.Contains(dep.ObjectMeta.Name, getConnectorNamePrefix()) {
			createConnectors[dep.ObjectMeta.Name] = false
			if _, exists := deleteConnectors[dep.ObjectMeta.Name]; !exists {
				deleteConnectors[dep.ObjectMeta.Name] = true
			}
		}
	}

	// Delete connectors
	for k, isDelete := range deleteConnectors {
		if isDelete {
			if err := r.deleteConnector(kog, k); err != nil {
				return err
			}
		}
	}

	// Create connectors
	for k, isCreate := range createConnectors {
		if isCreate {
			if err := r.createConnector(kog, k); err != nil {
				return err
			}
		}
	}

	// Update existing Connector deployments (e.g. for image change)
	for k, isDelete := range deleteConnectors {
		// Untouched Connectors were neither deleted nor created
		if !isDelete {
			if isCreate, _ := createConnectors[k]; !isCreate {
				ms := newConnectorMicroservice(kog.Spec.Connectors.Image)
				ms.name = k
				// Deployment
				if err := r.createDeployment(kog, ms); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *ReconcileKog) reconcileIofogController(kog *iofogv1.Kog) error {
	cp := &kog.Spec.ControlPlane
	// Configure
	trafficPolicy := ""
	if cp.ServiceType == "NodePort" {
		trafficPolicy = "Cluster"
	} else if cp.ServiceType == "LoadBalancer" {
		trafficPolicy = "Local"
	}
	ms := newControllerMicroservice(
		cp.ControllerReplicaCount,
		cp.ControllerImage,
		cp.ImagePullSecret,
		&cp.Database,
		cp.ServiceType,
		trafficPolicy,
		cp.LoadBalancerIP,
	)
	r.apiEndpoint = fmt.Sprintf("%s:%d", ms.name, ms.ports[0])

	// Service Account
	if err := r.createServiceAccount(kog, ms); err != nil {
		return err
	}

	// Deployment
	if err := r.createDeployment(kog, ms); err != nil {
		return err
	}

	// Service
	if err := r.createService(kog, ms); err != nil {
		return err
	}

	// Connect to cluster
	k8sClient, err := k8sclient.NewClient()
	if err != nil {
		return err
	}

	// Wait for Pods
	if err = k8sClient.WaitForPod(kog.ObjectMeta.Namespace, ms.name, 120); err != nil {
		return err
	}

	// Wait for external IP of LB Service
	if cp.ServiceType == "LoadBalancer" {
		_, err = k8sClient.WaitForLoadBalancer(kog.ObjectMeta.Namespace, ms.name, 240)
		if err != nil {
			return err
		}
	}

	// Wait for Controller REST API
	if err = r.waitForControllerAPI(); err != nil {
		return err
	}

	// Set up user
	if err = r.createIofogUser(&cp.IofogUser); err != nil {
		return err
	}

	return nil
}

func (r *ReconcileKog) reconcileIofogKubelet(kog *iofogv1.Kog) error {
	// Generate new token if required
	token := ""
	kubeletKey := client.ObjectKey{
		Name:      "kubelet",
		Namespace: kog.ObjectMeta.Namespace,
	}
	dep := appsv1.Deployment{}
	if err := r.client.Get(context.TODO(), kubeletKey, &dep); err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		// Not found, generate new token
		token, err = r.getKubeletToken(&kog.Spec.ControlPlane.IofogUser)
		if err != nil {
			return err
		}
	} else {
		// Found, use existing token
		token, err = getKubeletToken(dep.Spec.Template.Spec.Containers)
		if err != nil {
			return err
		}
	}

	// Configure
	ms := newKubeletMicroservice(kog.Spec.ControlPlane.KubeletImage, kog.ObjectMeta.Namespace, token, r.apiEndpoint)

	// Service Account
	if err := r.createServiceAccount(kog, ms); err != nil {
		return err
	}
	// ClusterRoleBinding
	if err := r.createClusterRoleBinding(kog, ms); err != nil {
		return err
	}
	// Deployment
	if err := r.createDeployment(kog, ms); err != nil {
		return err
	}

	return nil
}
