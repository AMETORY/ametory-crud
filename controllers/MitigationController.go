package controllers

import (
    db "ametory-crud/database"
	mdl "ametory-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// @Summary Get all Mitigation
// @Description Get all Mitigation
// @Tags Mitigation
// @Accept  json
// @Produce  json
// @Success 200 {object} mdl.MitigationResp
// @Router /Mitigation [get]
func GetMitigation(c *gin.Context) {
    var mitigation []mdl.Mitigation

    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    search := c.DefaultQuery("search", "")

    offset := (page - 1) * limit

    // Get the total count of records
    var count int64
    if err := db.DB.Model(&mitigation).Count(&count).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Paginate
    if err := db.DB.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+search+"%").Find(&mitigation).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, mdl.MitigationResp{
		Pagination: mdl.PaginationResponse{
			Total:  count,
			Limit:  limit,
			Page:  page,
		},
        Message: "success retrived mitigation data",
		Data: mitigation,
	})
}

// @Summary Get an Mitigation by ID
// @Description Get an Mitigation by ID
// @Tags Mitigation
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the Mitigation"
// @Success 200 {object} mdl.MitigationSingleResp
// @Router /Mitigation/{id} [get]
func GetOneMitigation(c *gin.Context) {
    id := c.Params.ByName("id")
    var mitigation mdl.Mitigation

    // Find the record by ID
    if err := db.DB.Where("id = ?", id).First(&mitigation).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    // Return the found record as JSON
    c.JSON(http.StatusOK, mdl.MitigationSingleResp{
		Message: "success retrived mitigation data",
		Data:    mitigation,
	})
}

// @Summary Create an Mitigation
// @Description Create an Mitigation
// @Tags Mitigation
// @Accept  json
// @Produce  json
// @Param input body requests.MitigationRequest true "The Mitigation to create"
// @Success 201 {object}  mdl.GeneralResp
// @Router /Mitigation [post]
func CreateMitigation(c *gin.Context) {
    var input mdl.Mitigation

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
		Message: "Created Mitigation successfully",
	})
}

// @Summary Update an Mitigation
// @Description Update an Mitigation
// @Tags Mitigation
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the Mitigation"
// @Param input body requests.MitigationRequest true "The Mitigation to update"
// @Success 200 {object}  mdl.GeneralResp
// @Router /Mitigation/{id} [put]
// @Security BearerAuth
func UpdateMitigation(c *gin.Context) {
    id := c.Params.ByName("id")
    var input mdl.Mitigation

    // Bind JSON to the request struct
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Before update, make sure the data is exist
    var data mdl.Mitigation
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
    c.JSON(http.StatusOK, gin.H{"message": "Updated Mitigation successfully"})
}

// @Summary Delete an Mitigation
// @Description Delete an Mitigation
// @Tags Mitigation
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the Mitigation"
// @Success 200 {object}  mdl.GeneralResp
// @Router /Mitigation/{id} [delete]
func DeleteMitigation(c *gin.Context) {
    id := c.Params.ByName("id")
    // Before delete, make sure the data exists
    var data mdl.Mitigation
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
    c.JSON(http.StatusOK, gin.H{"message": "Deleted Mitigation successfully"})
}

