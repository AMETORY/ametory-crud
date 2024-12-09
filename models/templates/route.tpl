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
			group.GET("", mid.PermissionMiddleware([]string{"read:{{ToSnakeCase .ModelName }}"}), controllers.Get{{ToPascalCase .ModelName }})
			group.GET("/:id", mid.PermissionMiddleware([]string{"read:{{ToSnakeCase .ModelName }}"}), controllers.GetOne{{ToPascalCase .ModelName }})
			group.POST("", mid.PermissionMiddleware([]string{"create:{{ToSnakeCase .ModelName }}"}), controllers.Create{{ToPascalCase .ModelName }})
			group.PUT("/:id", mid.PermissionMiddleware([]string{"update:{{ToSnakeCase .ModelName }}"}), controllers.Update{{ToPascalCase .ModelName }})
			group.DELETE("/:id", mid.PermissionMiddleware([]string{"delete:{{ToSnakeCase .ModelName }}"}), controllers.Delete{{ToPascalCase .ModelName }})
		}
	})
}

