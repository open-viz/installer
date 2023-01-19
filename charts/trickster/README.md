# Trickster

[Trickster](https://github.com/open-viz/trickster) - Trickster

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/trickster --version=v2022.06.14
$ helm upgrade -i trickster appscode/trickster -n kubeops --create-namespace --version=v2022.06.14
```

## Introduction

This chart deploys a Trickster Server on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `trickster`:

```bash
$ helm upgrade -i trickster appscode/trickster -n kubeops --create-namespace --version=v2022.06.14
```

The command deploys a Trickster Server on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `trickster`:

```bash
$ helm uninstall trickster -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `trickster` chart and their default values.

|                 Parameter                  |                                                                             Description                                                                             |                                                 Default                                                  |
|--------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------|
| replicaCount                               |                                                                                                                                                                     | <code>1</code>                                                                                           |
| registryFQDN                               | Docker registry fqdn used to pull app related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                | <code>ghcr.io</code>                                                                                     |
| image.registry                             | Docker registry used to pull app container image                                                                                                                    | <code>appscode</code>                                                                                    |
| image.repository                           | App container image                                                                                                                                                 | <code>trickster</code>                                                                                   |
| image.tag                                  | Overrides the image tag whose default is the chart appVersion.                                                                                                      | <code>""</code>                                                                                          |
| image.pullPolicy                           |                                                                                                                                                                     | <code>IfNotPresent</code>                                                                                |
| imagePullSecrets                           |                                                                                                                                                                     | <code>[]</code>                                                                                          |
| nameOverride                               |                                                                                                                                                                     | <code>""</code>                                                                                          |
| fullnameOverride                           |                                                                                                                                                                     | <code>""</code>                                                                                          |
| serviceAccount.create                      | Specifies whether a service account should be created                                                                                                               | <code>true</code>                                                                                        |
| serviceAccount.annotations                 | Annotations to add to the service account                                                                                                                           | <code>{}</code>                                                                                          |
| serviceAccount.name                        | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                              | <code></code>                                                                                            |
| podAnnotations                             |                                                                                                                                                                     | <code>{}</code>                                                                                          |
| podSecurityContext                         |                                                                                                                                                                     | <code>{}</code>                                                                                          |
| securityContext                            |                                                                                                                                                                     | <code>{}</code>                                                                                          |
| service.type                               |                                                                                                                                                                     | <code>ClusterIP</code>                                                                                   |
| service.port                               |                                                                                                                                                                     | <code>80</code>                                                                                          |
| ingress.enabled                            |                                                                                                                                                                     | <code>false</code>                                                                                       |
| ingress.className                          |                                                                                                                                                                     | <code>""</code>                                                                                          |
| ingress.annotations                        |                                                                                                                                                                     | <code>{}</code>                                                                                          |
| ingress.hosts                              | kubernetes.io/ingress.class: nginx kubernetes.io/tls-acme: "true"                                                                                                   | <code>[{"host":"chart-example.local","paths":[{"path":"/","pathType":"ImplementationSpecific"}]}]</code> |
| ingress.tls                                |                                                                                                                                                                     | <code>[]</code>                                                                                          |
| resources                                  |                                                                                                                                                                     | <code>{}</code>                                                                                          |
| autoscaling.enabled                        |                                                                                                                                                                     | <code>false</code>                                                                                       |
| autoscaling.minReplicas                    |                                                                                                                                                                     | <code>1</code>                                                                                           |
| autoscaling.maxReplicas                    |                                                                                                                                                                     | <code>100</code>                                                                                         |
| autoscaling.targetCPUUtilizationPercentage |                                                                                                                                                                     | <code>80</code>                                                                                          |
| nodeSelector                               |                                                                                                                                                                     | <code>{}</code>                                                                                          |
| tolerations                                |                                                                                                                                                                     | <code>[]</code>                                                                                          |
| affinity                                   |                                                                                                                                                                     | <code>{}</code>                                                                                          |
| monitoring.agent                           | Name of monitoring agent (eg "prometheus.io/operator")                                                                                                              | <code>""</code>                                                                                          |
| monitoring.port                            |                                                                                                                                                                     | <code>8080</code>                                                                                        |
| monitoring.serviceMonitor.labels           | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`. | <code>{}</code>                                                                                          |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i trickster appscode/trickster -n kubeops --create-namespace --version=v2022.06.14 --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i trickster appscode/trickster -n kubeops --create-namespace --version=v2022.06.14 --values values.yaml
```
