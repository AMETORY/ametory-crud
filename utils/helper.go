package utils

import (
	"ametory-crud/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWTToken(sub string) (string, error) {
	secretKey := config.App.Server.SecretKey
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = sub
	claims["exp"] = time.Now().Add(time.Hour * 24 * time.Duration(config.App.Server.ExpiredJWT)).Unix()
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

func CheckPasswordHash(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
