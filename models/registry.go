package models

import (
	db "ametory-crud/database"
)

var Models []interface{}

// RegisterModel appends a new model to the Models slice
func RegisterModel(model interface{}) {
	Models = append(Models, model)
}

func MigrateDatabase() {
	for _, model := range Models {
		db.DB.AutoMigrate(model)
	}
}