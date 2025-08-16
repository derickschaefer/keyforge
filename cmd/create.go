// cmd/create.go
package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/spf13/cobra"
)

var (
	length int
	count  int
	asJSON bool
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create keys and passwords",
	Long:  "Generate keys/passwords in various styles (easy, strong, 256wep, ...).",
}

var createEasyCmd = &cobra.Command{
	Use:   "easy",
	Short: "Create an easy (memorable) password",
	RunE: func(cmd *cobra.Command, args []string) error {
		results := make([]string, 0, count)
		for i := 0; i < count; i++ {
			results = append(results, genEasy(length))
		}
		return printResults(results, asJSON)
	},
}

var createStrongCmd = &cobra.Command{
	Use:   "strong",
	Short: "Create a strong password",
	RunE: func(cmd *cobra.Command, args []string) error {
		results := make([]string, 0, count)
		for i := 0; i < count; i++ {
			results = append(results, genStrong(length))
		}
		return printResults(results, asJSON)
	},
}

var createWEP64Cmd = &cobra.Command{
	Use:   "64wep",
	Short: "Create a 64-bit WEP key (40-bit key, 10 hex chars)",
	RunE: func(cmd *cobra.Command, args []string) error {
		results := make([]string, 0, count)
		for i := 0; i < count; i++ {
			results = append(results, genWEPHexBytes(5)) // 5 bytes -> 10 hex chars
		}
		return printResults(results, asJSON)
	},
}

var createWEP128Cmd = &cobra.Command{
	Use:   "128wep",
	Short: "Create a 128-bit WEP key (104-bit key, 26 hex chars)",
	RunE: func(cmd *cobra.Command, args []string) error {
		results := make([]string, 0, count)
		for i := 0; i < count; i++ {
			results = append(results, genWEPHexBytes(13)) // 13 bytes -> 26 hex chars
		}
		return printResults(results, asJSON)
	},
}

var createWEP256Cmd = &cobra.Command{
	Use:   "256wep",
	Short: "Create a 256-bit WEP key (232-bit key, 58 hex chars)",
	RunE: func(cmd *cobra.Command, args []string) error {
		results := make([]string, 0, count)
		for i := 0; i < count; i++ {
			results = append(results, genWEPHexBytes(29)) // 29 bytes -> 58 hex chars
		}
		return printResults(results, asJSON)
	},
}

func genWEPHexBytes(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		// extremely unlikely; fallback to empty string on error
		return ""
	}
	return hex.EncodeToString(b)
}

func init() {
	rootCmd.AddCommand(createCmd)

	// easy
	createCmd.AddCommand(createEasyCmd)
	createEasyCmd.Flags().IntVarP(&length, "length", "l", 12, "length of the password")
	createEasyCmd.Flags().IntVarP(&count, "count", "c", 1, "how many to generate")
	createEasyCmd.Flags().BoolVar(&asJSON, "json", false, "output as JSON")

	// strong
	createCmd.AddCommand(createStrongCmd)
	createStrongCmd.Flags().IntVarP(&length, "length", "l", 20, "length of the password")
	createStrongCmd.Flags().IntVarP(&count, "count", "c", 1, "how many to generate")
	createStrongCmd.Flags().BoolVar(&asJSON, "json", false, "output as JSON")

	// WEP 64
	createCmd.AddCommand(createWEP64Cmd)
	createWEP64Cmd.Flags().IntVarP(&count, "count", "c", 1, "how many to generate")
	createWEP64Cmd.Flags().BoolVar(&asJSON, "json", false, "output as JSON")

	// WEP 128
	createCmd.AddCommand(createWEP128Cmd)
	createWEP128Cmd.Flags().IntVarP(&count, "count", "c", 1, "how many to generate")
	createWEP128Cmd.Flags().BoolVar(&asJSON, "json", false, "output as JSON")

	// WEP 256
	createCmd.AddCommand(createWEP256Cmd)
	createWEP256Cmd.Flags().IntVarP(&count, "count", "c", 1, "how many to generate")
	createWEP256Cmd.Flags().BoolVar(&asJSON, "json", false, "output as JSON")
}

// ---- generators ----

func genEasy(n int) string {
	if n < 4 {
		n = 4
	}
	// Vowels/consonants to produce pronounceable-ish strings + digits
	const vowels = "aeiou"
	const cons = "bcdfghjklmnpqrstvwxyz"
	const digits = "0123456789"

	var b strings.Builder
	for i := 0; i < n; i++ {
		var pool string
		switch {
		case i%3 == 2:
			pool = digits
		case i%2 == 0:
			pool = cons
		default:
			pool = vowels
		}
		ch := randChoice(pool)
		b.WriteByte(ch)
	}
	return b.String()
}

func genStrong(n int) string {
	if n < 8 {
		n = 8
	}
	const pool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{};:,.<>/?~"
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteByte(randChoice(pool))
	}
	return b.String()
}

func randChoice(pool string) byte {
	max := big.NewInt(int64(len(pool)))
	ri, _ := rand.Int(rand.Reader, max)
	return pool[ri.Int64()]
}

func printResults(results []string, jsonOut bool) error {
	if jsonOut {
		data, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		return nil
	}

	for _, r := range results {
		fmt.Println(r)
	}
	return nil
}

// cmdOutWriter returns stdout now; later you can swap for buffers/tests.
func cmdOutWriter() *strings.Builder {
	// for now we just print directly in printResults; keeping placeholder if needed later
	return &strings.Builder{}
}
