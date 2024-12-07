package models

import (
	"ametory-crud/config"
	"ametory-crud/utils"
)

type Auth struct {
	Base
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (Auth) TableName() string {
	return config.App.Database.AuthTable
}

func (a *Auth) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(a.Password, password)
}
