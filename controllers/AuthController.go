package controllers

import (
	"ametory-crud/models"
	"ametory-crud/utils"
	"net/http"

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

	token, err := utils.GenerateJWTToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
