package model

type DecryptedToken struct {
	AccountType   string  `json:"account_type"`
	Name          string  `json:"name"`
	Issuer        *string `json:"issuer"`
	DecryptedSeed string  `json:"decrypted_seed"`
	Digits        int     `json:"digits"`
	Logo          *string `json:"logo"`
	UniqueID      string  `json:"unique_id"`
}
