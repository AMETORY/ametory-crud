package requests

// ProductRequest represents the structure for Product input (used for creating/updating)
type ProductRequest struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
	}

// ProductResponse represents the structure for Product output (used for returning data)
type ProductResponse struct {
	ID       string `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	}

