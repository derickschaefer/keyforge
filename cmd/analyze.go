// cmd/analyze.go
package cmd

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var fromStdin bool

var analyzeCmd = &cobra.Command{
	Use:   "analyze [password]",
	Short: "Analyze a password (offline entropy/heuristics; AI later)",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var pwd string
		if fromStdin {
			info, _ := os.Stdin.Stat()
			if (info.Mode() & os.ModeCharDevice) != 0 {
				return fmt.Errorf("--stdin provided but no piped input")
			}
			s := bufio.NewScanner(os.Stdin)
			if s.Scan() {
				pwd = strings.TrimSpace(s.Text())
			}
		} else if len(args) == 1 {
			pwd = args[0]
		} else {
			return fmt.Errorf("provide a password or use --stdin")
		}

		report := analyzePassword(pwd)
		fmt.Println(report)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().BoolVar(&fromStdin, "stdin", false, "read password from STDIN")
}

func analyzePassword(p string) string {
	length := len(p)
	classes := charClasses(p)
	entropy := shannonEntropy(p)

	var verdict string
	switch {
	case length >= 16 && classes >= 3 && entropy >= 3.5:
		verdict = "Strong"
	case length >= 12 && classes >= 2 && entropy >= 3.0:
		verdict = "Moderate"
	default:
		verdict = "Weak"
	}

	warnings := []string{}
	if hasKeyboardRun(p) {
		warnings = append(warnings, "Avoid keyboard runs (e.g., qwerty, asdf).")
	}
	if looksLikeDate(p) {
		warnings = append(warnings, "Avoid dates or year patterns.")
	}
	if repeats(p) {
		warnings = append(warnings, "Avoid repeated characters/sequences.")
	}

	// Don’t echo the password; show hash tail so users can compare runs privately.
	hash := sha256.Sum256([]byte(p))
	hashTail := fmt.Sprintf("%x", hash)[56:]

	var b strings.Builder
	fmt.Fprintf(&b, "Length: %d\n", length)
	fmt.Fprintf(&b, "Classes: %d (lower/upper/digit/symbol)\n", classes)
	fmt.Fprintf(&b, "Entropy: %.2f bits/char\n", entropy)
	fmt.Fprintf(&b, "Verdict: %s\n", verdict)
	if len(warnings) > 0 {
		fmt.Fprintf(&b, "Warnings:\n  - %s\n", strings.Join(warnings, "\n  - "))
	}
	fmt.Fprintf(&b, "Reference: sha256(...)=...%s\n", hashTail)
	return b.String()
}

func shannonEntropy(s string) float64 {
	if s == "" {
		return 0
	}
	freq := map[rune]float64{}
	for _, r := range s {
		freq[r]++
	}
	l := float64(len([]rune(s)))
	h := 0.0
	for _, c := range freq {
		p := c / l
		h += -p * math.Log2(p)
	}
	return h
}

func charClasses(s string) int {
	hasL, hasU, hasD, hasS := false, false, false, false
	for _, r := range s {
		switch {
		case unicode.IsLower(r):
			hasL = true
		case unicode.IsUpper(r):
			hasU = true
		case unicode.IsDigit(r):
			hasD = true
		default:
			hasS = true
		}
	}
	n := 0
	for _, v := range []bool{hasL, hasU, hasD, hasS} {
		if v {
			n++
		}
	}
	return n
}

func hasKeyboardRun(s string) bool {
	ss := strings.ToLower(s)
	runs := []string{"qwerty", "asdf", "zxcv", "12345", "0987"}
	for _, r := range runs {
		if strings.Contains(ss, r) {
			return true
		}
	}
	return false
}
func looksLikeDate(s string) bool {
	ss := strings.ToLower(s)
	return strings.Contains(ss, "2020") || strings.Contains(ss, "2021") ||
		strings.Contains(ss, "2022") || strings.Contains(ss, "2023") ||
		strings.Contains(ss, "2024") || strings.Contains(ss, "2025")
}
func repeats(s string) bool {
	if len(s) < 3 {
		return false
	}
	return strings.Contains(s, strings.Repeat(string(s[0]), 3))
}
