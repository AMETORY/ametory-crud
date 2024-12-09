package controllers

import (
    db "ametory-crud/database"
	mdl "ametory-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// @Summary Get all Risk
// @Description Get all Risk
// @Tags Risk
// @Accept  json
// @Produce  json
// @Success 200 {object} mdl.RiskResp
// @Router /Risk [get]
func GetRisk(c *gin.Context) {
    var risk []mdl.Risk

    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    search := c.DefaultQuery("search", "")

    offset := (page - 1) * limit

    // Get the total count of records
    var count int64
    if err := db.DB.Model(&risk).Count(&count).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Paginate
    if err := db.DB.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+search+"%").Find(&risk).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, mdl.RiskResp{
		Pagination: mdl.PaginationResponse{
			Total:  count,
			Limit:  limit,
			Page:  page,
		},
        Message: "success retrived risk data",
		Data: risk,
	})
}

// @Summary Get an Risk by ID
// @Description Get an Risk by ID
// @Tags Risk
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the Risk"
// @Success 200 {object} mdl.RiskSingleResp
// @Router /Risk/{id} [get]
func GetOneRisk(c *gin.Context) {
    id := c.Params.ByName("id")
    var risk mdl.Risk

    // Find the record by ID
    if err := db.DB.Where("id = ?", id).First(&risk).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    // Return the found record as JSON
    c.JSON(http.StatusOK, mdl.RiskSingleResp{
		Message: "success retrived risk data",
		Data:    risk,
	})
}

// @Summary Create an Risk
// @Description Create an Risk
// @Tags Risk
// @Accept  json
// @Produce  json
// @Param input body requests.RiskRequest true "The Risk to create"
// @Success 201 {object}  mdl.GeneralResp
// @Router /Risk [post]
func CreateRisk(c *gin.Context) {
    var input mdl.Risk

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
		Message: "Created Risk successfully",
	})
}

// @Summary Update an Risk
// @Description Update an Risk
// @Tags Risk
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the Risk"
// @Param input body requests.RiskRequest true "The Risk to update"
// @Success 200 {object}  mdl.GeneralResp
// @Router /Risk/{id} [put]
// @Security BearerAuth
func UpdateRisk(c *gin.Context) {
    id := c.Params.ByName("id")
    var input mdl.Risk

    // Bind JSON to the request struct
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Before update, make sure the data is exist
    var data mdl.Risk
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
    c.JSON(http.StatusOK, gin.H{"message": "Updated Risk successfully"})
}

// @Summary Delete an Risk
// @Description Delete an Risk
// @Tags Risk
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the Risk"
// @Success 200 {object}  mdl.GeneralResp
// @Router /Risk/{id} [delete]
func DeleteRisk(c *gin.Context) {
    id := c.Params.ByName("id")
    // Before delete, make sure the data exists
    var data mdl.Risk
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
    c.JSON(http.StatusOK, gin.H{"message": "Deleted Risk successfully"})
}

