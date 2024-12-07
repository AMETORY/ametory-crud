package routes

import (
	"ametory-crud/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	Register("AuthRoutes", func(router *gin.RouterGroup) {
		var group = router.Group("Auth")
		group.Use()
		{
			group.POST("/Login", controllers.LoginAuth)

		}
	})
}
