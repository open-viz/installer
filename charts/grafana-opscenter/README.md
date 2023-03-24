# Grafana Opscenter

[Grafana Opscenter by AppsCode](https://github.com/kubedb) - Grafana Opscenter

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/grafana-opscenter --version=v2023.03.23
$ helm upgrade -i grafana-opscenter appscode/grafana-opscenter -n kubeops --create-namespace --version=v2023.03.23
```

## Introduction

This chart deploys a Grafana Opscenter on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `grafana-opscenter`:

```bash
$ helm upgrade -i grafana-opscenter appscode/grafana-opscenter -n kubeops --create-namespace --version=v2023.03.23
```

The command deploys a Grafana Opscenter on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `grafana-opscenter`:

```bash
$ helm uninstall grafana-opscenter -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `grafana-opscenter` chart and their default values.

|                Parameter                |                                                                                                                         Description                                                                                                                          |      Default      |
|-----------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------|
| global.registry                         | Docker registry used to pull KubeDB related images                                                                                                                                                                                                           | <code>""</code>   |
| global.registryFQDN                     | Docker registry fqdn used to pull KubeDB related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                      | <code>""</code>   |
| global.imagePullSecrets                 | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/grafana-opscenter \` <br> `--set global.imagePullSecrets[0].name=sec0 \` <br> `--set global.imagePullSecrets[1].name=sec1` | <code>[]</code>   |
| global.monitoring.agent                 | Name of monitoring agent (one of "prometheus.io", "prometheus.io/operator", "prometheus.io/builtin")                                                                                                                                                         | <code>""</code>   |
| global.monitoring.serviceMonitor.labels | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                                          | <code>{}</code>   |
| grafana-operator.enabled                | If enabled, installs the grafana-operator chart                                                                                                                                                                                                              | <code>true</code> |
| grafana-ui-server.enabled               | If enabled, installs the grafana-ui-server chart                                                                                                                                                                                                             | <code>true</code> |
| kube-grafana-dashboards.enabled         | If enabled, installs the kube-grafana-dashboards chart                                                                                                                                                                                                       | <code>true</code> |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i grafana-opscenter appscode/grafana-opscenter -n kubeops --create-namespace --version=v2023.03.23 --set -- generate from values file --
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i grafana-opscenter appscode/grafana-opscenter -n kubeops --create-namespace --version=v2023.03.23 --values values.yaml
```
