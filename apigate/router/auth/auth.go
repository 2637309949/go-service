package auth

var (
	// DefaultAuth implementation
	DefaultAuth Auth
)

// Auth provides authentication and authorization
type Auth interface {
	// Generate a new account
	Generate(*Account, ...Option) (*Token, error)
	// Inspect a token
	Inspect(string, ...Option) (*Account, error)
	// Token generated using refresh token or credentials
	Verify(*Account, string, ...Option) error
	// Token generated using refresh token or credentials
	Refresh(string, ...Option) (*Token, error)
}
