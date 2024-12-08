package routes

import (
	"ametory-crud/config"
	"ametory-crud/docs"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

var RegistryMap = map[string]func(*gin.RouterGroup){}

// RegisterRoutes loads all route groups
func RegisterRoutes(r *gin.RouterGroup) {
	// var swaggerHost = strings.Replace(config.App.Server.ApiURL, "http://", "", -1)
	// swaggerHost = strings.Replace(swaggerHost, "https://", "", -1)
	docs.SwaggerInfo.Host = strings.Replace(config.App.Server.ApiURL, "http://", "", -1)
	docs.SwaggerInfo.Title = config.App.Server.AppName
	docs.SwaggerInfo.Description = config.App.Server.AppDesc
	docs.SwaggerInfo.Version = config.App.Server.AppDesc
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Example: Register user routes here
	// User routes can be added like this:
	// RegisterUserRoutes(r)
	// ADD BELOW
	for name, registerFunc := range RegistryMap {
		fmt.Println("REGISTER ROUTE", name)
		registerFunc(r)
	}

}

// Register adds a route registration function to the RegistryMap
func Register(name string, registerFunc func(*gin.RouterGroup)) {
	RegistryMap[name] = registerFunc
}
