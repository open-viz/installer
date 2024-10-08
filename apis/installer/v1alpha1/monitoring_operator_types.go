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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindMonitoringOperator = "MonitoringOperator"
	ResourceMonitoringOperator     = "monitoringoperator"
	ResourceMonitoringOperators    = "monitoringoperators"
)

// MonitoringOperator defines the schama for ui server installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=monitoringoperators,singular=monitoringoperator,categories={kubeops,appscode}
type MonitoringOperator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringOperatorSpec `json:"spec,omitempty"`
}

// MonitoringOperatorSpec is the schema for Identity Server values file
type MonitoringOperatorSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string       `json:"fullnameOverride"`
	ReplicaCount     int32        `json:"replicaCount"`
	RegistryFQDN     string       `json:"registryFQDN"`
	Image            ContianerRef `json:"image"`
	ImagePullPolicy  string       `json:"imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	Apiserver          EASSpec                  `json:"apiserver"`
	Monitoring         EASMonitoring            `json:"monitoring"`
	Platform           PlatformSpec             `json:"platform"`
	HubUID             string                   `json:"hubUID"`
	Rancher            PlatformSpec             `json:"rancher"`
}

type PlatformSpec struct {
	BaseURL string `json:"baseURL"`
	Token   string `json:"token"`
	// +optional
	CABundle string `json:"caBundle"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MonitoringOperatorList is a list of MonitoringOperators
type MonitoringOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitoringOperator CRD objects
	Items []MonitoringOperator `json:"items,omitempty"`
}
