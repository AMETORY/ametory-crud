package requests

import (

	"time"

)

// RiskAssessmentRequest represents the structure for RiskAssessment input (used for creating/updating)
type RiskAssessmentRequest struct {
	RiskId string `json:"risk_id"`
	Likelihood string `json:"likelihood"`
	Impact string `json:"impact"`
	RiskScore int `json:"risk_score"`
	AssessedBy string `json:"assessed_by"`
	AssessedAt time.Time `json:"assessed_at"`
	}

// RiskAssessmentResponse represents the structure for RiskAssessment output (used for returning data)
type RiskAssessmentResponse struct {
	ID       string `json:"id"`
	RiskId string `json:"risk_id"`
	Likelihood string `json:"likelihood"`
	Impact string `json:"impact"`
	RiskScore int `json:"risk_score"`
	AssessedBy string `json:"assessed_by"`
	AssessedAt time.Time `json:"assessed_at"`
	}

