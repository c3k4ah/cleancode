// command to generate scaffold
// Path: cmd/generate.go

package cmd

import (
	"cleancode/helpers"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new scaffold",
	Long:  `Generate a new scaffold`,
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		helpers.CreateAllDirs(args[0])
		color.Blue(" All directories created successfully")
		helpers.CreateAllFiles(args[0])
		color.Blue(" All files created successfully")
		color.Green(" Scaffold generated successfully")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
