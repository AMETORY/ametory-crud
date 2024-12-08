package cmd

import (
	"ametory-crud/database"
	"ametory-crud/models"
	"fmt"
	"os"
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
				Tag:       cases.Lower(language.English).String(strings.ReplaceAll(parts[0], "_", "")),
			}
			fields = append(fields, field)
		}
		switch generateType {
		case "model":

			generateModel(name, fields)
		case "request":

			generateRequestResponse(name, fields)
		case "controller":
			generateController(name, fields)
		case "route":
			generateRoute(name)

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
}

// Field holds details about each field in the model
type Field struct {
	ModelName string
	Name      string
	Type      string
	DBType    string
	Tag       string
}

func generateModel(modelName string, fields []Field) error {

	var crudMethods = [...]string{"create", "read", "update", "delete"}

	for _, crudMethod := range crudMethods {
		// Add permission logic
		permission := models.Permission{
			Name:        fmt.Sprintf("%s %s", cases.Title(language.English).String(modelName), cases.Title(language.English).String(crudMethod)),
			Description: fmt.Sprintf("Permission to %s %s", crudMethod, modelName),
			Key:         fmt.Sprintf("%s:%s", strings.ToLower(crudMethod), strings.ToLower(modelName)),
			Group:       strings.ToUpper(strings.ReplaceAll(modelName, " ", "_")),
		}

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
	}

	// Open the template file
	tmpl, err := template.ParseFiles("models/templates/model.tpl") // ensure template file exists
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
		return err
	}

	// Create the model file
	fileName := fmt.Sprintf("models/%s.go", modelName)
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

	fmt.Printf("Model %s generated successfully!\n", modelName)
	return nil
}
func generateRequestResponse(modelName string, fields []Field) error {
	// Define the model data
	modelData := ModelData{
		ModelName: modelName,
		Fields:    fields,
	}

	// Open the template file
	tmpl, err := template.New("request_response.tpl").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).ParseFiles("models/templates/request_response.tpl") // ensure template file exists
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
		return err
	}

	// Create the model file
	fileName := fmt.Sprintf("requests/%sReq.go", modelName)
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

	fmt.Printf("Model %s generated successfully!\n", modelName)
	return nil
}

// Fungsi untuk menghasilkan file controller
func generateController(feature string, fields []Field) error {
	// Define the controller data
	modelData := ModelData{
		ModelName: feature,
		Fields:    fields,
	}

	// Open the template file
	tmpl, err := template.New("controller.tpl").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).ParseFiles("models/templates/controller.tpl") // ensure template file exists
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
		return err
	}

	// Create the controller file
	fileName := fmt.Sprintf("controllers/%sController.go", feature)
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

	fmt.Printf("Controller %s generated successfully!\n", feature)
	return nil
}

// Fungsi untuk menghasilkan file route
func generateRoute(feature string) error {

	// Define the controller data
	modelData := ModelData{
		ModelName: feature,
	}

	// Open the template file
	tmpl, err := template.New("route.tpl").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).ParseFiles("models/templates/route.tpl") // ensure template file exists
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
		return err
	}

	// Create the controller file
	fileName := fmt.Sprintf("routes/%sRoute.go", feature)
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

	fmt.Printf("Controller %s generated successfully!\n", feature)
	return nil
}
