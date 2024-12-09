package routes

import (
	"ametory-crud/controllers"
	"ametory-crud/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	Register("MiscRoutes", func(router *gin.RouterGroup) {
		// var group = router.Group("Misc")
		router.Use(middlewares.AuthMiddleware())
		{
			router.POST("FileUpload", controllers.FileUpload)
		}
	})
}
