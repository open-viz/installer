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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindKubeGrafanaDashboards = "KubeGrafanaDashboards"
	ResourceKubeGrafanaDashboards     = "kubegrafanadashboards"
	ResourceKubeGrafanaDashboardss    = "kubegrafanadashboardss"
)

// KubeGrafanaDashboards defines the schama for ui server installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubegrafanadashboardss,singular=kubegrafanadashboards,categories={kubeops,appscode}
type KubeGrafanaDashboards struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeGrafanaDashboardsSpec `json:"spec,omitempty"`
}

// KubeGrafanaDashboardsSpec is the schema for Identity Server values file
type KubeGrafanaDashboardsSpec struct {
	NameOverride          string                         `json:"nameOverride"`
	FullnameOverride      string                         `json:"fullnameOverride"`
	Dashboard             DashboardSpec                  `json:"dashboard"`
	Grafana               GrafanaDashboard               `json:"grafana"`
	CoreDNS               CoreDNSDashboard               `json:"coreDns"`
	KubeEtcd              KubeEtcdDashboard              `json:"kubeEtcd"`
	KubeAPIServer         KubeAPIServerDashboard         `json:"kubeApiServer"`
	KubeControllerManager KubeControllerManagerDashboard `json:"kubeControllerManager"`
	Kubelet               KubeletDashboard               `json:"kubelet"`
	KubeProxy             KubeProxyDashboard             `json:"kubeProxy"`
	KubeScheduler         KubeSchedulerDashboard         `json:"kubeScheduler"`
	NodeExporter          NodeExporterDashboard          `json:"nodeExporter"`
	Prometheus            PrometheusDashboard            `json:"prometheus"`
}

type DashboardSpec struct {
	FolderID     int                   `json:"folderID"`
	Overwrite    bool                  `json:"overwrite"`
	Templatize   TemplatizeDashboard   `json:"templatize"`
	Multicluster MulticlusterDashboard `json:"multicluster"`
}

type TemplatizeDashboard struct {
	Title      bool `json:"title"`
	Datasource bool `json:"datasource"`
}

type MulticlusterDashboard struct {
	Global MulticlusterGlobalDashboard `json:"global"`
	Etcd   MulticlusterEtcdDashboard   `json:"etcd"`
}

type MulticlusterGlobalDashboard struct {
	Enabled bool `json:"enabled"`
}

type MulticlusterEtcdDashboard struct {
	Enabled bool `json:"enabled"`
}

type GrafanaDashboard struct {
	Name                      string `json:"name"`
	Namespace                 string `json:"namespace"`
	DefaultDashboardsTimezone string `json:"defaultDashboardsTimezone"`
}

type CoreDNSDashboard struct {
	Enabled bool `json:"enabled"`
}

type KubeEtcdDashboard struct {
	Enabled bool `json:"enabled"`
}

type KubeAPIServerDashboard struct {
	Enabled bool `json:"enabled"`
}

type KubeControllerManagerDashboard struct {
	Enabled bool `json:"enabled"`
}

type KubeletDashboard struct {
	Enabled bool `json:"enabled"`
}

type KubeProxyDashboard struct {
	Enabled bool `json:"enabled"`
}

type KubeSchedulerDashboard struct {
	Enabled bool `json:"enabled"`
}

type NodeExporterDashboard struct {
	Enabled bool `json:"enabled"`
}

type PrometheusDashboard struct {
	RemoteWriteDashboards bool `json:"remoteWriteDashboards"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeGrafanaDashboardsList is a list of KubeGrafanaDashboardss
type KubeGrafanaDashboardsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubeGrafanaDashboards CRD objects
	Items []KubeGrafanaDashboards `json:"items,omitempty"`
}
