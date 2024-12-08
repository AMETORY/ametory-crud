package controllers

import (
	"ametory-crud/config"
	"ametory-crud/database"
	"ametory-crud/models"
	"ametory-crud/objects"
	"ametory-crud/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginAuth(c *gin.Context) {
	var loginData models.Auth

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Example authentication logic
	user, err := models.FindUserByEmail(loginData.Email)
	if err != nil || !user.CheckPassword(loginData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if user.VerifiedAt == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please verify your email address"})
		return
	}

	token, err := utils.GenerateJWTToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func RegisterUser(c *gin.Context) {
	var registerData models.Auth

	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newPassword = registerData.Password
	if newPassword == "" {
		newPassword = utils.GenerateRandomString(12)
	}

	existingUser, err := models.FindUserByEmail(registerData.Email)
	if err == nil && existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	registerData.Password = hashedPassword

	auth, err := registerData.CreateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	link := fmt.Sprintf("%s/verification/%s", config.App.Server.FrontEndURL, auth.ID)

	emailData, _ := json.Marshal(objects.UserReg{
		Name:     auth.Name,
		Link:     link,
		Password: newPassword,
		Email:    auth.Email,
	})
	fmt.Println(objects.QueueSendMail, string(emailData))
	database.REDIS.RPush(objects.QueueSendMail, string(emailData))

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Verification(c *gin.Context) {
	userID := c.Param("id")
	auth, err := models.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if auth.VerifiedAt != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already verified"})
		return
	}
	now := time.Now()
	auth.VerifiedAt = &now
	if err := auth.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User verified successfully"})
}
