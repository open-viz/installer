# Grafana Operator

[Grafana Operator by AppsCode](https://github.com/open-viz/grafana-operator) - Grafana Operator for Kubernetes

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/grafana-operator --version=v2024.11.18
$ helm upgrade -i grafana-operator appscode/grafana-operator -n kubeops --create-namespace --version=v2024.11.18
```

## Introduction

This chart deploys a Grafana operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `grafana-operator`:

```bash
$ helm upgrade -i grafana-operator appscode/grafana-operator -n kubeops --create-namespace --version=v2024.11.18
```

The command deploys a Grafana operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `grafana-operator`:

```bash
$ helm uninstall grafana-operator -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `grafana-operator` chart and their default values.

|              Parameter               |                                                                                                                  Description                                                                                                                  |                                                                                            Default                                                                                             |
|--------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| nameOverride                         | Overrides name template                                                                                                                                                                                                                       | <code>""</code>                                                                                                                                                                                |
| fullnameOverride                     | Overrides fullname template                                                                                                                                                                                                                   | <code>""</code>                                                                                                                                                                                |
| replicaCount                         | Number of Grafana operator replicas to create (only 1 is supported)                                                                                                                                                                           | <code>1</code>                                                                                                                                                                                 |
| registryFQDN                         | Docker registry fqdn used to pull KubeDB related images Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                        | <code>ghcr.io</code>                                                                                                                                                                           |
| image.registry                       | Docker registry used to pull operator image                                                                                                                                                                                                   | <code>appscode</code>                                                                                                                                                                          |
| image.repository                     | Name of operator container image                                                                                                                                                                                                              | <code>grafana-tools</code>                                                                                                                                                                     |
| image.tag                            | Operator container image tag                                                                                                                                                                                                                  | <code>""</code>                                                                                                                                                                                |
| image.resources                      | Compute Resources required by the operator container                                                                                                                                                                                          | <code>{}</code>                                                                                                                                                                                |
| image.securityContext                | Security options this container should run with                                                                                                                                                                                               | <code>{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"readOnlyRootFilesystem":true,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}</code> |
| imagePullSecrets                     | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/grafana-operator \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1` | <code>[]</code>                                                                                                                                                                                |
| imagePullPolicy                      | Container image pull policy                                                                                                                                                                                                                   | <code>IfNotPresent</code>                                                                                                                                                                      |
| criticalAddon                        | If true, installs Grafana operator as critical addon                                                                                                                                                                                          | <code>false</code>                                                                                                                                                                             |
| logLevel                             | Log level for operator                                                                                                                                                                                                                        | <code>3</code>                                                                                                                                                                                 |
| annotations                          | Annotations applied to operator deployment                                                                                                                                                                                                    | <code>{}</code>                                                                                                                                                                                |
| podAnnotations                       | Annotations passed to operator pod(s).                                                                                                                                                                                                        | <code>{}</code>                                                                                                                                                                                |
| nodeSelector                         | Node labels for pod assignment                                                                                                                                                                                                                | <code>{"kubernetes.io/os":"linux"}</code>                                                                                                                                                      |
| tolerations                          | Tolerations for pod assignment                                                                                                                                                                                                                | <code>[]</code>                                                                                                                                                                                |
| affinity                             | Affinity rules for pod assignment                                                                                                                                                                                                             | <code>{}</code>                                                                                                                                                                                |
| podSecurityContext                   | Security options the operator pod should run with.                                                                                                                                                                                            | <code>{}</code>                                                                                                                                                                                |
| serviceAccount.create                | Specifies whether a service account should be created                                                                                                                                                                                         | <code>true</code>                                                                                                                                                                              |
| serviceAccount.annotations           | Annotations to add to the service account                                                                                                                                                                                                     | <code>{}</code>                                                                                                                                                                                |
| serviceAccount.name                  | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                                        | <code></code>                                                                                                                                                                                  |
| apiserver.useKubeapiserverFqdnForAks | If true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522 (default true)                                                                                                                        | <code>true</code>                                                                                                                                                                              |
| apiserver.healthcheck.enabled        | healthcheck configures the readiness and liveliness probes for the operator pod.                                                                                                                                                              | <code>true</code>                                                                                                                                                                              |
| apiserver.healthcheck.probePort      | The port the probe endpoint binds to                                                                                                                                                                                                          | <code>8081</code>                                                                                                                                                                              |
| monitoring.bindPort                  | The port the metric endpoint binds to                                                                                                                                                                                                         | <code>8080</code>                                                                                                                                                                              |
| monitoring.agent                     | Name of monitoring agent (one of "prometheus.io", "prometheus.io/operator", "prometheus.io/builtin")                                                                                                                                          | <code>""</code>                                                                                                                                                                                |
| monitoring.serviceMonitor.labels     | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                           | <code>{}</code>                                                                                                                                                                                |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i grafana-operator appscode/grafana-operator -n kubeops --create-namespace --version=v2024.11.18 --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i grafana-operator appscode/grafana-operator -n kubeops --create-namespace --version=v2024.11.18 --values values.yaml
```
