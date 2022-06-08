{{/*
Expand the name of the component.
*/}}
{{- define "shopping-cart.orderService.fullname" -}}
{{- printf "%s-order-service" (include "shopping-cart.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "shopping-cart.orderService.selectorLabels" -}}
{{ include "shopping-cart.selectorLabels" . }}
shopping-cart/service: order
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "shopping-cart.orderService.serviceAccountName" -}}
{{- if .Values.orderService.serviceAccount.create }}
{{- default (include "shopping-cart.orderService.fullname" .) .Values.orderService.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.orderService.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Address of the service endpoint.
*/}}
{{- define "shopping-cart.orderService.serviceAddress" -}}
{{- printf "%s:%v" (include "shopping-cart.orderService.fullname" .) .Values.orderService.service.port }}
{{- end }}

{{/*
ConfigMap name.
*/}}
{{- define "shopping-cart.orderService.configMap" -}}
{{- printf "%s-config" (include "shopping-cart.orderService.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Mongo hosts.
*/}}
{{- define "shopping-cart.orderService.mongoHosts" -}}
{{- $fullname := include "shopping-cart.fullname" . }}
{{- $mongoSvcName := printf "%s-mongodb-order" $fullname }}
{{- range until (index .Values "mongodb-order" "replicaCount" | int) }}
- {{ printf "%s-%v.%s-headless" $mongoSvcName . $mongoSvcName | quote }}
{{- end }}
{{- end }}
