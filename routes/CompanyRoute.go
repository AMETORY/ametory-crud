package routes

import (
	"ametory-crud/controllers"
	"github.com/gin-gonic/gin"
	mid "ametory-crud/middlewares"
)

func init() {
	Register("CompanyRoutes", func(router *gin.RouterGroup) {
		var group = router.Group("Company")
		group.Use(mid.AuthMiddleware())
		{
			group.GET("", mid.PermissionMiddleware([]string{"read:company"}), controllers.GetCompany)
			group.GET("/:id", mid.PermissionMiddleware([]string{"read:company"}), controllers.GetOneCompany)
			group.POST("", mid.PermissionMiddleware([]string{"create:company"}), controllers.CreateCompany)
			group.PUT("/:id", mid.PermissionMiddleware([]string{"update:company"}), controllers.UpdateCompany)
			group.DELETE("/:id", mid.PermissionMiddleware([]string{"delete:company"}), controllers.DeleteCompany)
		}
	})
}

