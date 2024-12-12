package cmd

import (
	"ametory-crud/services"
	"fmt"

	"github.com/spf13/cobra"
)

var geminiCmd = &cobra.Command{
	Use:   "gemini [prompt]",
	Short: "prompt gemini",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
		res, err := services.GeminiPrompt(args[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	},
}
