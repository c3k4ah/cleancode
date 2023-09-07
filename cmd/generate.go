// command to generate scaffold
// Path: cmd/generate.go

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/abdelkarim/cleancode/helpers"
)


var generateCmd = &cobra.Command{
	Use:   "generate",	
	Short: "Generate a new scaffold",
	Long: `Generate a new scaffold`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {	
		res := helpers.createDir(args[0])
		fmt.Println(res)
	},
	}

func init() {
	rootCmd.AddCommand(generateCmd)
}
