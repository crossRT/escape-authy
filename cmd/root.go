/*
Copyright © 2025 crossRT crossrt@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var decryptedTokensFilePath string

var rootCmd = &cobra.Command{
	Use:   "escape-authy",
	Short: "A cli tool to convert decrypted_tokens.json to preferred MFA application format",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&decryptedTokensFilePath, "decrypted-tokens-file-path", "./decrypted_tokens.json", "File path of the decrypted_tokens.json")
}
