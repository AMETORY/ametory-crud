package utils

import (
	"ametory-crud/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWTToken(sub string) (string, error) {
	secretKey := config.App.Server.SecretKey
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = sub
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
