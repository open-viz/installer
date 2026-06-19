/*
Copyright AppsCode Inc. and Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	core "k8s.io/api/core/v1"
)

// ThanosQuerierSpec is the schema for thanos-querier values file.
type ThanosQuerierSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`
	ReplicaCount     int32  `json:"replicaCount"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets"`
	ImagePullPolicy  string                      `json:"imagePullPolicy"`
	Image            ThanosQuerierImageSpec      `json:"image"`
	ThanosQuery      ThanosQueryFlagsSpec        `json:"thanosQuery"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints.
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	//+optional
	PriorityClassName string `json:"priorityClassName"`
	//+optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	Service            ThanosQuerierServiceSpec `json:"service"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	Secrets            ThanosQuerierSecretsSpec `json:"secrets"`
}

type ThanosQuerierImageSpec struct {
	Query          ThanosQuerierContainerSpec `json:"query"`
	KubeRbacProxy  ThanosQuerierContainerSpec `json:"kubeRbacProxy"`
	PromLabelProxy ThanosQuerierContainerSpec `json:"promLabelProxy"`
}

type ThanosQuerierContainerSpec struct {
	Repository string `json:"repository"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
	//+optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
}

type ThanosQueryFlagsSpec struct {
	ReplicaLabels []string `json:"replicaLabels"`
	EndpointFlags []string `json:"endpointFlags"`
	StoreFlags    []string `json:"storeFlags"`
	RuleFlags     []string `json:"ruleFlags"`
	TargetFlags   []string `json:"targetFlags"`
}

type ThanosQuerierServiceSpec struct {
	Type string `json:"type"`
}

type ThanosQuerierSecretsSpec struct {
	Create bool                         `json:"create"`
	Names  ThanosQuerierSecretNamesSpec `json:"names"`
}

type ThanosQuerierSecretNamesSpec struct {
	Dockercfg            string `json:"dockercfg"`
	GrpcTLS              string `json:"grpcTLS"`
	KubeRbacProxy        string `json:"kubeRbacProxy"`
	KubeRbacProxyWeb     string `json:"kubeRbacProxyWeb"`
	KubeRbacProxyRules   string `json:"kubeRbacProxyRules"`
	KubeRbacProxyMetrics string `json:"kubeRbacProxyMetrics"`
	TLS                  string `json:"tls"`
}
