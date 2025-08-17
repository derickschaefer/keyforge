// cmd/create_set.go
package cmd

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

// SetConfig holds configuration for password set generation
type SetConfig struct {
	Count  int
	AsJSON bool
}

// PasswordSet represents a collection of passwords/keys by type
type PasswordSet struct {
	Easy    []string `json:"easy"`
	Strong  []string `json:"strong"`
	WEP64   []string `json:"64wep"`
	WEP128  []string `json:"128wep"`
	WEP256  []string `json:"256wep"`
}

var createSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Create a set (grid) of passwords/keys",
	Long: `Generate a complete set of different password and key types.
This creates multiple instances of each supported password/key type:
- Easy (memorable) passwords
- Strong (cryptographically secure) passwords  
- WEP keys (64-bit, 128-bit, and 256-bit)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := getSetConfigFromFlags(cmd)
		
		passwordSet, err := generatePasswordSet(cfg)
		if err != nil {
			return fmt.Errorf("failed to generate password set: %w", err)
		}
		
		return printPasswordSet(passwordSet, cfg.AsJSON)
	},
}

// getSetConfigFromFlags extracts configuration from command flags
func getSetConfigFromFlags(cmd *cobra.Command) SetConfig {
	count, _ := cmd.Flags().GetInt("count")
	asJSON, _ := cmd.Flags().GetBool("json")
	
	return SetConfig{
		Count:  count,
		AsJSON: asJSON,
	}
}

// generatePasswordSet creates a complete set of passwords and keys
func generatePasswordSet(cfg SetConfig) (*PasswordSet, error) {
	set := &PasswordSet{}
	
	// Generate easy passwords
	for i := 0; i < cfg.Count; i++ {
		password, err := genEasyWithError(12)
		if err != nil {
			return nil, fmt.Errorf("failed to generate easy password %d: %w", i+1, err)
		}
		set.Easy = append(set.Easy, password)
	}
	
	// Generate strong passwords
	for i := 0; i < cfg.Count; i++ {
		password, err := genStrongWithError(20)
		if err != nil {
			return nil, fmt.Errorf("failed to generate strong password %d: %w", i+1, err)
		}
		set.Strong = append(set.Strong, password)
	}
	
	// Generate 64-bit WEP keys
	for i := 0; i < cfg.Count; i++ {
		key, err := genWEPHexBytesWithError(5) // 5 bytes = 10 hex chars
		if err != nil {
			return nil, fmt.Errorf("failed to generate 64-bit WEP key %d: %w", i+1, err)
		}
		set.WEP64 = append(set.WEP64, key)
	}
	
	// Generate 128-bit WEP keys
	for i := 0; i < cfg.Count; i++ {
		key, err := genWEPHexBytesWithError(13) // 13 bytes = 26 hex chars
		if err != nil {
			return nil, fmt.Errorf("failed to generate 128-bit WEP key %d: %w", i+1, err)
		}
		set.WEP128 = append(set.WEP128, key)
	}
	
	// Generate 256-bit WEP keys
	for i := 0; i < cfg.Count; i++ {
		key, err := genWEPHexBytesWithError(29) // 29 bytes = 58 hex chars
		if err != nil {
			return nil, fmt.Errorf("failed to generate 256-bit WEP key %d: %w", i+1, err)
		}
		set.WEP256 = append(set.WEP256, key)
	}
	
	return set, nil
}

// printPasswordSet outputs the password set in the requested format
func printPasswordSet(set *PasswordSet, asJSON bool) error {
	if asJSON {
		data, err := json.MarshalIndent(set, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal password set to JSON: %w", err)
		}
		fmt.Println(string(data))
		return nil
	}
	
	// Plain text output with consistent ordering
	sections := []struct {
		name     string
		passwords []string
	}{
		{"easy", set.Easy},
		{"strong", set.Strong},
		{"64wep", set.WEP64},
		{"128wep", set.WEP128},
		{"256wep", set.WEP256},
	}
	
	for i, section := range sections {
		fmt.Printf("== %s ==\n", section.name)
		for _, password := range section.passwords {
			fmt.Println(password)
		}
		
		// Add blank line between sections (except after the last one)
		if i < len(sections)-1 {
			fmt.Println()
		}
	}
	
	return nil
}

// Alternative implementation using map for backward compatibility
func generatePasswordSetLegacy(cfg SetConfig) (map[string][]string, error) {
	out := make(map[string][]string)
	
	// Generate easy passwords
	easy := make([]string, cfg.Count)
	for i := 0; i < cfg.Count; i++ {
		password, err := genEasyWithError(12)
		if err != nil {
			return nil, fmt.Errorf("failed to generate easy password %d: %w", i+1, err)
		}
		easy[i] = password
	}
	out["easy"] = easy
	
	// Generate strong passwords
	strong := make([]string, cfg.Count)
	for i := 0; i < cfg.Count; i++ {
		password, err := genStrongWithError(20)
		if err != nil {
			return nil, fmt.Errorf("failed to generate strong password %d: %w", i+1, err)
		}
		strong[i] = password
	}
	out["strong"] = strong
	
	// Generate WEP keys
	wep64 := make([]string, cfg.Count)
	for i := 0; i < cfg.Count; i++ {
		key, err := genWEPHexBytesWithError(5)
		if err != nil {
			return nil, fmt.Errorf("failed to generate 64-bit WEP key %d: %w", i+1, err)
		}
		wep64[i] = key
	}
	out["64wep"] = wep64
	
	wep128 := make([]string, cfg.Count)
	for i := 0; i < cfg.Count; i++ {
		key, err := genWEPHexBytesWithError(13)
		if err != nil {
			return nil, fmt.Errorf("failed to generate 128-bit WEP key %d: %w", i+1, err)
		}
		wep128[i] = key
	}
	out["128wep"] = wep128
	
	wep256 := make([]string, cfg.Count)
	for i := 0; i < cfg.Count; i++ {
		key, err := genWEPHexBytesWithError(29)
		if err != nil {
			return nil, fmt.Errorf("failed to generate 256-bit WEP key %d: %w", i+1, err)
		}
		wep256[i] = key
	}
	out["256wep"] = wep256
	
	return out, nil
}

// printPasswordSetLegacy outputs using the legacy map format
func printPasswordSetLegacy(out map[string][]string, asJSON bool) error {
	if asJSON {
		data, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal password set to JSON: %w", err)
		}
		fmt.Println(string(data))
		return nil
	}
	
	// Sort keys for consistent output
	keys := make([]string, 0, len(out))
	for k := range out {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	
	for i, key := range keys {
		fmt.Printf("== %s ==\n", key)
		for _, password := range out[key] {
			fmt.Println(password)
		}
		
		// Add blank line between sections (except after the last one)
		if i < len(keys)-1 {
			fmt.Println()
		}
	}
	
	return nil
}

func init() {
	createCmd.AddCommand(createSetCmd)
	
	// Add flags
	createSetCmd.Flags().IntP("count", "c", 4, "number of passwords/keys to generate for each type")
	createSetCmd.Flags().Bool("json", false, "output as JSON")
}
