package auth

import (
	tokens "github.com/zerospiel/i3-gmail-checker/internal"
)

// GenerateAuthURL return Google OAuth2 URL to auth the app
func GenerateAuthURL() (string, error) {
	tokens.Kek()
	return "kek", nil
}
