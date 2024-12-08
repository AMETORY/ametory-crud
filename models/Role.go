package models

import (
	"ametory-crud/requests"
	"encoding/json"

	"gorm.io/gorm"
)

type Role struct {
	Base
	Name         string       `gorm:"type:varchar(255)" json:"name"`
	Description  string       `gorm:"type:text" json:"description"`
	Permissions  []Permission `gorm:"many2many:role_permissions;"`
	IsSuperAdmin bool         `json:"isSuperAdmin"`
}

func (p *Role) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}

func (p Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(requests.RoleResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	})
}

type RoleResp struct {
	Pagination PaginationResponse `json:"pagination"`
	Data       []Role             `json:"data"`
	Message    string             `json:"msg"`
}

type RoleSingleResp struct {
	Data    Role   `json:"data"`
	Message string `json:"msg"`
}

func (p *Role) UnmarshalJSON(data []byte) error {
	var req requests.RoleRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	p.Name = req.Name
	p.Description = req.Description
	return nil
}

type Permission struct {
	Base
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Key         string `gorm:"type:varchar(255);unique" json:"key"`
	Group       string `gorm:"type:varchar(255)" json:"group"`
}

func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}
