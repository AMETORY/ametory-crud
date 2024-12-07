package models

import (
	db "ametory-crud/database"
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        string     `gorm:"type:char(36);primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}
type PaginationResponse struct {
	Total int64 `json:"total"`
	Limit int   `json:"limit"`
	Page  int   `json:"page"`
}

type GeneralResp struct {
	Message string `json:"msg"`
}

func FindUserByEmail(email string) (*Auth, error) {
	var user Auth
	result := db.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func generateUUID() string {
	return uuid.NewString()
}
