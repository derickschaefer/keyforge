// cmd/create_set.go
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var setAsJSON bool

var createSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Create a set (grid) of passwords/keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		out := map[string][]string{
			"easy":    {genEasy(12), genEasy(12), genEasy(12), genEasy(12)},
			"strong":  {genStrong(20), genStrong(20), genStrong(20), genStrong(20)},
			"64wep":   {genWEPHexBytes(5), genWEPHexBytes(5), genWEPHexBytes(5), genWEPHexBytes(5)},
			"128wep":  {genWEPHexBytes(13), genWEPHexBytes(13), genWEPHexBytes(13), genWEPHexBytes(13)},
			"256wep":  {genWEPHexBytes(29), genWEPHexBytes(29), genWEPHexBytes(29), genWEPHexBytes(29)},
		}
		if setAsJSON {
			b, _ := json.MarshalIndent(out, "", "  ")
			fmt.Println(string(b))
			return nil
		}
		for k, v := range out {
			fmt.Println("==", k, "==")
			for _, s := range v {
				fmt.Println(s)
			}
			fmt.Println()
		}
		return nil
	},
}

func init() {
	createCmd.AddCommand(createSetCmd)
	createSetCmd.Flags().BoolVar(&setAsJSON, "json", false, "output as JSON")
}
