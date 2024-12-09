package controllers

import (
	"ametory-crud/services"
	"ametory-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      File Upload
// @Description  Upload a file and return its path and URL
// @Tags         file
// @Accept       multipart/form-data
// @Produce      json
// @Param        file formData file true "File to upload"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /FileUpload [post]
// @Security BearerAuth
func FileUpload(c *gin.Context) {
	path, err := services.UploadFile(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "File uploaded successfully",
		"path":    path,
		"url":     utils.GetFileUrl(path),
	})
}
