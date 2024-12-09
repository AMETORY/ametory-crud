package requests

import (

	"time"

)

// MitigationRequest represents the structure for Mitigation input (used for creating/updating)
type MitigationRequest struct {
	ActionPlan string `json:"action_plan"`
	AssignedTo string `json:"assigned_to"`
	Status string `json:"status"`
	DueDate time.Time `json:"due_date"`
	RiskId string `json:"risk_id"`
	RiskAssessmentId string `json:"risk_assessment_id"`
	}

// MitigationResponse represents the structure for Mitigation output (used for returning data)
type MitigationResponse struct {
	ID       string `json:"id"`
	ActionPlan string `json:"action_plan"`
	AssignedTo string `json:"assigned_to"`
	Status string `json:"status"`
	DueDate time.Time `json:"due_date"`
	RiskId string `json:"risk_id"`
	RiskAssessmentId string `json:"risk_assessment_id"`
	}

