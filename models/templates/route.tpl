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
			group.GET("", mid.PermissionMiddleware([]string{"read:{{ .ModelName | ToLower }}"}), controllers.Get{{ .ModelName }})
			group.GET("/:id", mid.PermissionMiddleware([]string{"read:{{ .ModelName | ToLower }}"}), controllers.GetOne{{ .ModelName }})
			group.POST("", mid.PermissionMiddleware([]string{"create:{{ .ModelName | ToLower }}"}), controllers.Create{{ .ModelName }})
			group.PUT("/:id", mid.PermissionMiddleware([]string{"update:{{ .ModelName | ToLower }}"}), controllers.Update{{ .ModelName }})
			group.DELETE("/:id", mid.PermissionMiddleware([]string{"delete:{{ .ModelName | ToLower }}"}), controllers.Delete{{ .ModelName }})
		}
	})
}

