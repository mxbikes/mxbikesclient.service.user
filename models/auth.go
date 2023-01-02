package models

import "time"

type AuthResponse struct {
	AccessToken string    `json:"access_token"`
	Scope       string    `json:"scope"`
	ExpiresIn   int       `json:"expires_in"`
	ExpiresAt   time.Time `json:"expires_at"`
	TokenType   string    `json:"token_type"`
}
