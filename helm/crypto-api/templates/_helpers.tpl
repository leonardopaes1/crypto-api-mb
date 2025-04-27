{{/* 
================================================================================
  _helpers.tpl – Funções utilitárias reutilizadas nos demais templates
  Adapte os sufixos/prefixos se quiser mudar o estilo dos nomes dos recursos.
================================================================================
*/}}

{{/* Retorna o nome da aplicação (valor .Values.appName ou .Chart.Name) */}}
{{- define "chart.appName" -}}
{{- default .Chart.Name .Values.appName | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/* Retorna um nome “completo” único para o recurso (ex.: crypto-api-staging) */}}
{{- define "chart.fullname" -}}
{{- printf "%s-%s" (include "chart.appName" .) .Values.environment | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/* Rótulos comuns para todos os recursos */}}
{{- define "chart.labels" -}}
app.kubernetes.io/name: {{ include "chart.appName" . }}
app.kubernetes.io/instance: {{ include "chart.fullname" . }}
app.kubernetes.io/version: {{ .Chart.AppVersion }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Nome base da aplicação, reutilizado em todo o chart
*/}}
{{- define "app.name" -}}
{{ .Values.appName | default .Chart.Name }}
{{- end }}

{{/*
Rótulos comuns
*/}}
{{- define "app.labels" -}}
app.kubernetes.io/name: {{ include "app.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion }}
{{- end }}