package cmd

import (
	"ametory-crud/database"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var excelPath string

var generateExcelCmd = &cobra.Command{
	Use:   "generate-from-excel",
	Short: "Generate models, controllers, and routes from an Excel file",
	Run: func(cmd *cobra.Command, args []string) {
		if excelPath == "" {
			fmt.Println("Please provide the path to the Excel file using --path")
			return
		}
		database.ConnectDatabase()
		generateFromExcel(excelPath)
	},
}

func generateFromExcel(path string) error {
	var fields []Field

	// Open the Excel file
	f, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}

	// Get the sheet names (assuming the data is in the first sheet)
	sheet := f.GetSheetName(0)

	// Iterate through rows (skip the header row)
	rows, err := f.GetRows(sheet)

	// fmt.Println(path, rows)
	if err != nil {
		return err
	}

	// Loop over rows
	for i, row := range rows {
		// Skip header row
		if i == 0 {
			continue
		}

		// Read the data from each row
		if len(row) >= 4 {
			field := Field{
				ModelName: row[0],
				Name:      row[1],
				Type:      row[2],
				DBType:    row[3],
				Tag:       cases.Lower(language.English).String(strings.ReplaceAll(row[1], "_", "")),
			}
			fields = append(fields, field)
		}
	}

	modelFields := groupFieldsByModel(fields)

	// Generate models based on the grouped data
	for modelName, fields := range modelFields {
		err := generateModel(modelName, fields)
		if err != nil {
			log.Fatalf("Error generating model %s: %s", modelName, err)
		}
		err = generateRequestResponse(modelName, fields)
		if err != nil {
			log.Fatalf("Error generating request %s: %s", modelName, err)
		}
		generateController(modelName, fields)
		generateRoute(modelName)
	}

	return nil
}

func groupFieldsByModel(fields []Field) map[string][]Field {
	modelFields := make(map[string][]Field)
	for _, field := range fields {
		modelFields[field.ModelName] = append(modelFields[field.ModelName], field)
	}
	return modelFields
}

func generateGORMTag(constraints string) string {
	constraintsList := strings.Split(constraints, ",")
	var gormTags []string

	for _, c := range constraintsList {
		c = strings.TrimSpace(c)
		switch c {
		case "primaryKey":
			gormTags = append(gormTags, "primaryKey")
		case "not null":
			gormTags = append(gormTags, "not null")
		case "unique":
			gormTags = append(gormTags, "unique")
		default:
			gormTags = append(gormTags, c)
		}
	}
	return fmt.Sprintf(`gorm:"%s"`, strings.Join(gormTags, ";"))
}
