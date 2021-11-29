# Grafana Configurator

[Grafana Configurator by AppsCode](https://github.com/open-viz/installer) - Grafana Configurator for ByteBuilders

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install grafana-configurator appscode/grafana-configurator -n kubeops
```

## Introduction

This chart deploys a Grafana Configurator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `grafana-configurator`:

```console
$ helm install grafana-configurator appscode/grafana-configurator -n kubeops
```

The command deploys a Grafana Configurator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `grafana-configurator`:

```console
$ helm delete grafana-configurator -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `grafana-configurator` chart and their default values.

|             Parameter             | Description | Default |
|-----------------------------------|-------------|---------|
| nameOverride                      |             | `""`    |
| fullnameOverride                  |             | `""`    |
| grafana.url                       |             | `""`    |
| grafana.service.scheme            |             | `""`    |
| grafana.service.name              |             | `""`    |
| grafana.service.port              |             | `""`    |
| grafana.service.path              |             | `""`    |
| grafana.service.query             |             | `""`    |
| grafana.auth.apiKey               |             | `""`    |
| grafana.tls.insecureSkipTLSVerify |             | `false` |
| grafana.tls.caBundle              |             | `""`    |
| dashboard.datasource              |             | `""`    |
| dashboard.folderID                |             | `0`     |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install grafana-configurator appscode/grafana-configurator -n kubeops --set dashboard.folderID=0
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install grafana-configurator appscode/grafana-configurator -n kubeops --values values.yaml
```
