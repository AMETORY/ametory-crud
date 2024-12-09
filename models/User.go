package models


import (
	"encoding/json"
	"gorm.io/gorm"
	"ametory-crud/requests"


)

type User struct {
	Base
	CompanyId string `gorm:"type:char(36);NOT NULL" json:"company_id"`
	Name string `gorm:"type:varchar(255);NOT NULL" json:"name"`
	Email string `gorm:"type:varchar(255) ;NOT NULL" json:"email"`
	Role string `gorm:"type:user_role_enum;NOT NULL" json:"role"`
	}

func init() {
	RegisterModel(&User{})
}

func (p *User) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}

func (p User) MarshalJSON() ([]byte, error) {
	return json.Marshal(requests.UserResponse{
		ID:       p.ID,
		CompanyId: p.CompanyId,
		Name: p.Name,
		Email: p.Email,
		Role: p.Role,
		})
}

type UserResp struct {
	Pagination 	PaginationResponse 	`json:"pagination"`
	Data		[]User 	`json:"data"`
	Message 	string 				`json:"msg"`
}

type UserSingleResp struct {
	Data		User 	`json:"data"`
	Message 	string 				`json:"msg"`
}

func (p *User) UnmarshalJSON(data []byte) error {
	var req requests.UserRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	p.CompanyId = req.CompanyId
	p.Name = req.Name
	p.Email = req.Email
	p.Role = req.Role
	
	return nil
}
