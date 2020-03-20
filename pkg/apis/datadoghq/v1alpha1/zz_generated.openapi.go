// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/CharlyF/schedulerenhancer/pkg/apis/datadoghq/v1alpha1.SchedulerEnhancer":       schema_pkg_apis_datadoghq_v1alpha1_SchedulerEnhancer(ref),
		"github.com/CharlyF/schedulerenhancer/pkg/apis/datadoghq/v1alpha1.SchedulerEnhancerSpec":   schema_pkg_apis_datadoghq_v1alpha1_SchedulerEnhancerSpec(ref),
		"github.com/CharlyF/schedulerenhancer/pkg/apis/datadoghq/v1alpha1.SchedulerEnhancerStatus": schema_pkg_apis_datadoghq_v1alpha1_SchedulerEnhancerStatus(ref),
	}
}

func schema_pkg_apis_datadoghq_v1alpha1_SchedulerEnhancer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SchedulerEnhancer is the Schema for the schedulerenhancers API",
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
							Ref: ref("github.com/CharlyF/schedulerenhancer/pkg/apis/datadoghq/v1alpha1.SchedulerEnhancerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/CharlyF/schedulerenhancer/pkg/apis/datadoghq/v1alpha1.SchedulerEnhancerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/CharlyF/schedulerenhancer/pkg/apis/datadoghq/v1alpha1.SchedulerEnhancerSpec", "github.com/CharlyF/schedulerenhancer/pkg/apis/datadoghq/v1alpha1.SchedulerEnhancerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_datadoghq_v1alpha1_SchedulerEnhancerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SchedulerEnhancerSpec defines the desired state of SchedulerEnhancer",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_datadoghq_v1alpha1_SchedulerEnhancerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SchedulerEnhancerStatus defines the observed state of SchedulerEnhancer",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
