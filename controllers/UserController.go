package controllers

import (
    db "ametory-crud/database"
	mdl "ametory-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// @Summary Get all User
// @Description Get all User
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} mdl.UserResp
// @Router /User [get]
func GetUser(c *gin.Context) {
    var user []mdl.User

    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    search := c.DefaultQuery("search", "")

    offset := (page - 1) * limit

    // Get the total count of records
    var count int64
    if err := db.DB.Model(&user).Count(&count).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Paginate
    if err := db.DB.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+search+"%").Find(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, mdl.UserResp{
		Pagination: mdl.PaginationResponse{
			Total:  count,
			Limit:  limit,
			Page:  page,
		},
        Message: "success retrived user data",
		Data: user,
	})
}

// @Summary Get an User by ID
// @Description Get an User by ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the User"
// @Success 200 {object} mdl.UserSingleResp
// @Router /User/{id} [get]
func GetOneUser(c *gin.Context) {
    id := c.Params.ByName("id")
    var user mdl.User

    // Find the record by ID
    if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    // Return the found record as JSON
    c.JSON(http.StatusOK, mdl.UserSingleResp{
		Message: "success retrived user data",
		Data:    user,
	})
}

// @Summary Create an User
// @Description Create an User
// @Tags User
// @Accept  json
// @Produce  json
// @Param input body requests.UserRequest true "The User to create"
// @Success 201 {object}  mdl.GeneralResp
// @Router /User [post]
func CreateUser(c *gin.Context) {
    var input mdl.User

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
		Message: "Created User successfully",
	})
}

// @Summary Update an User
// @Description Update an User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the User"
// @Param input body requests.UserRequest true "The User to update"
// @Success 200 {object}  mdl.GeneralResp
// @Router /User/{id} [put]
// @Security BearerAuth
func UpdateUser(c *gin.Context) {
    id := c.Params.ByName("id")
    var input mdl.User

    // Bind JSON to the request struct
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Before update, make sure the data is exist
    var data mdl.User
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
    c.JSON(http.StatusOK, gin.H{"message": "Updated User successfully"})
}

// @Summary Delete an User
// @Description Delete an User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the User"
// @Success 200 {object}  mdl.GeneralResp
// @Router /User/{id} [delete]
func DeleteUser(c *gin.Context) {
    id := c.Params.ByName("id")
    // Before delete, make sure the data exists
    var data mdl.User
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
    c.JSON(http.StatusOK, gin.H{"message": "Deleted User successfully"})
}

