package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename Go.mod",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(args)
		if err := renameFiles(args[0]); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
	// renameCmd.Flags().String("module-name", "", "Name of the module to rename to")
}

func renameFiles(moduleName string) error {
	var filePaths []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			filePaths = append(filePaths, path)
		}
		return nil
	})

	filePaths = append(filePaths, "go.mod")
	filePaths = append(filePaths, "go.sum")
	if err != nil {
		return err
	}
	for _, filePath := range filePaths {
		input, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		output := strings.ReplaceAll(string(input), "ametory-crud", moduleName)
		if err := os.WriteFile(filePath, []byte(output), 0644); err != nil {
			return err
		}

		// fmt.Println(output)
	}
	return nil
}
