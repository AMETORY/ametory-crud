package models


import (
	"encoding/json"
	"gorm.io/gorm"
	"ametory-crud/requests"
)

type Product struct {
	Base
	Name string `gorm:"type:varchar(255)" json:"name"`
	Price float64 `gorm:"type:decimal(10,2)" json:"price"`
	}

func init() {
	RegisterModel(&Product{})
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(requests.ProductResponse{
		ID:       p.ID,
		Name: p.Name,
		Price: p.Price,
		})
}

type ProductResp struct {
	Pagination 	PaginationResponse 	`json:"pagination"`
	Data		[]Product 	`json:"data"`
	Message 	string 				`json:"msg"`
}

type ProductSingleResp struct {
	Data		Product 	`json:"data"`
	Message 	string 				`json:"msg"`
}

func (p *Product) UnmarshalJSON(data []byte) error {
	var req requests.ProductRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	p.Name = req.Name
	p.Price = req.Price
	
	return nil
}
