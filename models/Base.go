package models

import (
	"time"
)

type Base struct {
	ID        string     `gorm:"type:char(36);primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}
type PaginationResponse struct {
	Total  int64 `json:"total"`
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
}

type GeneralResp struct {
	Message string `json:"msg"`
}
