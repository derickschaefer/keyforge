package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of KeyForge",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("KeyForge v0.1.0 -- Forged in fire, cooled in entropy")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
