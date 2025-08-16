package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Create a set of keys/passwords",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating a set of keys/passwords (like RandomKeygen)...")
		// TODO: implement logic
	},
}

func init() {
	createCmd.AddCommand(createSetCmd)
}
