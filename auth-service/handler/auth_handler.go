package handler

import (
	"comm/mark"
	"context"
	pbApigate "proto/apigate"
	"time"

	gt "github.com/golang-jwt/jwt"
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/metadata"
)

// Generate generates a JWT based on account information and context metadata.
// This function implements the logic to generate a token, including obtaining secrets, setting token expiration, and generating the token.
func (h *Handler) Generate(ctx context.Context, acc *pbApigate.Account, rsp *pbApigate.Token) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Generate")()
	secret := ""
	host := ""
	md, ok := metadata.FromContext(ctx)
	if ok {
		secret = md["secret"]
		host = md["Host"]
	}
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
		logger.Errorf("SignedString fail[%+v]", err)
		return err
	}

	// return the token
	rsp.AccessToken = tok
	rsp.RefreshToken = tok
	rsp.Expiry = expiry.Unix()
	rsp.Created = now.Unix()
	return nil
}

// Inspect checks the provided Token and returns the corresponding Account information.
func (h *Handler) Inspect(ctx context.Context, req *pbApigate.Token, rsp *pbApigate.Account) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Inspect")()
	logger.Infof("Received Handler.Get request: %+v", req)
	host := ""
	md, ok := metadata.FromContext(ctx)
	if ok {
		host = md["Host"]
	}
	switch host {
	case "yourhost":
	default:

	}
	return nil
}

// Verify verifies the provided user credential.
// This method is responsible for validating the user's credential information.
func (h *Handler) Verify(ctx context.Context, req *pbApigate.Credential, rsp *pbApigate.Empty) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Verify")()
	logger.Infof("Received Handler.Get request: %+v", req)
	host := ""
	md, ok := metadata.FromContext(ctx)
	if ok {
		host = md["Host"]
	}
	switch host {
	case "yourhost":
	default:
	}
	return nil
}

// Refresh refreshes the authentication token.
// This method takes a context (ctx), a request containing the token to be refreshed (req),
// and a response object (rsp) where the refreshed token will be stored.
// It returns an error if any occurs during the refresh process.
func (h *Handler) Refresh(ctx context.Context, req *pbApigate.Token, rsp *pbApigate.Token) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Refresh")()
	logger.Infof("Received Handler.Get request: %+v", req)
	host := ""
	md, ok := metadata.FromContext(ctx)
	if ok {
		host = md["Host"]
	}
	switch host {
	case "yourhost":
	default:
	}
	return nil
}
