// command to generate scaffold
// Path: cmd/generate.go

package cmd

import (
	"cleancode/helpers"
	"fmt"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new scaffold",
	Long:  `Generate a new scaffold`,
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		helpers.CreateDir(args[0])
		fmt.Println("Scaffold generated successfully")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
