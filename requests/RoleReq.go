package requests // RoleRequest represents the structure for Role input (used for creating/updating)
type RoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// RoleResponse represents the structure for Role output (used for returning data)
type RoleResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}
