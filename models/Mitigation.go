package models


import (
	"encoding/json"
	"gorm.io/gorm"
	"ametory-crud/requests"

	"time"


)

type Mitigation struct {
	Base
	ActionPlan string `gorm:"type:text;NOT NULL" json:"action_plan"`
	AssignedTo string `gorm:"type:char(36);NOT NULL" json:"assigned_to"`
	Status string `gorm:"type:mitigation_status_enum ;DEFAULT Pending" json:"status"`
	DueDate time.Time `gorm:"type:date;NOT NULL" json:"due_date"`
	RiskId string `gorm:"type:char(36);NOT NULL" json:"risk_id"`
	RiskAssessmentId string `gorm:"type:char(36);NOT NULL" json:"risk_assessment_id"`
	}

func init() {
	RegisterModel(&Mitigation{})
}

func (p *Mitigation) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}

func (p Mitigation) MarshalJSON() ([]byte, error) {
	return json.Marshal(requests.MitigationResponse{
		ID:       p.ID,
		ActionPlan: p.ActionPlan,
		AssignedTo: p.AssignedTo,
		Status: p.Status,
		DueDate: p.DueDate,
		RiskId: p.RiskId,
		RiskAssessmentId: p.RiskAssessmentId,
		})
}

type MitigationResp struct {
	Pagination 	PaginationResponse 	`json:"pagination"`
	Data		[]Mitigation 	`json:"data"`
	Message 	string 				`json:"msg"`
}

type MitigationSingleResp struct {
	Data		Mitigation 	`json:"data"`
	Message 	string 				`json:"msg"`
}

func (p *Mitigation) UnmarshalJSON(data []byte) error {
	var req requests.MitigationRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	p.ActionPlan = req.ActionPlan
	p.AssignedTo = req.AssignedTo
	p.Status = req.Status
	p.DueDate = req.DueDate
	p.RiskId = req.RiskId
	p.RiskAssessmentId = req.RiskAssessmentId
	
	return nil
}
