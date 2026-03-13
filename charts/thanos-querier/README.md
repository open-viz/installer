# Thanos Querier

[Thanos Querier](https://github.com/go-openviz/installer) - Thanos Querier for Kubernetes

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/thanos-querier --version=v2026.3.13
$ helm upgrade -i thanos-querier appscode/thanos-querier -n openshift-monitoring --create-namespace --version=v2026.3.13
```

## Introduction

This chart deploys thanos querier on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `thanos-querier`:

```bash
$ helm upgrade -i thanos-querier appscode/thanos-querier -n openshift-monitoring --create-namespace --version=v2026.3.13
```

The command deploys thanos querier on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `thanos-querier`:

```bash
$ helm uninstall thanos-querier -n openshift-monitoring
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `thanos-querier` chart and their default values.

|                           Parameter                           |                                          Description                                           |                                                               Default                                                               |
|---------------------------------------------------------------|------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| nameOverride                                                  |                                                                                                | <code>""</code>                                                                                                                     |
| fullnameOverride                                              |                                                                                                | <code>""</code>                                                                                                                     |
| replicaCount                                                  |                                                                                                | <code>2</code>                                                                                                                      |
| annotations                                                   |                                                                                                | <code>{}</code>                                                                                                                     |
| podAnnotations                                                |                                                                                                | <code>{}</code>                                                                                                                     |
| imagePullSecrets                                              | Additional image pull secrets for pod spec.                                                    | <code>[]</code>                                                                                                                     |
| imagePullPolicy                                               | Container image pull policy used by all containers.                                            | <code>IfNotPresent</code>                                                                                                           |
| image.query.repository                                        |                                                                                                | <code>quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:b1d26a296f7030737a14e198a7c83a7d8720dad1d2b2ffaeaf7cbdbde274da78</code> |
| image.query.resources.requests.cpu                            |                                                                                                | <code>10m</code>                                                                                                                    |
| image.query.resources.requests.memory                         |                                                                                                | <code>12Mi</code>                                                                                                                   |
| image.query.securityContext.allowPrivilegeEscalation          |                                                                                                | <code>false</code>                                                                                                                  |
| image.query.securityContext.readOnlyRootFilesystem            |                                                                                                | <code>false</code>                                                                                                                  |
| image.query.securityContext.runAsNonRoot                      |                                                                                                | <code>true</code>                                                                                                                   |
| image.query.securityContext.seccompProfile.type               |                                                                                                | <code>RuntimeDefault</code>                                                                                                         |
| image.kubeRbacProxy.repository                                |                                                                                                | <code>quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:0e4654368f7dbc4e0808df96a3d1f85b3130185f5ba3523707980728d13fae43</code> |
| image.kubeRbacProxy.resources.requests.cpu                    |                                                                                                | <code>1m</code>                                                                                                                     |
| image.kubeRbacProxy.resources.requests.memory                 |                                                                                                | <code>15Mi</code>                                                                                                                   |
| image.kubeRbacProxy.securityContext.allowPrivilegeEscalation  |                                                                                                | <code>false</code>                                                                                                                  |
| image.promLabelProxy.repository                               |                                                                                                | <code>quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ab8dc6d8e30ef45c31d418718e17f1f5a500fa672a693058d61047a28a04adbd</code> |
| image.promLabelProxy.resources.requests.cpu                   |                                                                                                | <code>1m</code>                                                                                                                     |
| image.promLabelProxy.resources.requests.memory                |                                                                                                | <code>15Mi</code>                                                                                                                   |
| image.promLabelProxy.securityContext.allowPrivilegeEscalation |                                                                                                | <code>false</code>                                                                                                                  |
| nodeSelector.kubernetes.io/os                                 |                                                                                                | <code>linux</code>                                                                                                                  |
| tolerations                                                   |                                                                                                | <code>[]</code>                                                                                                                     |
| affinity                                                      | If empty, a default anti-affinity is applied in the deployment template.                       | <code>{}</code>                                                                                                                     |
| priorityClassName                                             |                                                                                                | <code>system-cluster-critical</code>                                                                                                |
| podSecurityContext.runAsNonRoot                               |                                                                                                | <code>true</code>                                                                                                                   |
| podSecurityContext.seccompProfile.type                        |                                                                                                | <code>RuntimeDefault</code>                                                                                                         |
| service.type                                                  |                                                                                                | <code>ClusterIP</code>                                                                                                              |
| serviceAccount.create                                         |                                                                                                | <code>true</code>                                                                                                                   |
| serviceAccount.annotations                                    |                                                                                                | <code>{}</code>                                                                                                                     |
| serviceAccount.name                                           |                                                                                                | <code></code>                                                                                                                       |
| secrets.create                                                | If false, secret manifests are not created and names below must refer to pre-existing secrets. | <code>true</code>                                                                                                                   |
| secrets.names.dockercfg                                       |                                                                                                | <code>""</code>                                                                                                                     |
| secrets.names.grpcTLS                                         |                                                                                                | <code>""</code>                                                                                                                     |
| secrets.names.kubeRbacProxy                                   |                                                                                                | <code>""</code>                                                                                                                     |
| secrets.names.kubeRbacProxyWeb                                |                                                                                                | <code>""</code>                                                                                                                     |
| secrets.names.kubeRbacProxyRules                              |                                                                                                | <code>""</code>                                                                                                                     |
| secrets.names.kubeRbacProxyMetrics                            |                                                                                                | <code>""</code>                                                                                                                     |
| secrets.names.tls                                             |                                                                                                | <code>""</code>                                                                                                                     |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i thanos-querier appscode/thanos-querier -n openshift-monitoring --create-namespace --version=v2026.3.13 --set replicaCount=2
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i thanos-querier appscode/thanos-querier -n openshift-monitoring --create-namespace --version=v2026.3.13 --values values.yaml
```
