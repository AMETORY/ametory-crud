package models


import (
	"encoding/json"
	"gorm.io/gorm"
	"ametory-crud/requests"

	"time"


)

type MitigationAction struct {
	Base
	MitigationId string `gorm:"type:char(36);NOT NULL" json:"mitigation_id"`
	Description string `gorm:"type:text" json:"description"`
	AssignedTo string `gorm:"type:char(36);NOT NULL" json:"assigned_to"`
	Status string `gorm:"type:mitigation_action_status_enum ;DEFAULT Pending" json:"status"`
	DueDate time.Time `gorm:"type:date;NOT NULL" json:"due_date"`
	}

func init() {
	RegisterModel(&MitigationAction{})
}

func (p *MitigationAction) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}

func (p MitigationAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(requests.MitigationActionResponse{
		ID:       p.ID,
		MitigationId: p.MitigationId,
		Description: p.Description,
		AssignedTo: p.AssignedTo,
		Status: p.Status,
		DueDate: p.DueDate,
		})
}

type MitigationActionResp struct {
	Pagination 	PaginationResponse 	`json:"pagination"`
	Data		[]MitigationAction 	`json:"data"`
	Message 	string 				`json:"msg"`
}

type MitigationActionSingleResp struct {
	Data		MitigationAction 	`json:"data"`
	Message 	string 				`json:"msg"`
}

func (p *MitigationAction) UnmarshalJSON(data []byte) error {
	var req requests.MitigationActionRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	p.MitigationId = req.MitigationId
	p.Description = req.Description
	p.AssignedTo = req.AssignedTo
	p.Status = req.Status
	p.DueDate = req.DueDate
	
	return nil
}
