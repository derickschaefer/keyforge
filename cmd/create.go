package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
    Use:   "create",
    Short: "Create keys and passwords",
    Long:  "Generate keys or passwords of various types such as easy, strong, or WEP.",
}

var createEasyCmd = &cobra.Command{
    Use:   "easy",
    Short: "Create an easy (memorable) password",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Creating an easy (memorable) password...")
        // TODO: implement logic
    },
}

var createStrongCmd = &cobra.Command{
    Use:   "strong",
    Short: "Create a strong password",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Creating a strong password...")
        // TODO: implement logic
    },
}

var createWepCmd = &cobra.Command{
    Use:   "256wep",
    Short: "Create a 256-bit WEP key",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Creating a 256-bit WEP key...")
        // TODO: implement logic
    },
}

func init() {
    rootCmd.AddCommand(createCmd)
    createCmd.AddCommand(createEasyCmd)
    createCmd.AddCommand(createStrongCmd)
    createCmd.AddCommand(createWepCmd)
}
