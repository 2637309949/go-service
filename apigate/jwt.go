package main

import (
	"apigate/router/auth"
	"errors"
	"time"

	gt "github.com/golang-jwt/jwt"
)

var (
	// ErrInvalidToken is when the token provided is not valid
	ErrInvalidToken = errors.New("invalid token provided")
	// ErrForbidden is when a user does not have the necessary scope to access a resource
	ErrForbidden  = errors.New("resource forbidden")
	DefaultSecret = "howh3ogf4q34hf9ogq34f"
)

type authClaims struct {
	Roles []string `json:"roles"`
	Name  string   `json:"name"`
	gt.StandardClaims
}

// jwt just for test scene
// you should write auth as service, so any change would not restart apigate
// also cache
type jwt struct{}

func (a *jwt) Generate(acc *auth.Account, opt ...auth.Option) (*auth.Token, error) {
	opts := auth.Options{}
	secret := ""
	for _, o := range opt {
		o(&opts)
	}

	// auth policy by hostname
	switch opts.Host {
	case "your.hostname":
	default:
		secret = DefaultSecret
	}

	// generate the JWT
	now := time.Now()
	expiry := now
	if opts.Expiry == 0 {
		nextDay := now.Add(24 * time.Hour)
		expiry = time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 5, 0, 0, 0, nextDay.Location())
	} else {
		expiry = now.Add(opts.Expiry)
	}
	t := gt.NewWithClaims(gt.SigningMethodRS256, authClaims{
		Roles: acc.Roles, Name: acc.Name,
		StandardClaims: gt.StandardClaims{
			Id:        acc.Id,
			Issuer:    acc.Issuer,
			ExpiresAt: expiry.Unix(),
		},
	})
	tok, err := t.SignedString(secret)
	if err != nil {
		return nil, err
	}

	// return the token
	return &auth.Token{
		AccessToken:  tok,
		RefreshToken: tok,
		Expiry:       expiry,
		Created:      time.Now(),
	}, nil
}

func (a *jwt) Inspect(token string, opt ...auth.Option) (*auth.Account, error) {
	opts := auth.Options{}
	acc := auth.Account{}
	for _, o := range opt {
		o(&opts)
	}

	// auth policy by hostname
	switch opts.Host {
	case "xxx":
	default:

	}
	// fmt.Printf("%+v\n", opts.Host)
	return &acc, nil
}

func (a *jwt) Verify(acc *auth.Account, scope string, opt ...auth.Option) error {
	opts := auth.Options{}
	for _, o := range opt {
		o(&opts)
	}

	// auth policy by hostname
	switch opts.Host {
	case "xxx":
	default:

	}
	return nil
}

func (a *jwt) Refresh(token string, opt ...auth.Option) (*auth.Token, error) {
	opts := auth.Options{}
	acc := auth.Token{}
	for _, o := range opt {
		o(&opts)
	}

	// auth policy by hostname
	switch opts.Host {
	case "xxx":
	default:

	}
	return &acc, nil
}
