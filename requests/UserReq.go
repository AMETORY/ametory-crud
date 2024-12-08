package requests

// UserRequest represents the structure for User input (used for creating/updating)
type UserRequest struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
	Password string `json:"password"`
	}

// UserResponse represents the structure for User output (used for returning data)
type UserResponse struct {
	ID       string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
	Password string `json:"password"`
	}

