package requests

import "time"

// LoginReq represents the structure for Login input (used for login)
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegistRequest represents the structure for User input (used for registering)
type RegistRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProfileResponse struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	VerifiedAt *time.Time    `json:"verified_at"`
	RoleID     *string       `json:"role_id"`
	Role       *RoleResponse `json:"role,omitempty"`
}
