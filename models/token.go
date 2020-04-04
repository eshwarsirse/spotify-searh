package models

import "time"

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type,omitempty"`
	Scope       string `json:"scope,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	Expiry      time.Time
}

func (t *Token) IsExpired() bool {
	return t.Expiry.After(time.Now())
}
