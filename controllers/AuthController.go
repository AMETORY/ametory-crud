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

// @Summary      Login
// @Description  login by username and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        LoginData body requests.LoginReq true "Login data"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /Auth/Login [post]
func LoginAuth(c *gin.Context) {
	var loginData models.Auth

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	// Example authentication logic
	user, err := models.FindUserByEmail(loginData.Email)
	if err != nil || !user.CheckPassword(loginData.Password) {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid username or password"})
		return
	}

	if user.VerifiedAt == nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Please verify your email address"})
		return
	}

	token, err := utils.GenerateJWTToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

// @Summary      Register
// @Description  register a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        RegisterData body requests.RegistRequest true "Register data"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /Auth/Registration [post]
func RegisterUser(c *gin.Context) {
	var registerData models.Auth

	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var newPassword = registerData.Password
	if newPassword == "" {
		newPassword = utils.GenerateRandomString(12)
	}

	existingUser, err := models.FindUserByEmail(registerData.Email)
	if err == nil && existingUser != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Email already in use"})
		return
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to hash password"})
		return
	}
	registerData.Password = hashedPassword

	auth, err := registerData.CreateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to register user"})
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

	c.JSON(http.StatusCreated, map[string]interface{}{"message": "User registered successfully"})
}

// @Summary      Profile
// @Description  get user profile
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  requests.ProfileResponse
// @Failure      401  {object}  map[string]interface{}
// @Router       /Auth/Profile [get]
// @Security BearerAuth
func Profile(c *gin.Context) {
	authData, exists := c.Get("auth")
	if !exists {
		c.JSON(http.StatusForbidden, map[string]interface{}{"error": "Authentication data not found"})
		c.Abort()
		return
	}

	auth := authData.(models.Auth)

	c.JSON(http.StatusOK, auth)
}

// @Summary      Verification
// @Description  verification by user id
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /Auth/Verification/{id} [get]
func Verification(c *gin.Context) {
	userID := c.Param("id")
	auth, err := models.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "User not found"})
		return
	}
	if auth.VerifiedAt != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "User already verified"})
		return
	}
	now := time.Now()
	auth.VerifiedAt = &now
	if err := auth.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "User verified successfully"})
}
