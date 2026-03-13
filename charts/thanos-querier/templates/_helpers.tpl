{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "thanos-querier.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
*/}}
{{- define "thanos-querier.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{- define "thanos-querier.instance" -}}
{{- .Release.Name -}}
{{- end }}

{{- define "thanos-querier.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "thanos-querier.selectorLabels" -}}
app.kubernetes.io/component: query-layer
app.kubernetes.io/instance: {{ include "thanos-querier.instance" . }}
app.kubernetes.io/name: thanos-query
app.kubernetes.io/part-of: openshift-monitoring
{{- end }}

{{- define "thanos-querier.labels" -}}
helm.sh/chart: {{ include "thanos-querier.chart" . }}
{{ include "thanos-querier.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{- define "thanos-querier.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "thanos-querier.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "thanos-querier.imagePullSecrets" -}}
{{- with .Values.imagePullSecrets -}}
imagePullSecrets:
{{- toYaml . | nindent 2 }}
{{- end -}}
{{- end }}

{{- define "thanos-querier.secretDockercfg" -}}
{{- if .Values.secrets.create -}}
{{- printf "%s-dockercfg" (include "thanos-querier.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- required "values.secrets.names.dockercfg is required when secrets.create=false" .Values.secrets.names.dockercfg -}}
{{- end -}}
{{- end }}

{{- define "thanos-querier.secretGrpcTLS" -}}
{{- if .Values.secrets.create -}}
{{- printf "%s-grpc-tls" (include "thanos-querier.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- required "values.secrets.names.grpcTLS is required when secrets.create=false" .Values.secrets.names.grpcTLS -}}
{{- end -}}
{{- end }}

{{- define "thanos-querier.secretKubeRbacProxy" -}}
{{- if .Values.secrets.create -}}
{{- printf "%s-kube-rbac-proxy" (include "thanos-querier.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- required "values.secrets.names.kubeRbacProxy is required when secrets.create=false" .Values.secrets.names.kubeRbacProxy -}}
{{- end -}}
{{- end }}

{{- define "thanos-querier.secretKubeRbacProxyWeb" -}}
{{- if .Values.secrets.create -}}
{{- printf "%s-kube-rbac-proxy-web" (include "thanos-querier.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- required "values.secrets.names.kubeRbacProxyWeb is required when secrets.create=false" .Values.secrets.names.kubeRbacProxyWeb -}}
{{- end -}}
{{- end }}

{{- define "thanos-querier.secretKubeRbacProxyRules" -}}
{{- if .Values.secrets.create -}}
{{- printf "%s-kube-rbac-proxy-rules" (include "thanos-querier.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- required "values.secrets.names.kubeRbacProxyRules is required when secrets.create=false" .Values.secrets.names.kubeRbacProxyRules -}}
{{- end -}}
{{- end }}

{{- define "thanos-querier.secretKubeRbacProxyMetrics" -}}
{{- if .Values.secrets.create -}}
{{- printf "%s-kube-rbac-proxy-metrics" (include "thanos-querier.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- required "values.secrets.names.kubeRbacProxyMetrics is required when secrets.create=false" .Values.secrets.names.kubeRbacProxyMetrics -}}
{{- end -}}
{{- end }}

{{- define "thanos-querier.secretTLS" -}}
{{- if .Values.secrets.create -}}
{{- printf "%s-tls" (include "thanos-querier.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- required "values.secrets.names.tls is required when secrets.create=false" .Values.secrets.names.tls -}}
{{- end -}}
{{- end }}

{{- define "thanos-querier.secretPrometheusSidecarTLS" -}}
{{- printf "%s-prometheus-k8s-thanos-sidecar-tls" (include "thanos-querier.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- end }}
