package requests

import (

)

// ObjectiveRequest represents the structure for Objective input (used for creating/updating)
type ObjectiveRequest struct {
	CompanyId string `json:"company_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	RiskAssessmentId string `json:"risk_assessment_id"`
	}

// ObjectiveResponse represents the structure for Objective output (used for returning data)
type ObjectiveResponse struct {
	ID       string `json:"id"`
	CompanyId string `json:"company_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	RiskAssessmentId string `json:"risk_assessment_id"`
	}

