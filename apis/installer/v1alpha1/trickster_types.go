/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

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
	ResourceKindTrickster = "Trickster"
	ResourceTrickster     = "trickster"
	ResourceTricksters    = "tricksters"
)

// Trickster defines the schama for Trickster Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=tricksters,singular=trickster,categories={kubeops,appscode}
type Trickster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TricksterSpec `json:"spec,omitempty"`
}

// TricksterSpec is the schema for Trickster Operator values file
type TricksterSpec struct {
	ReplicaCount int `json:"replicaCount"`
	//+optional
	RegistryFQDN string         `json:"registryFQDN"`
	Image        ImageReference `json:"image"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string             `json:"fullnameOverride"`
	ServiceAccount   ServiceAccountSpec `json:"serviceAccount"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	//+optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
	Service         AceServiceSpec        `json:"service"`
	//+optional
	Resources   core.ResourceRequirements `json:"resources"`
	Autoscaling AutoscalingSpec           `json:"autoscaling"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity   *core.Affinity   `json:"affinity"`
	Monitoring CustomMonitoring `json:"monitoring"`
	Ingress    AppIngress       `json:"ingress"`
}

type ImageReference struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	PullPolicy string `json:"pullPolicy"`
}

type AceServiceSpec struct {
	Type string `json:"type"`
	Port int    `json:"port"`
}

type AutoscalingSpec struct {
	Enabled     bool `json:"enabled"`
	MinReplicas int  `json:"minReplicas"`
	MaxReplicas int  `json:"maxReplicas"`
	// +optional
	TargetCPUUtilizationPercentage int `json:"targetCPUUtilizationPercentage,omitempty"`
	// +optional
	TargetMemoryUtilizationPercentage int `json:"targetMemoryUtilizationPercentage,omitempty"`
}

type CustomMonitoring struct {
	Agent          MonitoringAgent       `json:"agent"`
	Port           int                   `json:"port"`
	ServiceMonitor *ServiceMonitorLabels `json:"serviceMonitor"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TricksterList is a list of Tricksters
type TricksterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Trickster CRD objects
	Items []Trickster `json:"items,omitempty"`
}

type AppIngress struct {
	Enabled     bool              `json:"enabled"`
	ClassName   string            `json:"className"`
	Annotations map[string]string `json:"annotations"`
	Hosts       []IngressHost     `json:"hosts"`
	TLS         []IngressTLS      `json:"tls"`
}

type IngressHost struct {
	Host  string     `json:"host"`
	Paths []HostPath `json:"paths"`
}

type HostPath struct {
	Path     string `json:"path"`
	PathType string `json:"pathType"`
}

type IngressTLS struct {
	SecretName string   `json:"secretName"`
	Hosts      []string `json:"hosts"`
}
