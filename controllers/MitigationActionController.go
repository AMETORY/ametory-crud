package controllers

import (
    db "ametory-crud/database"
	mdl "ametory-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// @Summary Get all MitigationAction
// @Description Get all MitigationAction
// @Tags MitigationAction
// @Accept  json
// @Produce  json
// @Success 200 {object} mdl.MitigationActionResp
// @Router /MitigationAction [get]
func GetMitigationAction(c *gin.Context) {
    var mitigation_action []mdl.MitigationAction

    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    search := c.DefaultQuery("search", "")

    offset := (page - 1) * limit

    // Get the total count of records
    var count int64
    if err := db.DB.Model(&mitigation_action).Count(&count).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Paginate
    if err := db.DB.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+search+"%").Find(&mitigation_action).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, mdl.MitigationActionResp{
		Pagination: mdl.PaginationResponse{
			Total:  count,
			Limit:  limit,
			Page:  page,
		},
        Message: "success retrived mitigation_action data",
		Data: mitigation_action,
	})
}

// @Summary Get an MitigationAction by ID
// @Description Get an MitigationAction by ID
// @Tags MitigationAction
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the MitigationAction"
// @Success 200 {object} mdl.MitigationActionSingleResp
// @Router /MitigationAction/{id} [get]
func GetOneMitigationAction(c *gin.Context) {
    id := c.Params.ByName("id")
    var mitigation_action mdl.MitigationAction

    // Find the record by ID
    if err := db.DB.Where("id = ?", id).First(&mitigation_action).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    // Return the found record as JSON
    c.JSON(http.StatusOK, mdl.MitigationActionSingleResp{
		Message: "success retrived mitigation_action data",
		Data:    mitigation_action,
	})
}

// @Summary Create an MitigationAction
// @Description Create an MitigationAction
// @Tags MitigationAction
// @Accept  json
// @Produce  json
// @Param input body requests.MitigationActionRequest true "The MitigationAction to create"
// @Success 201 {object}  mdl.GeneralResp
// @Router /MitigationAction [post]
func CreateMitigationAction(c *gin.Context) {
    var input mdl.MitigationAction

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
		Message: "Created MitigationAction successfully",
	})
}

// @Summary Update an MitigationAction
// @Description Update an MitigationAction
// @Tags MitigationAction
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the MitigationAction"
// @Param input body requests.MitigationActionRequest true "The MitigationAction to update"
// @Success 200 {object}  mdl.GeneralResp
// @Router /MitigationAction/{id} [put]
// @Security BearerAuth
func UpdateMitigationAction(c *gin.Context) {
    id := c.Params.ByName("id")
    var input mdl.MitigationAction

    // Bind JSON to the request struct
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Before update, make sure the data is exist
    var data mdl.MitigationAction
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
    c.JSON(http.StatusOK, gin.H{"message": "Updated MitigationAction successfully"})
}

// @Summary Delete an MitigationAction
// @Description Delete an MitigationAction
// @Tags MitigationAction
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the MitigationAction"
// @Success 200 {object}  mdl.GeneralResp
// @Router /MitigationAction/{id} [delete]
func DeleteMitigationAction(c *gin.Context) {
    id := c.Params.ByName("id")
    // Before delete, make sure the data exists
    var data mdl.MitigationAction
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
    c.JSON(http.StatusOK, gin.H{"message": "Deleted MitigationAction successfully"})
}

