package models


import (
	"encoding/json"
	"gorm.io/gorm"
	"ametory-crud/requests"


)

type Risk struct {
	Base
	CompanyId string `gorm:"type:char(36);NOT NULL" json:"company_id"`
	Title string `gorm:"type:varchar(255);NOT NULL" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Category string `gorm:"type:risk_category_enum;NOT NULL" json:"category"`
	Status string `gorm:"type:risk_status_enum ;DEFAULT Identified" json:"status"`
	RiskOwnerId string `gorm:"type:char(36);NOT NULL" json:"risk_owner_id"`
	RiskTemplateId string `gorm:"type:char(36);NOT NULL" json:"risk_template_id"`
	}

func init() {
	RegisterModel(&Risk{})
}

func (p *Risk) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}

func (p Risk) MarshalJSON() ([]byte, error) {
	return json.Marshal(requests.RiskResponse{
		ID:       p.ID,
		CompanyId: p.CompanyId,
		Title: p.Title,
		Description: p.Description,
		Category: p.Category,
		Status: p.Status,
		RiskOwnerId: p.RiskOwnerId,
		RiskTemplateId: p.RiskTemplateId,
		})
}

type RiskResp struct {
	Pagination 	PaginationResponse 	`json:"pagination"`
	Data		[]Risk 	`json:"data"`
	Message 	string 				`json:"msg"`
}

type RiskSingleResp struct {
	Data		Risk 	`json:"data"`
	Message 	string 				`json:"msg"`
}

func (p *Risk) UnmarshalJSON(data []byte) error {
	var req requests.RiskRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	p.CompanyId = req.CompanyId
	p.Title = req.Title
	p.Description = req.Description
	p.Category = req.Category
	p.Status = req.Status
	p.RiskOwnerId = req.RiskOwnerId
	p.RiskTemplateId = req.RiskTemplateId
	
	return nil
}
