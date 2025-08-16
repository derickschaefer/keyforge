package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "keyforge",
	Short: "KeyForge - Keys and Passwords Forged with Strength",
	Long: `KeyForge is a CLI tool for generating and analyzing keys and passwords.
    
Commands follow a verb-noun pattern:
- create: generate a single key or password
- create set: generate a set of keys/passwords
- analyze: evaluate a password
- config: manage models, API keys, and settings
- help: usage information
`,
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
