package models


import (
	"encoding/json"
	"gorm.io/gorm"
	"ametory-crud/requests"
)

type {{.ModelName}} struct {
	Base
	{{range .Fields}}{{.Name}} {{.Type}} `gorm:"type:{{.DBType}}" json:"{{.Tag}}"`
	{{end}}}

func init() {
	RegisterModel(&{{.ModelName}}{})
}

func (p *{{.ModelName}}) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}

func (p {{.ModelName}}) MarshalJSON() ([]byte, error) {
	return json.Marshal(requests.{{.ModelName}}Response{
		ID:       p.ID,
		{{range .Fields}}{{.Name}}: p.{{.Name}},
		{{end}}})
}

type {{.ModelName}}Resp struct {
	Pagination 	PaginationResponse 	`json:"pagination"`
	Data		[]{{.ModelName}} 	`json:"data"`
	Message 	string 				`json:"msg"`
}

type {{.ModelName}}SingleResp struct {
	Data		{{.ModelName}} 	`json:"data"`
	Message 	string 				`json:"msg"`
}

func (p *{{.ModelName}}) UnmarshalJSON(data []byte) error {
	var req requests.{{.ModelName}}Request
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	{{range .Fields}}p.{{.Name}} = req.{{.Name}}
	{{end}}
	return nil
}
