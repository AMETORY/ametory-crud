package models


import (
	"encoding/json"
	"gorm.io/gorm"
	"ametory-crud/requests"
{{if .IsHasTime}}
	"time"
{{end}}

)

type {{ToPascalCase .ModelName}} struct {
	Base
	{{range .Fields}}{{ToPascalCase .Name}} {{.Type}} `gorm:"type:{{.DBType}}{{ .Default }}{{ .NotNull }}" json:"{{ToSnakeCase .Tag}}"`
	{{end}}}

func init() {
	RegisterModel(&{{ToPascalCase .ModelName}}{})
}

func (p *{{ToPascalCase .ModelName}}) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}

func (p {{ToPascalCase .ModelName}}) MarshalJSON() ([]byte, error) {
	return json.Marshal(requests.{{ToPascalCase .ModelName}}Response{
		ID:       p.ID,
		{{range .Fields}}{{ToPascalCase .Name}}: p.{{ToPascalCase .Name}},
		{{end}}})
}

type {{ToPascalCase .ModelName}}Resp struct {
	Pagination 	PaginationResponse 	`json:"pagination"`
	Data		[]{{ToPascalCase .ModelName}} 	`json:"data"`
	Message 	string 				`json:"msg"`
}

type {{ToPascalCase .ModelName}}SingleResp struct {
	Data		{{ToPascalCase .ModelName}} 	`json:"data"`
	Message 	string 				`json:"msg"`
}

func (p *{{ToPascalCase .ModelName}}) UnmarshalJSON(data []byte) error {
	var req requests.{{ToPascalCase .ModelName}}Request
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	{{range .Fields}}p.{{ToPascalCase .Name}} = req.{{ToPascalCase .Name}}
	{{end}}
	return nil
}
