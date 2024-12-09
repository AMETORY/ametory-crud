package routes

import (
	"ametory-crud/controllers"
	"github.com/gin-gonic/gin"
	mid "ametory-crud/middlewares"
)

func init() {
	Register("RiskAssessmentRoutes", func(router *gin.RouterGroup) {
		var group = router.Group("RiskAssessment")
		group.Use(mid.AuthMiddleware())
		{
			group.GET("", mid.PermissionMiddleware([]string{"read:risk_assessment"}), controllers.GetRiskAssessment)
			group.GET("/:id", mid.PermissionMiddleware([]string{"read:risk_assessment"}), controllers.GetOneRiskAssessment)
			group.POST("", mid.PermissionMiddleware([]string{"create:risk_assessment"}), controllers.CreateRiskAssessment)
			group.PUT("/:id", mid.PermissionMiddleware([]string{"update:risk_assessment"}), controllers.UpdateRiskAssessment)
			group.DELETE("/:id", mid.PermissionMiddleware([]string{"delete:risk_assessment"}), controllers.DeleteRiskAssessment)
		}
	})
}

