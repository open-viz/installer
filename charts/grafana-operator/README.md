# Grafana Operator

[Grafana Operator by AppsCode](https://github.com/open-viz/grafana-operator) - Grafana Operator for Kubernetes

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/grafana-operator --version=v2022.03.20
$ helm upgrade -i grafana-operator appscode/grafana-operator -n kubeops --create-namespace --version=v2022.03.20
```

## Introduction

This chart deploys a Grafana operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `grafana-operator`:

```bash
$ helm upgrade -i grafana-operator appscode/grafana-operator -n kubeops --create-namespace --version=v2022.03.20
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

|               Parameter               |                                                                                                                                                                            Description                                                                                                                                                                             |                  Default                  |
|---------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------|
| nameOverride                          | Overrides name template                                                                                                                                                                                                                                                                                                                                            | <code>""</code>                           |
| fullnameOverride                      | Overrides fullname template                                                                                                                                                                                                                                                                                                                                        | <code>""</code>                           |
| replicaCount                          | Number of Grafana operator replicas to create (only 1 is supported)                                                                                                                                                                                                                                                                                                | <code>1</code>                            |
| operator.registry                     | Docker registry used to pull operator image                                                                                                                                                                                                                                                                                                                        | <code>appscode</code>                     |
| operator.repository                   | Name of operator container image                                                                                                                                                                                                                                                                                                                                   | <code>grafana-tools</code>                |
| operator.tag                          | Operator container image tag                                                                                                                                                                                                                                                                                                                                       | <code>v0.0.1</code>                       |
| operator.resources                    | Compute Resources required by the operator container                                                                                                                                                                                                                                                                                                               | <code>{}</code>                           |
| operator.securityContext              | Security options the operator container should run with                                                                                                                                                                                                                                                                                                            | <code>{}</code>                           |
| cleaner.registry                      | Docker registry used to pull Webhook cleaner image                                                                                                                                                                                                                                                                                                                 | <code>appscode</code>                     |
| cleaner.repository                    | Webhook cleaner container image                                                                                                                                                                                                                                                                                                                                    | <code>kubectl</code>                      |
| cleaner.tag                           | Webhook cleaner container image tag                                                                                                                                                                                                                                                                                                                                | <code>v1.22</code>                        |
| cleaner.skip                          | Skip generating cleaner YAML                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                        |
| imagePullSecrets                      | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/grafana-operator \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1`                                                                                                                      | <code>[]</code>                           |
| imagePullPolicy                       | Container image pull policy                                                                                                                                                                                                                                                                                                                                        | <code>IfNotPresent</code>                 |
| criticalAddon                         | If true, installs Grafana operator as critical addon                                                                                                                                                                                                                                                                                                               | <code>false</code>                        |
| logLevel                              | Log level for operator                                                                                                                                                                                                                                                                                                                                             | <code>3</code>                            |
| annotations                           | Annotations applied to operator deployment                                                                                                                                                                                                                                                                                                                         | <code>{}</code>                           |
| podAnnotations                        | Annotations passed to operator pod(s).                                                                                                                                                                                                                                                                                                                             | <code>{}</code>                           |
| nodeSelector                          | Node labels for pod assignment                                                                                                                                                                                                                                                                                                                                     | <code>{"kubernetes.io/os":"linux"}</code> |
| tolerations                           | Tolerations for pod assignment                                                                                                                                                                                                                                                                                                                                     | <code>[]</code>                           |
| affinity                              | Affinity rules for pod assignment                                                                                                                                                                                                                                                                                                                                  | <code>{}</code>                           |
| podSecurityContext                    | Security options the operator pod should run with.                                                                                                                                                                                                                                                                                                                 | <code>{"fsGroup":65535}</code>            |
| serviceAccount.create                 | Specifies whether a service account should be created                                                                                                                                                                                                                                                                                                              | <code>true</code>                         |
| serviceAccount.annotations            | Annotations to add to the service account                                                                                                                                                                                                                                                                                                                          | <code>{}</code>                           |
| serviceAccount.name                   | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                                                                                                                                                             | <code></code>                             |
| apiserver.groupPriorityMinimum        | The minimum priority the webhook api group should have at least. Please see https://github.com/kubernetes/kube-aggregator/blob/release-1.9/pkg/apis/apiregistration/v1beta1/types.go#L58-L64 for more information on proper values of this field.                                                                                                                  | <code>10000</code>                        |
| apiserver.versionPriority             | The ordering of the webhook api inside of the group. Please see https://github.com/kubernetes/kube-aggregator/blob/release-1.9/pkg/apis/apiregistration/v1beta1/types.go#L66-L70 for more information on proper values of this field                                                                                                                               | <code>15</code>                           |
| apiserver.enableMutatingWebhook       | If true, mutating webhook is configured for Grafana CRDss                                                                                                                                                                                                                                                                                                          | <code>false</code>                        |
| apiserver.enableValidatingWebhook     | If true, validating webhook is configured for Grafana CRDss                                                                                                                                                                                                                                                                                                        | <code>false</code>                        |
| apiserver.ca                          | CA certificate used by the Kubernetes api server. This field is automatically assigned by the operator.                                                                                                                                                                                                                                                            | <code>not-ca-cert</code>                  |
| apiserver.bypassValidatingWebhookXray | If true, bypasses checks that validating webhook is actually enabled in the Kubernetes cluster.                                                                                                                                                                                                                                                                    | <code>false</code>                        |
| apiserver.useKubeapiserverFqdnForAks  | If true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522 (default true)                                                                                                                                                                                                                                             | <code>true</code>                         |
| apiserver.healthcheck.enabled         | If true, enables the readiness and liveliness probes for the operator pod.                                                                                                                                                                                                                                                                                         | <code>false</code>                        |
| apiserver.servingCerts.generate       | If true, generates on install/upgrade the certs that allow the kube-apiserver (and potentially ServiceMonitor) to authenticate operators pods. Otherwise specify certs in `apiserver.servingCerts.{caCrt, serverCrt, serverKey}`. See also: [example terraform](https://github.com/searchlight/installer/blob/master/charts/grafana-operator/example-terraform.tf) | <code>true</code>                         |
| apiserver.servingCerts.caCrt          | CA certficate used by serving certificate of webhook server.                                                                                                                                                                                                                                                                                                       | <code>""</code>                           |
| apiserver.servingCerts.serverCrt      | Serving certficate used by webhook server.                                                                                                                                                                                                                                                                                                                         | <code>""</code>                           |
| apiserver.servingCerts.serverKey      | Private key for the serving certificate used by webhook server.                                                                                                                                                                                                                                                                                                    | <code>""</code>                           |
| monitoring.agent                      | Name of monitoring agent (either "prometheus.io/operator" or "prometheus.io/builtin")                                                                                                                                                                                                                                                                              | <code>"none"</code>                       |
| monitoring.serviceMonitor.labels      | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                                                                                                                                                | <code>{}</code>                           |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i grafana-operator appscode/grafana-operator -n kubeops --create-namespace --version=v2022.03.20 --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i grafana-operator appscode/grafana-operator -n kubeops --create-namespace --version=v2022.03.20 --values values.yaml
```
