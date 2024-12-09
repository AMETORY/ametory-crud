package models


import (
	"encoding/json"
	"gorm.io/gorm"
	"ametory-crud/requests"


)

type Objective struct {
	Base
	CompanyId string `gorm:"type:char(36);NOT NULL" json:"company_id"`
	Title string `gorm:"type:varchar(255);NOT NULL" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	RiskAssessmentId string `gorm:"type:char(36);NOT NULL" json:"risk_assessment_id"`
	}

func init() {
	RegisterModel(&Objective{})
}

func (p *Objective) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}

func (p Objective) MarshalJSON() ([]byte, error) {
	return json.Marshal(requests.ObjectiveResponse{
		ID:       p.ID,
		CompanyId: p.CompanyId,
		Title: p.Title,
		Description: p.Description,
		RiskAssessmentId: p.RiskAssessmentId,
		})
}

type ObjectiveResp struct {
	Pagination 	PaginationResponse 	`json:"pagination"`
	Data		[]Objective 	`json:"data"`
	Message 	string 				`json:"msg"`
}

type ObjectiveSingleResp struct {
	Data		Objective 	`json:"data"`
	Message 	string 				`json:"msg"`
}

func (p *Objective) UnmarshalJSON(data []byte) error {
	var req requests.ObjectiveRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	p.CompanyId = req.CompanyId
	p.Title = req.Title
	p.Description = req.Description
	p.RiskAssessmentId = req.RiskAssessmentId
	
	return nil
}
