package controllers

import (
    db "ametory-crud/database"
	mdl "ametory-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// @Summary Get all Objective
// @Description Get all Objective
// @Tags Objective
// @Accept  json
// @Produce  json
// @Success 200 {object} mdl.ObjectiveResp
// @Router /Objective [get]
func GetObjective(c *gin.Context) {
    var objective []mdl.Objective

    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    search := c.DefaultQuery("search", "")

    offset := (page - 1) * limit

    // Get the total count of records
    var count int64
    if err := db.DB.Model(&objective).Count(&count).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Paginate
    if err := db.DB.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+search+"%").Find(&objective).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, mdl.ObjectiveResp{
		Pagination: mdl.PaginationResponse{
			Total:  count,
			Limit:  limit,
			Page:  page,
		},
        Message: "success retrived objective data",
		Data: objective,
	})
}

// @Summary Get an Objective by ID
// @Description Get an Objective by ID
// @Tags Objective
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the Objective"
// @Success 200 {object} mdl.ObjectiveSingleResp
// @Router /Objective/{id} [get]
func GetOneObjective(c *gin.Context) {
    id := c.Params.ByName("id")
    var objective mdl.Objective

    // Find the record by ID
    if err := db.DB.Where("id = ?", id).First(&objective).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    // Return the found record as JSON
    c.JSON(http.StatusOK, mdl.ObjectiveSingleResp{
		Message: "success retrived objective data",
		Data:    objective,
	})
}

// @Summary Create an Objective
// @Description Create an Objective
// @Tags Objective
// @Accept  json
// @Produce  json
// @Param input body requests.ObjectiveRequest true "The Objective to create"
// @Success 201 {object}  mdl.GeneralResp
// @Router /Objective [post]
func CreateObjective(c *gin.Context) {
    var input mdl.Objective

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
		Message: "Created Objective successfully",
	})
}

// @Summary Update an Objective
// @Description Update an Objective
// @Tags Objective
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the Objective"
// @Param input body requests.ObjectiveRequest true "The Objective to update"
// @Success 200 {object}  mdl.GeneralResp
// @Router /Objective/{id} [put]
// @Security BearerAuth
func UpdateObjective(c *gin.Context) {
    id := c.Params.ByName("id")
    var input mdl.Objective

    // Bind JSON to the request struct
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Before update, make sure the data is exist
    var data mdl.Objective
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
    c.JSON(http.StatusOK, gin.H{"message": "Updated Objective successfully"})
}

// @Summary Delete an Objective
// @Description Delete an Objective
// @Tags Objective
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the Objective"
// @Success 200 {object}  mdl.GeneralResp
// @Router /Objective/{id} [delete]
func DeleteObjective(c *gin.Context) {
    id := c.Params.ByName("id")
    // Before delete, make sure the data exists
    var data mdl.Objective
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
    c.JSON(http.StatusOK, gin.H{"message": "Deleted Objective successfully"})
}

