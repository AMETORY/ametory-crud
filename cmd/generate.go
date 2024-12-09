package cmd

import (
	"ametory-crud/config"
	"ametory-crud/database"
	"ametory-crud/models"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"

	"log"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var generateCmd = &cobra.Command{
	Use:   "generate [type] [name]",
	Short: "Generate a new model, controller, or route",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		generateType := strings.ToLower(args[0])
		name := cases.Title(language.English).String(args[1])
		fields := []Field{}
		isHasTime := false
		for _, arg := range args[2:] {
			parts := strings.Split(arg, ":")
			fmt.Println(parts)
			if len(parts) != 3 {
				fmt.Printf("Invalid field format: %s\n", arg)
				return
			}
			field := Field{
				ModelName: name,
				Name:      cases.Title(language.English).String(parts[0]),
				Type:      parts[1],
				DBType:    parts[2],
				Tag:       cases.Lower(language.English).String(strings.ReplaceAll(parts[0], " ", "_")),
			}

			if parts[1] == "time.Time" {
				isHasTime = true
			}

			fields = append(fields, field)
		}
		switch generateType {
		case "model":

			generateModel(name, fields, isHasTime)
		case "request":

			generateRequestResponse(name, fields, isHasTime)
		case "controller":
			generateController(name, fields, isHasTime)
		case "route":
			generateRoute(name, isHasTime)

		default:
			fmt.Println("Invalid type. Use 'model', 'controller', or 'route'.")
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(generateExcelCmd)
	generateExcelCmd.Flags().StringVar(&excelPath, "path", "", "Path to the Excel file")
}

// Fungsi untuk menghasilkan file model

// ModelData holds the information about the model and its fields
type ModelData struct {
	ModelName string
	Fields    []Field
	IsHasTime bool
}

// Field holds details about each field in the model
type Field struct {
	ModelName string
	Name      string
	Type      string
	DBType    string
	Tag       string
}

func generateModel(modelName string, fields []Field, IsHasTime bool) error {

	var crudMethods = [...]string{"create", "read", "update", "delete"}

	for _, crudMethod := range crudMethods {
		// Add permission logic
		permission := models.Permission{
			Name:        fmt.Sprintf("%s %s", cases.Title(language.English).String(modelName), cases.Title(language.English).String(crudMethod)),
			Description: fmt.Sprintf("Permission to %s %s", crudMethod, modelName),
			Key:         fmt.Sprintf("%s:%s", strings.ToLower(crudMethod), ToSnakeCase(modelName)),
			Group:       ToPascalCase(modelName),
		}
		permission.ID = models.GenUUID()

		// Assuming you have a function to add permissions to your system
		err := database.DB.Create(&permission).Error
		if err != nil {
			log.Println("Error adding permission: %v", err)
		}
	}

	// Define the model data
	modelData := ModelData{
		ModelName: modelName,
		Fields:    fields,
		IsHasTime: IsHasTime,
	}

	if config.App.Database.Type == "postgres" {
		for i, v := range fields {
			if strings.Contains(v.DBType, "enum") {
				enumTypeName := fmt.Sprintf("%s_%s_enum", ToSnakeCase(modelName), ToSnakeCase(v.Name))
				cleanDBType := strings.ReplaceAll(v.DBType, "NOT NULL", "")
				cleanDBType = regexp.MustCompile(`DEFAULT '([^']*)'`).ReplaceAllString(cleanDBType, "")
				cleanDBType = strings.ReplaceAll(cleanDBType, "DEFAULT", "")
				rawSQL := fmt.Sprintf("CREATE TYPE %s AS %s;", enumTypeName, cleanDBType)
				if err := database.DB.Exec(rawSQL).Error; err != nil {
					log.Printf("Error creating enum type: %v", err)
				}
				fields[i].DBType = enumTypeName
			}
		}
	}

	// Open the template file
	tmpl, err := template.New("model.tpl").Funcs(template.FuncMap{
		"ToLower":      strings.ToLower,
		"ToPascalCase": ToPascalCase,
		"ToSnakeCase":  ToSnakeCase,
	}).ParseFiles("models/templates/model.tpl") // ensure template file exists
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
		return err
	}

	// Create the model file
	fileName := fmt.Sprintf("models/%s.go", ToPascalCase(modelName))
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error creating model file: %v", err)
		return err
	}
	defer file.Close()

	// Execute the template and generate the model
	err = tmpl.Execute(file, modelData)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
		return err
	}

	fmt.Printf("Model %s generated successfully!\n", ToPascalCase(modelName))
	return nil
}
func generateRequestResponse(modelName string, fields []Field, IsHasTime bool) error {
	// Define the model data
	modelData := ModelData{
		ModelName: modelName,
		Fields:    fields,
		IsHasTime: IsHasTime,
	}

	// Open the template file
	tmpl, err := template.New("request_response.tpl").Funcs(template.FuncMap{
		"ToLower":      strings.ToLower,
		"ToPascalCase": ToPascalCase,
		"ToSnakeCase":  ToSnakeCase,
	}).ParseFiles("models/templates/request_response.tpl") // ensure template file exists
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
		return err
	}

	// Create the model file
	fileName := fmt.Sprintf("requests/%sReq.go", ToPascalCase(modelName))
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error creating requests file: %v", err)
		return err
	}
	defer file.Close()

	// Execute the template and generate the model
	err = tmpl.Execute(file, modelData)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
		return err
	}

	fmt.Printf("Model %s generated successfully!\n", ToPascalCase(modelName))
	return nil
}

// Fungsi untuk menghasilkan file controller
func generateController(feature string, fields []Field, IsHasTime bool) error {
	// Define the controller data
	modelData := ModelData{
		ModelName: feature,
		Fields:    fields,
		IsHasTime: IsHasTime,
	}

	// Open the template file
	tmpl, err := template.New("controller.tpl").Funcs(template.FuncMap{
		"ToLower":      strings.ToLower,
		"ToPascalCase": ToPascalCase,
		"ToSnakeCase":  ToSnakeCase,
	}).ParseFiles("models/templates/controller.tpl") // ensure template file exists
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
		return err
	}

	// Create the controller file
	fileName := fmt.Sprintf("controllers/%sController.go", ToPascalCase(feature))
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error creating controller file: %v", err)
		return err
	}
	defer file.Close()

	// Execute the template and generate the controller
	err = tmpl.Execute(file, modelData)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
		return err
	}

	fmt.Printf("Controller %s generated successfully!\n", ToPascalCase(feature))
	return nil
}

// Fungsi untuk menghasilkan file route
func generateRoute(feature string, IsHasTime bool) error {

	// Define the controller data
	modelData := ModelData{
		ModelName: feature,
		IsHasTime: IsHasTime,
	}

	// Open the template file
	tmpl, err := template.New("route.tpl").Funcs(template.FuncMap{
		"ToLower":      strings.ToLower,
		"ToPascalCase": ToPascalCase,
		"ToSnakeCase":  ToSnakeCase,
	}).ParseFiles("models/templates/route.tpl") // ensure template file exists
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
		return err
	}

	// Create the controller file
	fileName := fmt.Sprintf("routes/%sRoute.go", ToPascalCase(feature))
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error creating route file: %v", err)
		return err
	}
	defer file.Close()

	// Execute the template and generate the sRoute
	err = tmpl.Execute(file, modelData)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
		return err
	}

	fmt.Printf("Controller %s generated successfully!\n", ToPascalCase(feature))
	return nil
}

// ToPascalCase converts a string to Pascal case
func ToPascalCase(str string) string {
	return strings.ReplaceAll(cases.Title(language.English).String(str), " ", "")
}

// ToSnakeCase converts a string to snake case
func ToSnakeCase(str string) string {
	return strings.ToLower(strings.ReplaceAll(str, " ", "_"))
}
