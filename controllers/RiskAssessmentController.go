package controllers

import (
    db "ametory-crud/database"
	mdl "ametory-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// @Summary Get all RiskAssessment
// @Description Get all RiskAssessment
// @Tags RiskAssessment
// @Accept  json
// @Produce  json
// @Success 200 {object} mdl.RiskAssessmentResp
// @Router /RiskAssessment [get]
func GetRiskAssessment(c *gin.Context) {
    var risk_assessment []mdl.RiskAssessment

    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    search := c.DefaultQuery("search", "")

    offset := (page - 1) * limit

    // Get the total count of records
    var count int64
    if err := db.DB.Model(&risk_assessment).Count(&count).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Paginate
    if err := db.DB.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+search+"%").Find(&risk_assessment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, mdl.RiskAssessmentResp{
		Pagination: mdl.PaginationResponse{
			Total:  count,
			Limit:  limit,
			Page:  page,
		},
        Message: "success retrived risk_assessment data",
		Data: risk_assessment,
	})
}

// @Summary Get an RiskAssessment by ID
// @Description Get an RiskAssessment by ID
// @Tags RiskAssessment
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the RiskAssessment"
// @Success 200 {object} mdl.RiskAssessmentSingleResp
// @Router /RiskAssessment/{id} [get]
func GetOneRiskAssessment(c *gin.Context) {
    id := c.Params.ByName("id")
    var risk_assessment mdl.RiskAssessment

    // Find the record by ID
    if err := db.DB.Where("id = ?", id).First(&risk_assessment).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    // Return the found record as JSON
    c.JSON(http.StatusOK, mdl.RiskAssessmentSingleResp{
		Message: "success retrived risk_assessment data",
		Data:    risk_assessment,
	})
}

// @Summary Create an RiskAssessment
// @Description Create an RiskAssessment
// @Tags RiskAssessment
// @Accept  json
// @Produce  json
// @Param input body requests.RiskAssessmentRequest true "The RiskAssessment to create"
// @Success 201 {object}  mdl.GeneralResp
// @Router /RiskAssessment [post]
func CreateRiskAssessment(c *gin.Context) {
    var input mdl.RiskAssessment

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
		Message: "Created RiskAssessment successfully",
	})
}

// @Summary Update an RiskAssessment
// @Description Update an RiskAssessment
// @Tags RiskAssessment
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the RiskAssessment"
// @Param input body requests.RiskAssessmentRequest true "The RiskAssessment to update"
// @Success 200 {object}  mdl.GeneralResp
// @Router /RiskAssessment/{id} [put]
// @Security BearerAuth
func UpdateRiskAssessment(c *gin.Context) {
    id := c.Params.ByName("id")
    var input mdl.RiskAssessment

    // Bind JSON to the request struct
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Before update, make sure the data is exist
    var data mdl.RiskAssessment
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
    c.JSON(http.StatusOK, gin.H{"message": "Updated RiskAssessment successfully"})
}

// @Summary Delete an RiskAssessment
// @Description Delete an RiskAssessment
// @Tags RiskAssessment
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the RiskAssessment"
// @Success 200 {object}  mdl.GeneralResp
// @Router /RiskAssessment/{id} [delete]
func DeleteRiskAssessment(c *gin.Context) {
    id := c.Params.ByName("id")
    // Before delete, make sure the data exists
    var data mdl.RiskAssessment
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
    c.JSON(http.StatusOK, gin.H{"message": "Deleted RiskAssessment successfully"})
}

