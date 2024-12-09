package routes

import (
	"ametory-crud/controllers"
	"ametory-crud/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	Register("RoleRoutes", func(router *gin.RouterGroup) {
		var group = router.Group("Role")
		group.Use(middlewares.AuthMiddleware(), middlewares.SuperAdminMiddleware())
		{
			group.GET("", controllers.GetRoles)
			group.GET("/:id", controllers.GetOneRole)
			group.POST("", controllers.CreateRole)
			group.PUT("/:id", controllers.UpdateRole)
			group.DELETE("/:id", controllers.DeleteRole)
		}
	})
}
