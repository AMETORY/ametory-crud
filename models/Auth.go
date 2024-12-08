package models

import (
	"ametory-crud/config"
	db "ametory-crud/database"
	"ametory-crud/utils"
	"time"

	"gorm.io/gorm"
)

type Auth struct {
	Base
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	VerifiedAt *time.Time `json:"verified_at"`
	RoleID     *string    `json:"role_id"`
	Role       Role       `gorm:"foreignKey:RoleID" json:"role"`
}

func (Auth) TableName() string {
	return config.App.Database.AuthTable
}

func (p *Auth) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}

func (a *Auth) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(a.Password, password)
}

func (a *Auth) CreateUser() (*Auth, error) {
	result := db.DB.Create(&a)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}

func (a *Auth) Save() error {
	result := db.DB.Save(&a)
	return result.Error
}

func (a *Auth) GetPermissions() ([]string, error) {
	var role Role
	if err := db.DB.Preload("Permissions").First(&role, "id = ?", a.RoleID).Error; err != nil {
		return nil, err
	}

	if role.IsSuperAdmin {
		var permissions []Permission
		db.DB.Find(&permissions)
		var keys []string
		for _, perm := range permissions {
			keys = append(keys, perm.Key)
		}
		return keys, nil
	}

	var keys []string
	for _, perm := range role.Permissions {
		keys = append(keys, perm.Key)
	}
	return keys, nil
}