package controllers

import (
	"ametory-crud/services"
	"ametory-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FileUpload(c *gin.Context) {
	path, err := services.UploadFile(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"path":    path,
		"url":     utils.GetFileUrl(path),
	})
}
