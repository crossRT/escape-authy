package model

type RaivoEntry struct {
	Kind      string  `json:"kind"`
	Account   string  `json:"account"`
	Secret    string  `json:"secret"`
	IconType  string  `json:"iconType"`
	Issuer    string  `json:"issuer"`
	Timer     string  `json:"timer"`
	Digits    string  `json:"digits"`
	Counter   string  `json:"counter"`
	Algorithm string  `json:"algorithm"`
	IconValue *string `json:"iconValue"`
}
