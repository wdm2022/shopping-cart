{{/*
Expand the name of the component.
*/}}
{{- define "shopping-cart.stockService.fullname" -}}
{{- printf "%s-stock-service" (include "shopping-cart.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "shopping-cart.stockService.selectorLabels" -}}
{{ include "shopping-cart.selectorLabels" . }}
shopping-cart/service: stock
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "shopping-cart.stockService.serviceAccountName" -}}
{{- if .Values.stockService.serviceAccount.create }}
{{- default (include "shopping-cart.stockService.fullname" .) .Values.stockService.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.stockService.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Address of the service endpoint.
*/}}
{{- define "shopping-cart.stockService.serviceAddress" -}}
{{- printf "%s:%v" (include "shopping-cart.stockService.fullname" .) .Values.stockService.service.port }}
{{- end }}

{{/*
ConfigMap name.
*/}}
{{- define "shopping-cart.stockService.configMap" -}}
{{- printf "%s-config" (include "shopping-cart.stockService.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Mongo hosts.
*/}}
{{- define "shopping-cart.stockService.mongoHosts" -}}
{{- $fullname := include "shopping-cart.fullname" . }}
{{- $mongoSvcName := printf "%s-mongodb-stock" $fullname }}
{{- range until (index .Values "mongodb-stock" "replicaCount" | int) }}
- {{ printf "%s-%v.%s-headless" $mongoSvcName . $mongoSvcName | quote }}
{{- end }}
{{- end }}
