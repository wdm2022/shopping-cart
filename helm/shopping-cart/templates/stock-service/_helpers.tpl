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
