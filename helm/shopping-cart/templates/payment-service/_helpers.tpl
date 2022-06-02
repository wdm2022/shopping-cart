{{/*
Expand the name of the component.
*/}}
{{- define "shopping-cart.paymentService.fullname" -}}
{{- printf "%s-payment-service" (include "shopping-cart.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "shopping-cart.paymentService.selectorLabels" -}}
{{ include "shopping-cart.selectorLabels" . }}
shopping-cart/service: payment
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "shopping-cart.paymentService.serviceAccountName" -}}
{{- if .Values.paymentService.serviceAccount.create }}
{{- default (include "shopping-cart.paymentService.fullname" .) .Values.paymentService.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.paymentService.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Address of the service endpoint.
*/}}
{{- define "shopping-cart.paymentService.serviceAddress" -}}
{{- printf "%s:%v" (include "shopping-cart.paymentService.fullname" .) .Values.paymentService.service.port }}
{{- end }}

{{/*
ConfigMap name.
*/}}
{{- define "shopping-cart.paymentService.configMap" -}}
{{- printf "%s-config" (include "shopping-cart.paymentService.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}
