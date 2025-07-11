/*
Copyright © 2025 crossRT crossrt@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/crossRT/escape-authy/helper"
	"github.com/crossRT/escape-authy/model"
	"github.com/spf13/cobra"
)

// toEnteCmd represents the toEnte command
var toEnteCmd = &cobra.Command{
	Use:   "toEnte",
	Short: "Convert Authy decrypted_tokens.json to ente format",
	Run: func(cmd *cobra.Command, args []string) {
		decryptedTokensFilePath, _ = rootCmd.PersistentFlags().GetString("decrypted-tokens-file-path")

		file, err := os.ReadFile(decryptedTokensFilePath)
		if err != nil {
			panic(err)
		}

		// Parse the JSON
		var dtf model.DecryptedTokensFile
		err = json.Unmarshal(file, &dtf)
		if err != nil {
			panic(err)
		}

		var entries []string
		for _, token := range dtf.DecryptedAuthenticatorTokens {

			issuer, name, err := helper.ExtractPartsFromTokenName(token.Name)

			if err != nil {
				issuer = helper.DetermineIssuer(token)
				name = strings.TrimSpace(token.Name)
			}

			account := fmt.Sprintf("%s:%s", issuer, name)
			if issuer == "" {
				account = name
			}

			line := fmt.Sprintf("otpauth://totp/%s?secret=%s", account, token.DecryptedSeed)
			entries = append(entries, line)
		}

		err = helper.ExportPlainText(entries, "ente.txt")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(toEnteCmd)
}
