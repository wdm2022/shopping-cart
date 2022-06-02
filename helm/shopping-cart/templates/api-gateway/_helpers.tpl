{{/*
Expand the name of the component.
*/}}
{{- define "shopping-cart.apiGateway.fullname" -}}
{{- printf "%s-api-gateway" (include "shopping-cart.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "shopping-cart.apiGateway.selectorLabels" -}}
{{ include "shopping-cart.selectorLabels" . }}
shopping-cart/service: api-gateway
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "shopping-cart.apiGateway.serviceAccountName" -}}
{{- if .Values.apiGateway.serviceAccount.create }}
{{- default (include "shopping-cart.apiGateway.fullname" .) .Values.apiGateway.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.apiGateway.serviceAccount.name }}
{{- end }}
{{- end }}
