package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
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
		generateFromExcel(excelPath)
	},
}

func generateFromExcel(path string) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Printf("Error opening Excel file: %v\n", err)
		return
	}
	defer f.Close()

	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		fmt.Printf("Error reading rows: %v\n", err)
		return
	}

	features := map[string][]string{}
	columns := map[string][]string{}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		feature := row[0]
		columnName := row[1]
		columnType := row[2]
		constraints := row[3]

		colDef := fmt.Sprintf("%s %s `%s`", columnName, columnType, generateGORMTag(constraints))
		features[feature] = append(features[feature], colDef)
		columns[feature] = append(columns[feature], columnName)
	}

	for feature, colDefs := range features {
		fmt.Printf("Generating files for feature: %s\n", feature)
		generateModel(feature, colDefs)
		generateController(feature)
		generateRoute(feature)
	}
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
