package requests

import (

)

// RiskRequest represents the structure for Risk input (used for creating/updating)
type RiskRequest struct {
	CompanyId string `json:"company_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Category string `json:"category"`
	Status string `json:"status"`
	RiskOwnerId string `json:"risk_owner_id"`
	RiskTemplateId string `json:"risk_template_id"`
	}

// RiskResponse represents the structure for Risk output (used for returning data)
type RiskResponse struct {
	ID       string `json:"id"`
	CompanyId string `json:"company_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Category string `json:"category"`
	Status string `json:"status"`
	RiskOwnerId string `json:"risk_owner_id"`
	RiskTemplateId string `json:"risk_template_id"`
	}

