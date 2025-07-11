package helper

import (
	"errors"
	"strings"

	"github.com/crossRT/escape-authy/model"
)

func SafeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func GetTimerByDigits(d int) int {
	if d == 7 {
		return 10
	}

	return 30
}

// not sure why authy issuer sometimes is null
// thus we need to determine issuer conditionally
func DetermineIssuer(token model.DecryptedToken) string {
	issuer := SafeString(token.Issuer)

	if token.Issuer == nil && token.AccountType == "google" {
		issuer = "Google"
	}

	if token.Issuer == nil && token.AccountType == "digitalocean" {
		issuer = "DigitalOcean"
	}

	if token.Issuer == nil && token.Name == "luno" {
		issuer = "Luno"
	}

	if token.Issuer == nil && token.Name == "bybit" {
		issuer = "Bybit"
	}

	// if issuer is still empty here, try extract from token.Name from format "XXX: example@abc.com"
	parts := strings.SplitN(token.Name, ":", 2)
	if len(parts) == 2 {
		issuer = strings.TrimSpace(parts[0])
	}

	return issuer
}

func ExtractPartsFromTokenName(str string) (issuer string, name string, err error) {
	parts := strings.SplitN(str, ":", 2)

	if len(parts) != 2 {
		return "", "", errors.New("input does not contain ':' separator")
	}

	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), nil
}
