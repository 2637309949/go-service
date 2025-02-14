package auth

import "time"

// Account provided by an auth provider
type Account struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Roles []string `json:"scopes"`
	// 区分是那个系统的签发
	Issuer string `json:"issuer"`
}

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
