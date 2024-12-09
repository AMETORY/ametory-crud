package controllers

import (
	"net/http"
	"strconv"

	db "ametory-crud/database"
	"ametory-crud/models"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	var roles []models.Role

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	search := c.DefaultQuery("search", "")

	offset := (page - 1) * limit

	var count int64
	if err := db.DB.Model(&roles).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+search+"%").Find(&roles).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.RoleResp{
		Pagination: models.PaginationResponse{
			Total: count,
			Limit: limit,
			Page:  page,
		},
		Data:    roles,
		Message: "success retrieved roles data",
	})
}

func GetOneRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role

	if err := db.DB.Where("id = ?", id).First(&role).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, models.RoleSingleResp{
		Message: "success retrieved role data",
		Data:    role,
	})
}

func CreateRole(c *gin.Context) {
	var input models.Role

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permissions := []models.Permission{}
	for _, p := range input.Permissions {
		var permission models.Permission
		db.DB.Where("key = ?", p.Key).First(&permission)
		permissions = append(permissions, permission)
	}

	if err := db.DB.Omit("Permissions").Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(permissions) > 0 {
		db.DB.Model(&input).Association("Permissions").Append(permissions)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created role successfully"})
}

func UpdateRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var input models.Role

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var role models.Role
	if err := db.DB.Where("id = ?", id).First(&role).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}

	if err := db.DB.Model(&role).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated role successfully"})
}

func DeleteRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role

	if err := db.DB.Where("id = ?", id).First(&role).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}

	if err := db.DB.Delete(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted role successfully"})
}
