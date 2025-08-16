// cmd/config.go
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration (API keys, model, settings)",
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List current configuration (redacts secrets)",
	RunE: func(cmd *cobra.Command, args []string) error {
		loadConfig()
		apiKey := viper.GetString("openai_api_key")
		model := viper.GetString("model")
		if apiKey != "" && len(apiKey) > 6 {
			apiKey = apiKey[:3] + "…" + apiKey[len(apiKey)-3:]
		}
		fmt.Printf("config file: %s\n", viper.ConfigFileUsed())
		fmt.Printf("model:       %s\n", model)
		fmt.Printf("openai key:  %s\n", apiKey)
		return nil
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a configuration value (model|openai_api_key)",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		loadConfig()
		viper.Set(args[0], args[1])
		return saveConfig()
	},
}

var configTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Sanity-check config presence (no external calls yet)",
	RunE: func(cmd *cobra.Command, args []string) error {
		loadConfig()
		model := viper.GetString("model")
		key := viper.GetString("openai_api_key")
		if model == "" {
			fmt.Println("model: (not set)")
		} else {
			fmt.Println("model:", model)
		}
		if key == "" {
			fmt.Println("openai_api_key: (not set)")
		} else {
			fmt.Println("openai_api_key: present")
		}
		fmt.Println("OK")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configTestCmd)

	// allow custom config path later: keyforge --config /path/to/file config list
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.keyforge.yaml)")
}

func loadConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, _ := os.UserHomeDir()
		viper.AddConfigPath(home)
		viper.SetConfigName(".keyforge")
		viper.SetConfigType("yaml")
	}
	// env overrides
	viper.SetEnvPrefix("keyforge")
	viper.AutomaticEnv() // KEYFORGE_MODEL, KEYFORGE_OPENAI_API_KEY

	_ = viper.ReadInConfig() // ignore missing file
}

func saveConfig() error {
	path := viper.ConfigFileUsed()
	if path == "" {
		home, _ := os.UserHomeDir()
		path = filepath.Join(home, ".keyforge.yaml")
		viper.SetConfigFile(path)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return viper.WriteConfig()
}
