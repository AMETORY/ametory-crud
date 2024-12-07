package cmd

import (
	db "ametory-crud/database"
	mdl "ametory-crud/models"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		db.ConnectDatabase()
		mdl.MigrateDatabase()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
