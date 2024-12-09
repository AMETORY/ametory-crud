package controllers

import (
    db "ametory-crud/database"
	mdl "ametory-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// @Summary Get all Company
// @Description Get all Company
// @Tags Company
// @Accept  json
// @Produce  json
// @Success 200 {object} mdl.CompanyResp
// @Router /Company [get]
func GetCompany(c *gin.Context) {
    var company []mdl.Company

    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    search := c.DefaultQuery("search", "")

    offset := (page - 1) * limit

    // Get the total count of records
    var count int64
    if err := db.DB.Model(&company).Count(&count).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Paginate
    if err := db.DB.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+search+"%").Find(&company).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, mdl.CompanyResp{
		Pagination: mdl.PaginationResponse{
			Total:  count,
			Limit:  limit,
			Page:  page,
		},
        Message: "success retrived company data",
		Data: company,
	})
}

// @Summary Get an Company by ID
// @Description Get an Company by ID
// @Tags Company
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the Company"
// @Success 200 {object} mdl.CompanySingleResp
// @Router /Company/{id} [get]
func GetOneCompany(c *gin.Context) {
    id := c.Params.ByName("id")
    var company mdl.Company

    // Find the record by ID
    if err := db.DB.Where("id = ?", id).First(&company).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    // Return the found record as JSON
    c.JSON(http.StatusOK, mdl.CompanySingleResp{
		Message: "success retrived company data",
		Data:    company,
	})
}

// @Summary Create an Company
// @Description Create an Company
// @Tags Company
// @Accept  json
// @Produce  json
// @Param input body requests.CompanyRequest true "The Company to create"
// @Success 201 {object}  mdl.GeneralResp
// @Router /Company [post]
func CreateCompany(c *gin.Context) {
    var input mdl.Company

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
		Message: "Created Company successfully",
	})
}

// @Summary Update an Company
// @Description Update an Company
// @Tags Company
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the Company"
// @Param input body requests.CompanyRequest true "The Company to update"
// @Success 200 {object}  mdl.GeneralResp
// @Router /Company/{id} [put]
// @Security BearerAuth
func UpdateCompany(c *gin.Context) {
    id := c.Params.ByName("id")
    var input mdl.Company

    // Bind JSON to the request struct
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Before update, make sure the data is exist
    var data mdl.Company
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
    c.JSON(http.StatusOK, gin.H{"message": "Updated Company successfully"})
}

// @Summary Delete an Company
// @Description Delete an Company
// @Tags Company
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the Company"
// @Success 200 {object}  mdl.GeneralResp
// @Router /Company/{id} [delete]
func DeleteCompany(c *gin.Context) {
    id := c.Params.ByName("id")
    // Before delete, make sure the data exists
    var data mdl.Company
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
    c.JSON(http.StatusOK, gin.H{"message": "Deleted Company successfully"})
}

