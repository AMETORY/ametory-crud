package requests

import (

	"time"

)

// MitigationActionRequest represents the structure for MitigationAction input (used for creating/updating)
type MitigationActionRequest struct {
	MitigationId string `json:"mitigation_id"`
	Description string `json:"description"`
	AssignedTo string `json:"assigned_to"`
	Status string `json:"status"`
	DueDate time.Time `json:"due_date"`
	}

// MitigationActionResponse represents the structure for MitigationAction output (used for returning data)
type MitigationActionResponse struct {
	ID       string `json:"id"`
	MitigationId string `json:"mitigation_id"`
	Description string `json:"description"`
	AssignedTo string `json:"assigned_to"`
	Status string `json:"status"`
	DueDate time.Time `json:"due_date"`
	}

