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
	ResourceKindGrafanaOpscenter = "GrafanaOpscenter"
	ResourceGrafanaOpscenter     = "grafanaopscenter"
	ResourceGrafanaOpscenters    = "grafanaopscenters"
)

// GrafanaOpscenter defines the schama for KubeDB combined Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=grafanaopscenters,singular=grafanaopscenter,categories={grafanaopscenter,appscode}
type GrafanaOpscenter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GrafanaOpscenterSpec `json:"spec,omitempty"`
}

// GrafanaOpscenterSpec is the schema for grafanaopscenter chart values file
type GrafanaOpscenterSpec struct {
	Global GlobalValues `json:"global"`

	//+optional
	GrafanaOperator GrafanaOperatorValues `json:"grafana-operator"`

	//+optional
	GrafanaUiServer GrafanaUiServerValues `json:"grafana-ui-server"`

	//+optional
	KubeGrafanaDashboards KubeGrafanaDashboardsValues `json:"kube-grafana-dashboards"`
}

type GrafanaOperatorValues struct {
	Enabled              *bool `json:"enabled"`
	*GrafanaOperatorSpec `json:",inline,omitempty"`
}

type GrafanaUiServerValues struct {
	Enabled              *bool `json:"enabled"`
	*GrafanaUiServerSpec `json:",inline,omitempty"`
}

type KubeGrafanaDashboardsValues struct {
	Enabled                    bool `json:"enabled"`
	*KubeGrafanaDashboardsSpec `json:",inline,omitempty"`
}

type GlobalValues struct {
	Registry     string `json:"registry"`
	RegistryFQDN string `json:"registryFQDN"`
	//+optional
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets"`
	Monitoring       EASMonitoring               `json:"monitoring"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GrafanaOpscenterList is a list of GrafanaOpscenters
type GrafanaOpscenterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GrafanaOpscenter CRD objects
	Items []GrafanaOpscenter `json:"items,omitempty"`
}
