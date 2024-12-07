package cmd

import (
	"ametory-crud/config"
	"ametory-crud/database"
	"ametory-crud/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the Gin server",
	Long:  "Start the Gin server and load routes",
	Run: func(cmd *cobra.Command, args []string) {
		database.ConnectDatabase()
		startServer()
	},
}
var swagCmd = &cobra.Command{
	Use:   "swag",
	Short: "Swag init",
	Long:  "Swag init",
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("swag", "init")
		command.Stdout = os.Stdout
		command.Stderr = os.Stdout
		if err := command.Run(); err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(swagCmd)
	rootCmd.AddCommand(runCmd)
	// runCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server")
}

func startServer() {
	r := gin.Default()

	// Load application routes
	apiV1 := r.Group("/api/v1")
	routes.RegisterRoutes(apiV1)

	fmt.Printf("Server running on http://localhost:%s\n", config.App.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.App.Server.Port), r))
}
