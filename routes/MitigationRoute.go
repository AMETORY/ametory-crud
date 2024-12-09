package routes

import (
	"ametory-crud/controllers"
	"github.com/gin-gonic/gin"
	mid "ametory-crud/middlewares"
)

func init() {
	Register("MitigationRoutes", func(router *gin.RouterGroup) {
		var group = router.Group("Mitigation")
		group.Use(mid.AuthMiddleware())
		{
			group.GET("", mid.PermissionMiddleware([]string{"read:mitigation"}), controllers.GetMitigation)
			group.GET("/:id", mid.PermissionMiddleware([]string{"read:mitigation"}), controllers.GetOneMitigation)
			group.POST("", mid.PermissionMiddleware([]string{"create:mitigation"}), controllers.CreateMitigation)
			group.PUT("/:id", mid.PermissionMiddleware([]string{"update:mitigation"}), controllers.UpdateMitigation)
			group.DELETE("/:id", mid.PermissionMiddleware([]string{"delete:mitigation"}), controllers.DeleteMitigation)
		}
	})
}

