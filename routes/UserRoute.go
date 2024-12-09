package routes

import (
	"ametory-crud/controllers"
	"github.com/gin-gonic/gin"
	mid "ametory-crud/middlewares"
)

func init() {
	Register("UserRoutes", func(router *gin.RouterGroup) {
		var group = router.Group("User")
		group.Use(mid.AuthMiddleware())
		{
			group.GET("", mid.PermissionMiddleware([]string{"read:user"}), controllers.GetUser)
			group.GET("/:id", mid.PermissionMiddleware([]string{"read:user"}), controllers.GetOneUser)
			group.POST("", mid.PermissionMiddleware([]string{"create:user"}), controllers.CreateUser)
			group.PUT("/:id", mid.PermissionMiddleware([]string{"update:user"}), controllers.UpdateUser)
			group.DELETE("/:id", mid.PermissionMiddleware([]string{"delete:user"}), controllers.DeleteUser)
		}
	})
}

