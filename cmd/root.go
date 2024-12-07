package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gen",
	Short: "CLI to generate CRUD apps with Gin Framework",
}

func Execute() error {
	return rootCmd.Execute()
}
