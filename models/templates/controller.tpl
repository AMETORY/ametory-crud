package controllers

import (
    db "ametory-crud/database"
	mdl "ametory-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// @Summary Get all {{ToPascalCase .ModelName }}
// @Description Get all {{ToPascalCase .ModelName }}
// @Tags {{ToPascalCase .ModelName }}
// @Accept  json
// @Produce  json
// @Success 200 {object} mdl.{{ToPascalCase .ModelName }}Resp
// @Router /{{ToPascalCase .ModelName }} [get]
func Get{{ToPascalCase .ModelName }}(c *gin.Context) {
    var {{ToLower .ModelName }} []mdl.{{ToPascalCase .ModelName }}

    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    search := c.DefaultQuery("search", "")

    offset := (page - 1) * limit

    // Get the total count of records
    var count int64
    if err := db.DB.Model(&{{ToLower .ModelName }}).Count(&count).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Paginate
    if err := db.DB.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+search+"%").Find(&{{ToLower .ModelName }}).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, mdl.{{ToPascalCase .ModelName }}Resp{
		Pagination: mdl.PaginationResponse{
			Total:  count,
			Limit:  limit,
			Page:  page,
		},
        Message: "success retrived {{ToLower .ModelName }} data",
		Data: {{ToLower .ModelName }},
	})
}

// @Summary Get an {{ToPascalCase .ModelName }} by ID
// @Description Get an {{ToPascalCase .ModelName }} by ID
// @Tags {{ToPascalCase .ModelName }}
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the {{ToPascalCase .ModelName }}"
// @Success 200 {object} mdl.{{ToPascalCase .ModelName }}SingleResp
// @Router /{{ToPascalCase .ModelName }}/{id} [get]
func GetOne{{ToPascalCase .ModelName }}(c *gin.Context) {
    id := c.Params.ByName("id")
    var {{ToLower .ModelName }} mdl.{{ToPascalCase .ModelName }}

    // Find the record by ID
    if err := db.DB.Where("id = ?", id).First(&{{ToLower .ModelName }}).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    // Return the found record as JSON
    c.JSON(http.StatusOK, mdl.{{ToPascalCase .ModelName }}SingleResp{
		Message: "success retrived {{ToLower .ModelName }} data",
		Data:    {{ToLower .ModelName }},
	})
}

// @Summary Create an {{ToPascalCase .ModelName }}
// @Description Create an {{ToPascalCase .ModelName }}
// @Tags {{ToPascalCase .ModelName }}
// @Accept  json
// @Produce  json
// @Param input body requests.{{.ModelName}}Request true "The {{ToPascalCase .ModelName }} to create"
// @Success 201 {object}  mdl.GeneralResp
// @Router /{{ToPascalCase .ModelName }} [post]
func Create{{ToPascalCase .ModelName }}(c *gin.Context) {
    var input mdl.{{.ModelName}}

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
		Message: "Created {{.ModelName}} successfully",
	})
}

// @Summary Update an {{ToPascalCase .ModelName }}
// @Description Update an {{ToPascalCase .ModelName }}
// @Tags {{ToPascalCase .ModelName }}
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the {{ToPascalCase .ModelName }}"
// @Param input body requests.{{.ModelName}}Request true "The {{ToPascalCase .ModelName }} to update"
// @Success 200 {object}  mdl.GeneralResp
// @Router /{{ToPascalCase .ModelName }}/{id} [put]
// @Security BearerAuth
func Update{{ToPascalCase .ModelName }}(c *gin.Context) {
    id := c.Params.ByName("id")
    var input mdl.{{.ModelName}}

    // Bind JSON to the request struct
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Before update, make sure the data is exist
    var data mdl.{{.ModelName}}
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
    c.JSON(http.StatusOK, gin.H{"message": "Updated {{ToPascalCase .ModelName }} successfully"})
}

// @Summary Delete an {{ToPascalCase .ModelName }}
// @Description Delete an {{ToPascalCase .ModelName }}
// @Tags {{ToPascalCase .ModelName }}
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the {{ToPascalCase .ModelName }}"
// @Success 200 {object}  mdl.GeneralResp
// @Router /{{ToPascalCase .ModelName }}/{id} [delete]
func Delete{{ToPascalCase .ModelName }}(c *gin.Context) {
    id := c.Params.ByName("id")
    // Before delete, make sure the data exists
    var data mdl.{{.ModelName}}
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
    c.JSON(http.StatusOK, gin.H{"message": "Deleted {{ToPascalCase .ModelName }} successfully"})
}

