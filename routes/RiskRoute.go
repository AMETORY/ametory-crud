package routes

import (
	"ametory-crud/controllers"
	"github.com/gin-gonic/gin"
	mid "ametory-crud/middlewares"
)

func init() {
	Register("RiskRoutes", func(router *gin.RouterGroup) {
		var group = router.Group("Risk")
		group.Use(mid.AuthMiddleware())
		{
			group.GET("", mid.PermissionMiddleware([]string{"read:risk"}), controllers.GetRisk)
			group.GET("/:id", mid.PermissionMiddleware([]string{"read:risk"}), controllers.GetOneRisk)
			group.POST("", mid.PermissionMiddleware([]string{"create:risk"}), controllers.CreateRisk)
			group.PUT("/:id", mid.PermissionMiddleware([]string{"update:risk"}), controllers.UpdateRisk)
			group.DELETE("/:id", mid.PermissionMiddleware([]string{"delete:risk"}), controllers.DeleteRisk)
		}
	})
}

