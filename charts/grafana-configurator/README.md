# Grafana Configurator

[Grafana Configurator by AppsCode](https://github.com/open-viz/installer) - Grafana Configurator for ByteBuilders

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/grafana-configurator --version=v2021.11.16
$ helm upgrade -i grafana-configurator appscode/grafana-configurator -n kubeops --create-namespace --version=v2021.11.16
```

## Introduction

This chart deploys a Grafana Configurator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `grafana-configurator`:

```bash
$ helm upgrade -i grafana-configurator appscode/grafana-configurator -n kubeops --create-namespace --version=v2021.11.16
```

The command deploys a Grafana Configurator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `grafana-configurator`:

```bash
$ helm uninstall grafana-configurator -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `grafana-configurator` chart and their default values.

|             Parameter             | Description |      Default       |
|-----------------------------------|-------------|--------------------|
| nameOverride                      |             | <code>""</code>    |
| fullnameOverride                  |             | <code>""</code>    |
| grafana.default                   |             | <code>false</code> |
| grafana.url                       |             | <code>""</code>    |
| grafana.service.scheme            |             | <code>""</code>    |
| grafana.service.name              |             | <code>""</code>    |
| grafana.service.port              |             | <code>""</code>    |
| grafana.service.path              |             | <code>""</code>    |
| grafana.service.query             |             | <code>""</code>    |
| grafana.auth.apiKey               |             | <code>""</code>    |
| grafana.tls.insecureSkipTLSVerify |             | <code>false</code> |
| grafana.tls.caBundle              |             | <code>""</code>    |
| dashboard.datasource              |             | <code>""</code>    |
| dashboard.folderID                |             | <code>0</code>     |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i grafana-configurator appscode/grafana-configurator -n kubeops --create-namespace --version=v2021.11.16 --set dashboard.folderID=0
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i grafana-configurator appscode/grafana-configurator -n kubeops --create-namespace --version=v2021.11.16 --values values.yaml
```
