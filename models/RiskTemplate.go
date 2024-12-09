package models


import (
	"encoding/json"
	"gorm.io/gorm"
	"ametory-crud/requests"


)

type RiskTemplate struct {
	Base
	Title string `gorm:"type:varchar(255);NOT NULL" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Category string `gorm:"type:risk_template_category_enum;NOT NULL" json:"category"`
	Status string `gorm:"type:risk_template_status_enum ;DEFAULT Identified" json:"status"`
	}

func init() {
	RegisterModel(&RiskTemplate{})
}

func (p *RiskTemplate) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}

func (p RiskTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(requests.RiskTemplateResponse{
		ID:       p.ID,
		Title: p.Title,
		Description: p.Description,
		Category: p.Category,
		Status: p.Status,
		})
}

type RiskTemplateResp struct {
	Pagination 	PaginationResponse 	`json:"pagination"`
	Data		[]RiskTemplate 	`json:"data"`
	Message 	string 				`json:"msg"`
}

type RiskTemplateSingleResp struct {
	Data		RiskTemplate 	`json:"data"`
	Message 	string 				`json:"msg"`
}

func (p *RiskTemplate) UnmarshalJSON(data []byte) error {
	var req requests.RiskTemplateRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	p.Title = req.Title
	p.Description = req.Description
	p.Category = req.Category
	p.Status = req.Status
	
	return nil
}
