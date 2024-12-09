package routes

import (
	"ametory-crud/controllers"
	"github.com/gin-gonic/gin"
	mid "ametory-crud/middlewares"
)

func init() {
	Register("RiskTemplateRoutes", func(router *gin.RouterGroup) {
		var group = router.Group("RiskTemplate")
		group.Use(mid.AuthMiddleware())
		{
			group.GET("", mid.PermissionMiddleware([]string{"read:risk_template"}), controllers.GetRiskTemplate)
			group.GET("/:id", mid.PermissionMiddleware([]string{"read:risk_template"}), controllers.GetOneRiskTemplate)
			group.POST("", mid.PermissionMiddleware([]string{"create:risk_template"}), controllers.CreateRiskTemplate)
			group.PUT("/:id", mid.PermissionMiddleware([]string{"update:risk_template"}), controllers.UpdateRiskTemplate)
			group.DELETE("/:id", mid.PermissionMiddleware([]string{"delete:risk_template"}), controllers.DeleteRiskTemplate)
		}
	})
}

