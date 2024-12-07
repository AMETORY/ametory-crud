package main

import (
	"ametory-crud/cmd"
	"ametory-crud/config"
	"fmt"
	"log"
	"os"
)

func main() {
	// Load environment variables
	_, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Error initializing config: %s", err)
	}

	// Use config values

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Execute CLI
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
