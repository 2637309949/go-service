package auth

// Account provided by an auth provider
type Account struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Roles  []string `json:"scopes"`
	Issuer string   `json:"issuer"`
}
