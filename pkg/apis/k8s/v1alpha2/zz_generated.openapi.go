// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha2

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/eclipse-iofog/iofog-operator/pkg/apis/k8s/v1alpha2.Kog":       schema_pkg_apis_k8s_v1alpha2_Kog(ref),
		"github.com/eclipse-iofog/iofog-operator/pkg/apis/k8s/v1alpha2.KogSpec":   schema_pkg_apis_k8s_v1alpha2_KogSpec(ref),
		"github.com/eclipse-iofog/iofog-operator/pkg/apis/k8s/v1alpha2.KogStatus": schema_pkg_apis_k8s_v1alpha2_KogStatus(ref),
	}
}

func schema_pkg_apis_k8s_v1alpha2_Kog(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Kog is the Schema for the kogs API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/eclipse-iofog/iofog-operator/pkg/apis/k8s/v1alpha2.KogSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/eclipse-iofog/iofog-operator/pkg/apis/k8s/v1alpha2.KogStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/eclipse-iofog/iofog-operator/pkg/apis/k8s/v1alpha2.KogSpec", "github.com/eclipse-iofog/iofog-operator/pkg/apis/k8s/v1alpha2.KogStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_k8s_v1alpha2_KogSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogSpec defines the desired state of Kog",
				Properties: map[string]spec.Schema{
					"controlPlane": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("github.com/eclipse-iofog/iofog-operator/pkg/apis/k8s/v1alpha2.ControlPlane"),
						},
					},
					"connectors": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/eclipse-iofog/iofog-operator/pkg/apis/k8s/v1alpha2.Connectors"),
						},
					},
				},
				Required: []string{"controlPlane", "connectors"},
			},
		},
		Dependencies: []string{
			"github.com/eclipse-iofog/iofog-operator/pkg/apis/k8s/v1alpha2.Connectors", "github.com/eclipse-iofog/iofog-operator/pkg/apis/k8s/v1alpha2.ControlPlane"},
	}
}

func schema_pkg_apis_k8s_v1alpha2_KogStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogStatus defines the observed state of Kog",
				Properties: map[string]spec.Schema{
					"controllerPods": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"controllerPods"},
			},
		},
		Dependencies: []string{},
	}
}
