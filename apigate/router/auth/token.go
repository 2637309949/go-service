package auth

import "time"

// Token can be short or long lived
type Token struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Created      time.Time `json:"created"`
	Expiry       time.Time `json:"expiry"`
}

// Expired returns a boolean indicating if the token needs to be refreshed
func (t *Token) Expired() bool {
	return t.Expiry.Unix() < time.Now().Unix()
}
