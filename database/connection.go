package database

import (
	"ametory-crud/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dbType := config.App.Database.Type

	if dbType == "postgres" {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			config.App.Database.Host,
			config.App.Database.User,
			config.App.Database.Password,
			config.App.Database.Name,
			config.App.Database.Port,
		)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else if dbType == "mysql" {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.App.Database.User,
			config.App.Database.Password,
			config.App.Database.Host,
			config.App.Database.Port,
			config.App.Database.Name,
		)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		log.Fatalf("Unsupported database type: %s", dbType)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Printf("Connected to %s database successfully.\n", dbType)
}
