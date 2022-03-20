# Kubernetes Grafana Dashboards

[Kubernetes Grafana Dashboards by AppsCode](https://github.com/open-viz/installer) - Grafana dashboards for end-to-end Kubernetes cluster monitoring

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kube-grafana-dashboards --version=v2022.03.20
$ helm upgrade -i kube-grafana-dashboards appscode/kube-grafana-dashboards -n kubeops --create-namespace --version=v2022.03.20
```

## Introduction

This chart deploys Grafana dashboards on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `kube-grafana-dashboards`:

```bash
$ helm upgrade -i kube-grafana-dashboards appscode/kube-grafana-dashboards -n kubeops --create-namespace --version=v2022.03.20
```

The command deploys Grafana dashboards on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kube-grafana-dashboards`:

```bash
$ helm uninstall kube-grafana-dashboards -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kube-grafana-dashboards` chart and their default values.

|               Parameter               |                                                  Description                                                  |      Default       |
|---------------------------------------|---------------------------------------------------------------------------------------------------------------|--------------------|
| nameOverride                          | Overrides name template                                                                                       | <code>""</code>    |
| fullnameOverride                      | Overrides fullname template                                                                                   | <code>""</code>    |
| dashboard.folderID                    | ID of Grafana folder where these dashboards will be applied                                                   | <code>0</code>     |
| dashboard.overwrite                   | If true, dashboard with matching uid will be overwritten                                                      | <code>true</code>  |
| dashboard.templatize.title            | If true, datasource will be prefixed to dashboard name                                                        | <code>false</code> |
| dashboard.templatize.datasource       | If true, datasource will be hardcoded in the dashboard                                                        | <code>true</code>  |
| dashboard.multicluster.global.enabled |                                                                                                               | <code>false</code> |
| dashboard.multicluster.etcd.enabled   |                                                                                                               | <code>false</code> |
| grafana.name                          | Name of Grafana Appbinding where these dashboards are applied                                                 | <code>""</code>    |
| grafana.namespace                     | Namespace of Grafana Appbinding where these dashboards are applied                                            | <code>""</code>    |
| grafana.defaultDashboardsTimezone     | Timezone for the default dashboards Other options are: browser or a specific timezone, i.e. Europe/Luxembourg | <code>utc</code>   |
| coreDns.enabled                       |                                                                                                               | <code>true</code>  |
| kubeEtcd.enabled                      |                                                                                                               | <code>true</code>  |
| kubeApiServer.enabled                 |                                                                                                               | <code>true</code>  |
| kubeControllerManager.enabled         |                                                                                                               | <code>true</code>  |
| kubelet.enabled                       |                                                                                                               | <code>true</code>  |
| kubeProxy.enabled                     |                                                                                                               | <code>true</code>  |
| kubeScheduler.enabled                 |                                                                                                               | <code>true</code>  |
| nodeExporter.enabled                  |                                                                                                               | <code>true</code>  |
| prometheus.remoteWriteDashboards      | Enable/Disable Grafana dashboards provisioning for prometheus remote write feature                            | <code>false</code> |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kube-grafana-dashboards appscode/kube-grafana-dashboards -n kubeops --create-namespace --version=v2022.03.20 --set dashboard.folderID=0
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kube-grafana-dashboards appscode/kube-grafana-dashboards -n kubeops --create-namespace --version=v2022.03.20 --values values.yaml
```
