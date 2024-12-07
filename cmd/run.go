package cmd

import (
	"ametory-crud/config"
	"ametory-crud/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the Gin server",
	Long:  "Start the Gin server and load routes",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	// runCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server")
}

func startServer() {
	r := gin.Default()

	// Load application routes
	api := r.Group("/api")
	routes.RegisterRoutes(api)

	fmt.Printf("Server running on http://localhost:%s\n", config.App.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.App.Server.Port), r))
}
