package apis

import (
	extsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	appsv3 "github.com/eclipse-iofog/iofog-operator/v3/apis/apps/v3"
	cpv3 "github.com/eclipse-iofog/iofog-operator/v3/apis/controlplanes/v3"
)

func NewControlPlaneCustomResource() *extsv1.CustomResourceDefinition {
	apiVersions := []string{"v3", "v2"}
	versions := make([]extsv1.CustomResourceDefinitionVersion, len(apiVersions))
	for idx, version := range apiVersions {
		versions[idx].Name = version
		versions[idx].Served = true
		if idx == 0 {
			versions[idx].Storage = true
		}
	}
	return &extsv1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "controlplanes.iofog.org",
		},
		Spec: extsv1.CustomResourceDefinitionSpec{
			Group: "iofog.org",
			Names: extsv1.CustomResourceDefinitionNames{
				Kind:     "ControlPlane",
				ListKind: "ControlPlaneList",
				Plural:   "controlplanes",
				Singular: "controlplane",
			},
			Scope:    extsv1.NamespaceScoped,
			Versions: versions,
			Subresources: &extsv1.CustomResourceSubresources{
				Status: &extsv1.CustomResourceSubresourceStatus{},
			},
		},
	}
}

func NewAppCustomResource() *extsv1.CustomResourceDefinition {
	apiVersions := []string{"v3", "v2", "v1"}
	versions := make([]extsv1.CustomResourceDefinitionVersion, len(apiVersions))
	for idx, version := range apiVersions {
		versions[idx].Name = version
		versions[idx].Served = true
		if idx == 0 {
			versions[idx].Storage = true
		}
	}
	return &extsv1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "apps.iofog.org",
		},
		Spec: extsv1.CustomResourceDefinitionSpec{
			Group: "iofog.org",
			Names: extsv1.CustomResourceDefinitionNames{
				Kind:     "Application",
				ListKind: "ApplicationList",
				Plural:   "apps",
				Singular: "app",
			},
			Scope:    extsv1.NamespaceScoped,
			Versions: versions,
			Subresources: &extsv1.CustomResourceSubresources{
				Status: &extsv1.CustomResourceSubresourceStatus{},
			},
		},
	}
}

func sameVersionsSupported(left, right *extsv1.CustomResourceDefinition) bool {
	for _, leftVersion := range left.Spec.Versions {
		matched := false
		for _, rightVersion := range right.Spec.Versions {
			if leftVersion.Name == rightVersion.Name {
				matched = true
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

func IsSupportedCustomResource(crd *extsv1.CustomResourceDefinition) bool {
	cpCR := NewControlPlaneCustomResource()
	if crd.Name == cpCR.Name {
		return sameVersionsSupported(cpCR, crd)
	}
	appCR := NewAppCustomResource()
	if crd.Name == appCR.Name {
		return sameVersionsSupported(appCR, crd)
	}
	return false
}

func InitClientScheme() *runtime.Scheme {
	scheme := runtime.NewScheme()
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(appsv3.AddToScheme(scheme))
	utilruntime.Must(cpv3.AddToScheme(scheme))
	return scheme
}
