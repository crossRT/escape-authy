package model

type DecryptedTokensFile struct {
	Message string `json:"message"`
	Success bool   `json:"success"`

	DecryptedAuthenticatorTokens []DecryptedToken `json:"decrypted_authenticator_tokens"`
}
