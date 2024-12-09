package cmd

import (
	"ametory-crud/config"
	"ametory-crud/database"
	"ametory-crud/routes"
	"ametory-crud/services"
	"ametory-crud/workers"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the Gin server",
	Long:  "Start the Gin server and load routes",
	Run: func(cmd *cobra.Command, args []string) {
		initLog()
		database.ConnectDatabase()
		database.InitRedis()
		services.InitMail()

		gocron.Every(5).Seconds().Do(workers.SendRegMail)
		gocron.Start()
		fmt.Println("GO CRON STARTED")
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

func initLog() {
	t := time.Now()
	filename := t.Format("2006-01-02")
	logDir := "log"
	logPath := filepath.Join(logDir, filename+".log")

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("error creating directory: %v", err)
	}
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Println("Log file created:", logPath)
}

func startServer() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Origin frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Load application routes
	apiV1 := r.Group("/api/v1")

	routes.RegisterRoutes(apiV1)
	r.Static("/assets", "./assets")

	fmt.Printf("Server running on http://localhost:%s\n", config.App.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.App.Server.Port), r))
}
