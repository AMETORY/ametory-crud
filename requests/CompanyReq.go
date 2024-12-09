package requests

import (

)

// CompanyRequest represents the structure for Company input (used for creating/updating)
type CompanyRequest struct {
	Name string `json:"name"`
	Industry string `json:"industry"`
	}

// CompanyResponse represents the structure for Company output (used for returning data)
type CompanyResponse struct {
	ID       string `json:"id"`
	Name string `json:"name"`
	Industry string `json:"industry"`
	}

