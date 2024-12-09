package requests

import (

)

// UserRequest represents the structure for User input (used for creating/updating)
type UserRequest struct {
	CompanyId string `json:"company_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	}

// UserResponse represents the structure for User output (used for returning data)
type UserResponse struct {
	ID       string `json:"id"`
	CompanyId string `json:"company_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	}

