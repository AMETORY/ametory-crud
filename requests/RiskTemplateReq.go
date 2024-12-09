package requests

import (

)

// RiskTemplateRequest represents the structure for RiskTemplate input (used for creating/updating)
type RiskTemplateRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Category string `json:"category"`
	Status string `json:"status"`
	}

// RiskTemplateResponse represents the structure for RiskTemplate output (used for returning data)
type RiskTemplateResponse struct {
	ID       string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Category string `json:"category"`
	Status string `json:"status"`
	}

