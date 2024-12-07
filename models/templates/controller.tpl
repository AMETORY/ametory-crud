package controllers

import (
    db "ametory-crud/database"
	mdl "ametory-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// @Summary Get all {{ .ModelName }}
// @Description Get all {{ .ModelName }}
// @Tags {{ .ModelName }}
// @Accept  json
// @Produce  json
// @Success 200 {object} mdl.{{ .ModelName }}Resp
// @Router /{{ .ModelName }} [get]
func Get{{ .ModelName }}(c *gin.Context) {
    var {{ .ModelName | ToLower }} []mdl.{{ .ModelName }}

    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
    search := c.DefaultQuery("search", "")

    // Get the total count of records
    var count int64
    if err := db.DB.Model(&{{ .ModelName | ToLower }}).Count(&count).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Paginate
    if err := db.DB.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+search+"%").Find(&{{ .ModelName | ToLower }}).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, mdl.{{ .ModelName }}Resp{
		Pagination: mdl.PaginationResponse{
			Total:  count,
			Limit:  limit,
			Offset: offset,
		},
		Data: {{ .ModelName | ToLower }},
	})
}

// @Summary Get an {{ .ModelName }} by ID
// @Description Get an {{ .ModelName }} by ID
// @Tags {{ .ModelName }}
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the {{ .ModelName }}"
// @Success 200 {object} mdl.{{ .ModelName }}SingleResp
// @Router /{{ .ModelName }}/{id} [get]
func GetOne{{ .ModelName }}(c *gin.Context) {
    id := c.Params.ByName("id")
    var {{ .ModelName | ToLower }} mdl.{{ .ModelName }}

    // Find the record by ID
    if err := db.DB.Where("id = ?", id).First(&{{ .ModelName | ToLower }}).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    // Return the found record as JSON
    c.JSON(http.StatusOK, mdl.{{ .ModelName }}SingleResp{
		Message: "success retrived {{ .ModelName | ToLower }} data",
		Data:    {{ .ModelName | ToLower }},
	})
}

// @Summary Create an {{ .ModelName }}
// @Description Create an {{ .ModelName }}
// @Tags {{ .ModelName }}
// @Accept  json
// @Produce  json
// @Param input body requests.{{.ModelName}}Request true "The {{ .ModelName }} to create"
// @Success 201 {object}  mdl.GeneralResp
// @Router /{{ .ModelName }} [post]
func Create{{ .ModelName }}(c *gin.Context) {
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

// @Summary Update an {{ .ModelName }}
// @Description Update an {{ .ModelName }}
// @Tags {{ .ModelName }}
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the {{ .ModelName }}"
// @Param input body requests.{{.ModelName}}Request true "The {{ .ModelName }} to update"
// @Success 200 {object}  mdl.GeneralResp
// @Router /{{ .ModelName }}/{id} [put]
func Update{{ .ModelName }}(c *gin.Context) {
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
    c.JSON(http.StatusOK, gin.H{"message": "Updated {{ .ModelName }} successfully"})
}

// @Summary Delete an {{ .ModelName }}
// @Description Delete an {{ .ModelName }}
// @Tags {{ .ModelName }}
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the {{ .ModelName }}"
// @Success 200 {object}  mdl.GeneralResp
// @Router /{{ .ModelName }}/{id} [delete]
func Delete{{ .ModelName }}(c *gin.Context) {
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
    c.JSON(http.StatusOK, gin.H{"message": "Deleted {{ .ModelName }} successfully"})
}

