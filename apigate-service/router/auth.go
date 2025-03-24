package router

import (
	pb "proto/apigate"
	"context"
	"errors"
	"time"

	gt "github.com/golang-jwt/jwt"
	"go-micro.dev/v5/client"
	"go-micro.dev/v5/metadata"
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

func (a *jwt) Generate(ctx context.Context, acc *pb.Account, opts ...client.CallOption) (*pb.Token, error) {
	md, _ := metadata.FromContext(ctx)
	secret := md["secret"]
	host := md["Host"]
	// auth policy by hostname
	switch host {
	case "yourhost":
	default:
		secret = DefaultSecret
	}

	// generate the JWT
	now := time.Now()
	expiry := now
	expirys := time.Duration(0)
	if expirys == 0 {
		nextDay := now.Add(24 * time.Hour)
		expiry = time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 5, 0, 0, 0, nextDay.Location())
	} else {
		expiry = now.Add(expirys)
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
	return &pb.Token{
		AccessToken:  tok,
		RefreshToken: tok,
		Expiry:       expiry.Unix(),
		Created:      now.Unix(),
	}, nil
}

func (a *jwt) Inspect(ctx context.Context, token *pb.Token, opt ...client.CallOption) (*pb.Account, error) {
	md, _ := metadata.FromContext(ctx)
	host := md["Host"]
	acc := pb.Account{}
	switch host {
	case "yourhost":
	default:

	}
	return &acc, nil
}

func (a *jwt) Verify(ctx context.Context, ct *pb.Credential, opts ...client.CallOption) (*pb.Empty, error) {
	md, _ := metadata.FromContext(ctx)
	host := md["Host"]
	switch host {
	case "yourhost":
	default:
	}
	return &pb.Empty{}, nil
}

func (a *jwt) Refresh(ctx context.Context, tk *pb.Token, opts ...client.CallOption) (*pb.Token, error) {
	md, _ := metadata.FromContext(ctx)
	host := md["Host"]
	acc := pb.Token{}
	switch host {
	case "yourhost":
	default:
	}
	return &acc, nil
}

func NewJWT() pb.AuthService {
	return &jwt{}
}
