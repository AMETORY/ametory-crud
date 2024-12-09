package routes

import (
	"ametory-crud/controllers"
	"ametory-crud/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	Register("AuthRoutes", func(router *gin.RouterGroup) {
		var group = router.Group("Auth")
		group.Use()
		{
			group.POST("/Login", controllers.LoginAuth)
			group.POST("/Registration", controllers.RegisterUser)
			group.GET("/Verification/:id", controllers.Verification)
			group.GET("/Profile", middlewares.AuthMiddleware(), controllers.Profile)

		}
	})
}
