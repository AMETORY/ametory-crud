package models

type {{ .Feature }} struct {
	{{- range .Columns }}
	{{ . }}
	{{- end }}
}
