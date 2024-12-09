package routes

import (
	"ametory-crud/controllers"
	"github.com/gin-gonic/gin"
	mid "ametory-crud/middlewares"
)

func init() {
	Register("{{ToPascalCase .ModelName }}Routes", func(router *gin.RouterGroup) {
		var group = router.Group("{{ToPascalCase .ModelName }}")
		group.Use(mid.AuthMiddleware())
		{
			group.GET("", mid.PermissionMiddleware([]string{"read:{{ToLower .ModelName }}"}), controllers.Get{{ToPascalCase .ModelName }})
			group.GET("/:id", mid.PermissionMiddleware([]string{"read:{{ToLower .ModelName }}"}), controllers.GetOne{{ToPascalCase .ModelName }})
			group.POST("", mid.PermissionMiddleware([]string{"create:{{ToLower .ModelName }}"}), controllers.Create{{ToPascalCase .ModelName }})
			group.PUT("/:id", mid.PermissionMiddleware([]string{"update:{{ToLower .ModelName }}"}), controllers.Update{{ToPascalCase .ModelName }})
			group.DELETE("/:id", mid.PermissionMiddleware([]string{"delete:{{ToLower .ModelName }}"}), controllers.Delete{{ToPascalCase .ModelName }})
		}
	})
}

