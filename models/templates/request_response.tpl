package requests

// {{.ModelName}}Request represents the structure for {{.ModelName}} input (used for creating/updating)
type {{.ModelName}}Request struct {
	{{range .Fields}}{{.Name}} {{.Type}} `json:"{{.Name | ToLower}}"`
	{{end}}}

// {{.ModelName}}Response represents the structure for {{.ModelName}} output (used for returning data)
type {{.ModelName}}Response struct {
	{{range .Fields}}{{.Name}} {{.Type}} `json:"{{.Name | ToLower}}"`
	{{end}}}

