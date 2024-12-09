package routes

import (
	"ametory-crud/controllers"
	"github.com/gin-gonic/gin"
	mid "ametory-crud/middlewares"
)

func init() {
	Register("ObjectiveRoutes", func(router *gin.RouterGroup) {
		var group = router.Group("Objective")
		group.Use(mid.AuthMiddleware())
		{
			group.GET("", mid.PermissionMiddleware([]string{"read:objective"}), controllers.GetObjective)
			group.GET("/:id", mid.PermissionMiddleware([]string{"read:objective"}), controllers.GetOneObjective)
			group.POST("", mid.PermissionMiddleware([]string{"create:objective"}), controllers.CreateObjective)
			group.PUT("/:id", mid.PermissionMiddleware([]string{"update:objective"}), controllers.UpdateObjective)
			group.DELETE("/:id", mid.PermissionMiddleware([]string{"delete:objective"}), controllers.DeleteObjective)
		}
	})
}

