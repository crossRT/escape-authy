/*
Copyright © 2025 crossRT crossrt@gmail.com
*/
package cmd

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/crossRT/escape-authy/helper"
	"github.com/crossRT/escape-authy/model"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// toAegisCmd represents the toAegis command
var toAegisCmd = &cobra.Command{
	Use:   "toAegis",
	Short: "Convert Authy decrypted_tokens.json to aegis format",
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

		var aegisEntries []model.AegisEntry
		for _, token := range dtf.DecryptedAuthenticatorTokens {

			timer := helper.GetTimerByDigits(int(token.Digits))

			issuer := helper.DetermineIssuer(token)

			aegisEntry := model.AegisEntry{
				Type:   "totp",
				UUID:   uuid.New().String(),
				Name:   strings.TrimSpace(token.Name), // use original token.Name as account in order to avoid confusion
				Issuer: issuer,
				Note:   "",
				Icon:   nil,
				Info: model.AegisInfo{
					Secret: token.DecryptedSeed,
					Algo:   "SHA1",
					Digits: token.Digits,
					Period: timer,
				},
			}

			aegisEntries = append(aegisEntries, aegisEntry)
		}

		err = helper.ExportFile(aegisEntries, "aegis.json")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(toAegisCmd)
}
