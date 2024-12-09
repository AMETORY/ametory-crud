package utils

import (
	"ametory-crud/config"
	"fmt"
	"math/rand"
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

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func GetFileUrl(filename string) string {
	if config.App.Server.StorageProvider == "google" {
		return fmt.Sprintf("https://storage.googleapis.com/%s/%s", config.App.Google.FirebaseStorageBucket, filename)
	}
	if config.App.Server.StorageProvider == "s3" {
		return fmt.Sprintf("%s/%s", config.App.S3.PublicURL, filename)
	}
	return fmt.Sprintf("%s/%s", config.App.Server.ApiURL, filename)
}
