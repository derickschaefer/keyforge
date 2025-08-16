package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
    Use:   "config",
    Short: "Manage KeyForge configuration (API keys, models, settings)",
}

var configListCmd = &cobra.Command{
    Use:   "list",
    Short: "List current configuration",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Listing current configuration...")
        // TODO: implement config read
    },
}

var configSetCmd = &cobra.Command{
    Use:   "set [key] [value]",
    Short: "Set a configuration value",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Setting %s to %s\n", args[0], args[1])
        // TODO: implement config write
    },
}

var configTestCmd = &cobra.Command{
    Use:   "test",
    Short: "Test OpenAI connection",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Testing OpenAI connection...")
        // TODO: implement test
    },
}

func init() {
    rootCmd.AddCommand(configCmd)
    configCmd.AddCommand(configListCmd)
    configCmd.AddCommand(configSetCmd)
    configCmd.AddCommand(configTestCmd)
}
