package routes

import (
	"ametory-crud/controllers"
	"github.com/gin-gonic/gin"
	mid "ametory-crud/middlewares"
)

func init() {
	Register("{{ .ModelName }}Routes", func(router *gin.RouterGroup) {
		var group = router.Group("{{ .ModelName }}")
		group.Use(mid.AuthMiddleware())
		{
			group.GET("", controllers.Get{{ .ModelName }})
			group.GET("/:id", controllers.GetOne{{ .ModelName }})
			group.POST("", controllers.Create{{ .ModelName }})
			group.PUT("/:id", controllers.Update{{ .ModelName }})
			group.DELETE("/:id", controllers.Delete{{ .ModelName }})
		}
	})
}

