package routes

import (
	"ametory-crud/controllers"
	"github.com/gin-gonic/gin"
)

func Register{{ .Feature }}Routes(r *gin.RouterGroup) {
	r.GET("/{{ .Feature | ToLower }}", controllers.Get{{ .Feature }})
	r.POST("/{{ .Feature | ToLower }}", controllers.Create{{ .Feature }})
	r.PUT("/{{ .Feature | ToLower }}/:id", controllers.Update{{ .Feature }})
	r.DELETE("/{{ .Feature | ToLower }}/:id", controllers.Delete{{ .Feature }})
}
