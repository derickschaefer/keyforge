package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze [password]",
	Short: "Analyze a password",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		password := args[0]
		fmt.Printf("Analyzing password: %s\n", password)
		// TODO: entropy check + AI hook
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
