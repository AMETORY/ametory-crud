package routes

import (
	"ametory-crud/controllers"
	"github.com/gin-gonic/gin"
	mid "ametory-crud/middlewares"
)

func init() {
	Register("MitigationActionRoutes", func(router *gin.RouterGroup) {
		var group = router.Group("MitigationAction")
		group.Use(mid.AuthMiddleware())
		{
			group.GET("", mid.PermissionMiddleware([]string{"read:mitigation_action"}), controllers.GetMitigationAction)
			group.GET("/:id", mid.PermissionMiddleware([]string{"read:mitigation_action"}), controllers.GetOneMitigationAction)
			group.POST("", mid.PermissionMiddleware([]string{"create:mitigation_action"}), controllers.CreateMitigationAction)
			group.PUT("/:id", mid.PermissionMiddleware([]string{"update:mitigation_action"}), controllers.UpdateMitigationAction)
			group.DELETE("/:id", mid.PermissionMiddleware([]string{"delete:mitigation_action"}), controllers.DeleteMitigationAction)
		}
	})
}

