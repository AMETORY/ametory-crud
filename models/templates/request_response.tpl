package requests

import (
{{if .IsHasTime}}
	"time"
{{end}}
)

// {{ToPascalCase .ModelName}}Request represents the structure for {{ToPascalCase .ModelName}} input (used for creating/updating)
type {{ToPascalCase .ModelName}}Request struct {
	{{range .Fields}}{{ToPascalCase .Name}} {{.Type}} `json:"{{ToSnakeCase .Tag }}"`
	{{end}}}

// {{ToPascalCase .ModelName}}Response represents the structure for {{ToPascalCase .ModelName}} output (used for returning data)
type {{ToPascalCase .ModelName}}Response struct {
	ID       string `json:"id"`
	{{range .Fields}}{{ToPascalCase .Name}} {{.Type}} `json:"{{ToSnakeCase .Tag }}"`
	{{end}}}

