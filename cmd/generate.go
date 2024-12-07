package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var generateCmd = &cobra.Command{
	Use:   "generate [type] [name]",
	Short: "Generate a new model, controller, or route",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		generateType := strings.ToLower(args[0])
		name := cases.Title(language.English).String(args[1])
		switch generateType {
		case "model":
			generateModel(name, []string{})
		case "controller":
			generateController(name)
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
func generateModel(feature string, colDefs []string) {
	// Menyiapkan template untuk model
	tmpl, err := template.New("model").ParseFiles("models/templates/model.tpl")
	if err != nil {
		fmt.Printf("Error parsing model template: %v\n", err)
		return
	}

	// Menyiapkan data untuk template
	data := struct {
		Feature string
		Columns []string
	}{
		Feature: feature,
		Columns: colDefs,
	}

	// Menyimpan hasil render template ke file
	fileName := fmt.Sprintf("models/%s.go", feature)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating model file: %v\n", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error executing model template: %v\n", err)
		return
	}

	fmt.Printf("Model file created: %s\n", fileName)
}

// Fungsi untuk menghasilkan file controller
func generateController(feature string) {
	// Menyiapkan template untuk controller
	tmpl, err := template.New("controller").ParseFiles("models/templates/controller.tpl")
	if err != nil {
		fmt.Printf("Error parsing controller template: %v\n", err)
		return
	}

	// Menyiapkan data untuk template
	data := struct {
		Feature string
	}{
		Feature: feature,
	}

	// Menyimpan hasil render template ke file
	fileName := fmt.Sprintf("controllers/%s_controller.go", feature)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating controller file: %v\n", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error executing controller template: %v\n", err)
		return
	}

	fmt.Printf("Controller file created: %s\n", fileName)
}

// Fungsi untuk menghasilkan file route
func generateRoute(feature string) {
	// Menyiapkan template untuk route
	tmpl, err := template.New("route").ParseFiles("models/templates/route.tpl")
	if err != nil {
		fmt.Printf("Error parsing route template: %v\n", err)
		return
	}

	// Menyiapkan data untuk template
	data := struct {
		Feature string
	}{
		Feature: feature,
	}

	// Menyimpan hasil render template ke file
	fileName := fmt.Sprintf("routes/%s_route.go", feature)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating route file: %v\n", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error executing route template: %v\n", err)
		return
	}

	fmt.Printf("Route file created: %s\n", fileName)
}
