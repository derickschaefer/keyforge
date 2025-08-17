// cmd/create.go
package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Config holds password generation parameters
type Config struct {
	Length int
	Count  int
	AsJSON bool
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create keys and passwords",
	Long:  "Generate keys/passwords in various styles (easy, strong, WEP keys).",
}

var createEasyCmd = &cobra.Command{
	Use:   "easy",
	Short: "Create an easy (memorable) password",
	Long:  "Generate memorable passwords using alternating consonants, vowels, and digits",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := getConfigFromFlags(cmd)
		results := make([]string, 0, cfg.Count)
		
		for i := 0; i < cfg.Count; i++ {
			password, err := genEasyWithError(cfg.Length)
			if err != nil {
				return fmt.Errorf("failed to generate easy password: %w", err)
			}
			results = append(results, password)
		}
		return printResults(results, cfg.AsJSON)
	},
}

var createStrongCmd = &cobra.Command{
	Use:   "strong",
	Short: "Create a strong password",
	Long:  "Generate cryptographically strong passwords using mixed character sets",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := getConfigFromFlags(cmd)
		results := make([]string, 0, cfg.Count)
		
		for i := 0; i < cfg.Count; i++ {
			password, err := genStrongWithError(cfg.Length)
			if err != nil {
				return fmt.Errorf("failed to generate strong password: %w", err)
			}
			results = append(results, password)
		}
		return printResults(results, cfg.AsJSON)
	},
}

var createWEP64Cmd = &cobra.Command{
	Use:   "64wep",
	Short: "Create a 64-bit WEP key (40-bit key, 10 hex chars)",
	Long:  "Generate a 64-bit WEP key with 40 bits of actual key material (5 bytes = 10 hex characters)",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := getConfigFromFlags(cmd)
		results := make([]string, 0, cfg.Count)
		
		for i := 0; i < cfg.Count; i++ {
			key, err := genWEPHexBytesWithError(5) // 5 bytes -> 10 hex chars
			if err != nil {
				return fmt.Errorf("failed to generate 64-bit WEP key: %w", err)
			}
			results = append(results, key)
		}
		return printResults(results, cfg.AsJSON)
	},
}

var createWEP128Cmd = &cobra.Command{
	Use:   "128wep",
	Short: "Create a 128-bit WEP key (104-bit key, 26 hex chars)",
	Long:  "Generate a 128-bit WEP key with 104 bits of actual key material (13 bytes = 26 hex characters)",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := getConfigFromFlags(cmd)
		results := make([]string, 0, cfg.Count)
		
		for i := 0; i < cfg.Count; i++ {
			key, err := genWEPHexBytesWithError(13) // 13 bytes -> 26 hex chars
			if err != nil {
				return fmt.Errorf("failed to generate 128-bit WEP key: %w", err)
			}
			results = append(results, key)
		}
		return printResults(results, cfg.AsJSON)
	},
}

var createWEP256Cmd = &cobra.Command{
	Use:   "256wep",
	Short: "Create a 256-bit WEP key (232-bit key, 58 hex chars)",
	Long:  "Generate a 256-bit WEP key with 232 bits of actual key material (29 bytes = 58 hex characters)",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := getConfigFromFlags(cmd)
		results := make([]string, 0, cfg.Count)
		
		for i := 0; i < cfg.Count; i++ {
			key, err := genWEPHexBytesWithError(29) // 29 bytes -> 58 hex chars
			if err != nil {
				return fmt.Errorf("failed to generate 256-bit WEP key: %w", err)
			}
			results = append(results, key)
		}
		return printResults(results, cfg.AsJSON)
	},
}

// getConfigFromFlags extracts configuration from command flags
func getConfigFromFlags(cmd *cobra.Command) Config {
	length, _ := cmd.Flags().GetInt("length")
	count, _ := cmd.Flags().GetInt("count")
	asJSON, _ := cmd.Flags().GetBool("json")
	
	return Config{
		Length: length,
		Count:  count,
		AsJSON: asJSON,
	}
}

// genWEPHexBytes generates n random bytes and returns them as a hex string
func genWEPHexBytes(n int) string {
	result, err := genWEPHexBytesWithError(n)
	if err != nil {
		// Log error but don't panic - return empty string as fallback
		// This matches the original behavior
		return ""
	}
	return result
}

// genWEPHexBytesWithError is the internal version that returns errors
func genWEPHexBytesWithError(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("byte count must be positive, got: %d", n)
	}
	
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to read random bytes: %w", err)
	}
	return hex.EncodeToString(b), nil
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Add subcommands
	createCmd.AddCommand(createEasyCmd)
	createCmd.AddCommand(createStrongCmd)
	createCmd.AddCommand(createWEP64Cmd)
	createCmd.AddCommand(createWEP128Cmd)
	createCmd.AddCommand(createWEP256Cmd)

	// Easy password flags
	createEasyCmd.Flags().IntP("length", "l", 12, "length of the password (minimum 4)")
	createEasyCmd.Flags().IntP("count", "c", 1, "number of passwords to generate")
	createEasyCmd.Flags().Bool("json", false, "output as JSON array")

	// Strong password flags
	createStrongCmd.Flags().IntP("length", "l", 20, "length of the password (minimum 8)")
	createStrongCmd.Flags().IntP("count", "c", 1, "number of passwords to generate")
	createStrongCmd.Flags().Bool("json", false, "output as JSON array")

	// WEP key flags (no length needed as they're fixed size)
	createWEP64Cmd.Flags().IntP("count", "c", 1, "number of keys to generate")
	createWEP64Cmd.Flags().Bool("json", false, "output as JSON array")

	createWEP128Cmd.Flags().IntP("count", "c", 1, "number of keys to generate")
	createWEP128Cmd.Flags().Bool("json", false, "output as JSON array")

	createWEP256Cmd.Flags().IntP("count", "c", 1, "number of keys to generate")
	createWEP256Cmd.Flags().Bool("json", false, "output as JSON array")
}

// ---- Password Generators ----

// genEasy generates a memorable password using alternating consonants, vowels, and digits
func genEasy(n int) string {
	result, err := genEasyWithError(n)
	if err != nil {
		// Log error but don't panic - return a fallback
		fmt.Fprintf(os.Stderr, "Warning: Error generating easy password: %v\n", err)
		return genEasyFallback(n)
	}
	return result
}

// genEasyWithError is the internal version that returns errors
func genEasyWithError(n int) (string, error) {
	if n < 4 {
		n = 4
	}
	
	// Character pools for pronounceable passwords
	const vowels = "aeiou"
	const consonants = "bcdfghjklmnpqrstvwxyz"
	const digits = "0123456789"

	var b strings.Builder
	b.Grow(n) // Pre-allocate capacity
	
	for i := 0; i < n; i++ {
		var pool string
		switch {
		case i%3 == 2: // Every third character is a digit
			pool = digits
		case i%2 == 0: // Even positions get consonants
			pool = consonants
		default: // Odd positions get vowels
			pool = vowels
		}
		
		ch, err := randChoice(pool)
		if err != nil {
			return "", fmt.Errorf("failed to select random character: %w", err)
		}
		b.WriteByte(ch)
	}
	return b.String(), nil
}

// genStrong generates a cryptographically strong password using mixed character sets
func genStrong(n int) string {
	result, err := genStrongWithError(n)
	if err != nil {
		// Log error but don't panic - return a fallback
		fmt.Fprintf(os.Stderr, "Warning: Error generating strong password: %v\n", err)
		return genStrongFallback(n)
	}
	return result
}

// genStrongWithError is the internal version that returns errors
func genStrongWithError(n int) (string, error) {
	if n < 8 {
		n = 8
	}
	
	const pool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{};:,.<>/?~"
	
	var b strings.Builder
	b.Grow(n) // Pre-allocate capacity
	
	for i := 0; i < n; i++ {
		ch, err := randChoice(pool)
		if err != nil {
			return "", fmt.Errorf("failed to select random character: %w", err)
		}
		b.WriteByte(ch)
	}
	return b.String(), nil
}

// Fallback functions for when crypto/rand fails (extremely unlikely)
func genEasyFallback(n int) string {
	if n < 4 {
		n = 4
	}
	const vowels = "aeiou"
	const consonants = "bcdfghjklmnpqrstvwxyz"
	const digits = "0123456789"

	var b strings.Builder
	b.Grow(n)
	
	for i := 0; i < n; i++ {
		var pool string
		switch {
		case i%3 == 2:
			pool = digits
		case i%2 == 0:
			pool = consonants
		default:
			pool = vowels
		}
		// Use a simple modulo fallback (not cryptographically secure)
		b.WriteByte(pool[i%len(pool)])
	}
	return b.String()
}

func genStrongFallback(n int) string {
	if n < 8 {
		n = 8
	}
	const pool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{};:,.<>/?~"
	
	var b strings.Builder
	b.Grow(n)
	
	for i := 0; i < n; i++ {
		// Use a simple modulo fallback (not cryptographically secure)
		b.WriteByte(pool[i%len(pool)])
	}
	return b.String()
}

// randChoice selects a random character from the given pool
func randChoice(pool string) (byte, error) {
	if len(pool) == 0 {
		return 0, fmt.Errorf("character pool cannot be empty")
	}
	
	max := big.NewInt(int64(len(pool)))
	ri, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, fmt.Errorf("failed to generate random number: %w", err)
	}
	return pool[ri.Int64()], nil
}

// printResults outputs the results either as JSON or plain text
func printResults(results []string, jsonOut bool) error {
	if len(results) == 0 {
		return fmt.Errorf("no results to print")
	}
	
	if jsonOut {
		data, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
		fmt.Println(string(data))
		return nil
	}

	// Plain text output
	for _, result := range results {
		fmt.Println(result)
	}
	return nil
}
