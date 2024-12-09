package controllers

import (
    db "ametory-crud/database"
	mdl "ametory-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// @Summary Get all RiskTemplate
// @Description Get all RiskTemplate
// @Tags RiskTemplate
// @Accept  json
// @Produce  json
// @Success 200 {object} mdl.RiskTemplateResp
// @Router /RiskTemplate [get]
func GetRiskTemplate(c *gin.Context) {
    var risk_template []mdl.RiskTemplate

    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    search := c.DefaultQuery("search", "")

    offset := (page - 1) * limit

    // Get the total count of records
    var count int64
    if err := db.DB.Model(&risk_template).Count(&count).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Paginate
    if err := db.DB.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+search+"%").Find(&risk_template).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, mdl.RiskTemplateResp{
		Pagination: mdl.PaginationResponse{
			Total:  count,
			Limit:  limit,
			Page:  page,
		},
        Message: "success retrived risk_template data",
		Data: risk_template,
	})
}

// @Summary Get an RiskTemplate by ID
// @Description Get an RiskTemplate by ID
// @Tags RiskTemplate
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the RiskTemplate"
// @Success 200 {object} mdl.RiskTemplateSingleResp
// @Router /RiskTemplate/{id} [get]
func GetOneRiskTemplate(c *gin.Context) {
    id := c.Params.ByName("id")
    var risk_template mdl.RiskTemplate

    // Find the record by ID
    if err := db.DB.Where("id = ?", id).First(&risk_template).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    // Return the found record as JSON
    c.JSON(http.StatusOK, mdl.RiskTemplateSingleResp{
		Message: "success retrived risk_template data",
		Data:    risk_template,
	})
}

// @Summary Create an RiskTemplate
// @Description Create an RiskTemplate
// @Tags RiskTemplate
// @Accept  json
// @Produce  json
// @Param input body requests.RiskTemplateRequest true "The RiskTemplate to create"
// @Success 201 {object}  mdl.GeneralResp
// @Router /RiskTemplate [post]
func CreateRiskTemplate(c *gin.Context) {
    var input mdl.RiskTemplate

    // Bind JSON to the request struct
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save to DB (example using GORM)
    if err := db.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return the response as JSON
    c.JSON(http.StatusCreated, mdl.GeneralResp{
		Message: "Created RiskTemplate successfully",
	})
}

// @Summary Update an RiskTemplate
// @Description Update an RiskTemplate
// @Tags RiskTemplate
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the RiskTemplate"
// @Param input body requests.RiskTemplateRequest true "The RiskTemplate to update"
// @Success 200 {object}  mdl.GeneralResp
// @Router /RiskTemplate/{id} [put]
// @Security BearerAuth
func UpdateRiskTemplate(c *gin.Context) {
    id := c.Params.ByName("id")
    var input mdl.RiskTemplate

    // Bind JSON to the request struct
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Before update, make sure the data is exist
    var data mdl.RiskTemplate
    if err := db.DB.Where("id = ?", id).First(&data).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    // Update to DB (example using GORM)
    if err := db.DB.Model(&input).Where("id = ?", id).Updates(input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return the response as JSON
    c.JSON(http.StatusOK, gin.H{"message": "Updated RiskTemplate successfully"})
}

// @Summary Delete an RiskTemplate
// @Description Delete an RiskTemplate
// @Tags RiskTemplate
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the RiskTemplate"
// @Success 200 {object}  mdl.GeneralResp
// @Router /RiskTemplate/{id} [delete]
func DeleteRiskTemplate(c *gin.Context) {
    id := c.Params.ByName("id")
    // Before delete, make sure the data exists
    var data mdl.RiskTemplate
    if err := db.DB.Where("id = ?", id).First(&data).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    // Delete from DB (example using GORM)
    if err := db.DB.Delete(&data).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return the response as JSON
    c.JSON(http.StatusOK, gin.H{"message": "Deleted RiskTemplate successfully"})
}

