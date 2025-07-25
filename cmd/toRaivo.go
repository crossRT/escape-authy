/*
Copyright © 2025 crossRT crossrt@gmail.com
*/
package cmd

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/crossRT/escape-authy/helper"
	"github.com/crossRT/escape-authy/model"
	"github.com/spf13/cobra"
)

// toRaivoCmd represents the toRaivo command
var toRaivoCmd = &cobra.Command{
	Use:   "toRaivo",
	Short: "Convert Authy decrypted_tokens.json to raivo format",
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

		var raivoEntries []model.RaivoEntry
		for _, token := range dtf.DecryptedAuthenticatorTokens {

			timer := helper.GetTimerByDigits(int(token.Digits))

			issuer := helper.DetermineIssuer(token)

			raivoEntry := model.RaivoEntry{
				Kind:      "TOTP",
				Account:   strings.TrimSpace(token.Name), // use original token.Name as account in order to avoid confusion
				Secret:    token.DecryptedSeed,
				IconType:  "raivo_repository",
				Issuer:    issuer,
				Timer:     strconv.Itoa(timer),
				Digits:    strconv.Itoa(token.Digits),
				Counter:   "0",
				Algorithm: "SHA1",
				IconValue: nil,
			}

			raivoEntries = append(raivoEntries, raivoEntry)
		}

		err = helper.ExportFile(raivoEntries, "raivo.json")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(toRaivoCmd)
}
