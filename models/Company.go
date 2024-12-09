package models


import (
	"encoding/json"
	"gorm.io/gorm"
	"ametory-crud/requests"


)

type Company struct {
	Base
	Name string `gorm:"type:varchar(255) ;NOT NULL" json:"name"`
	Industry string `gorm:"type:varchar(100)" json:"industry"`
	}

func init() {
	RegisterModel(&Company{})
}

func (p *Company) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}

func (p Company) MarshalJSON() ([]byte, error) {
	return json.Marshal(requests.CompanyResponse{
		ID:       p.ID,
		Name: p.Name,
		Industry: p.Industry,
		})
}

type CompanyResp struct {
	Pagination 	PaginationResponse 	`json:"pagination"`
	Data		[]Company 	`json:"data"`
	Message 	string 				`json:"msg"`
}

type CompanySingleResp struct {
	Data		Company 	`json:"data"`
	Message 	string 				`json:"msg"`
}

func (p *Company) UnmarshalJSON(data []byte) error {
	var req requests.CompanyRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	p.Name = req.Name
	p.Industry = req.Industry
	
	return nil
}
