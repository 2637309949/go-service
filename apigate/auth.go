package main

import (
	"context"
	"errors"
	"net/http"
	"net/textproto"
	"strings"

	"apigate/api/router/registry"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/metadata"
)

var (
	// ErrInvalidToken is when the token provided is not valid
	ErrInvalidToken = errors.New("invalid token provided")
	// ErrForbidden is when a user does not have the necessary scope to access a resource
	ErrForbidden = errors.New("resource forbidden")
)

const (
	// BearerScheme used for Authorization header
	BearerScheme = "Bearer "
	// TokenCookieName is the name of the cookie which stores the auth token
	TokenCookieName = "micro-token"
)

// Account provided by an auth provider
type Account struct {
	ID string `json:"id"`
	// Name of the account. User friendly name that might change e.g. a username or email
	Name string `json:"name"`
	// Any other associated metadata
	Metadata map[string]string `json:"metadata"`
	// Scopes the account has access to
	Scopes []string `json:"scopes"`
	// Secret for the account, e.g. the password
	Secret string `json:"secret"`
}

// Resource is an entity such as a user or
type Resource struct {
	// Name of the resource, e.g. go.micro.service.notes
	Name string `json:"name"`
	// Type of resource, e.g. service
	Type string `json:"type"`
	// Endpoint resource e.g NotesService.Create
	Endpoint string `json:"endpoint"`
}

type VerifyOptions struct {
	Context   context.Context
	Namespace string
}

type VerifyOption func(o *VerifyOptions)

func VerifyContext(ctx context.Context) VerifyOption {
	return func(o *VerifyOptions) {
		o.Context = ctx
	}
}
func VerifyNamespace(ns string) VerifyOption {
	return func(o *VerifyOptions) {
		o.Namespace = ns
	}
}

func auth() gin.HandlerFunc {
	resolver := registry.NewResolver()
	return func(c *gin.Context) {
		// Determine the name of the service being requested
		req := c.Request
		endpoint := resolver.Resolve(req)
		ctrx := context.WithValue(req.Context(), registry.Endpoint{}, endpoint)
		*req = *req.Clone(ctrx)

		// Set the metadata so we can access it in micro api / web
		req = req.WithContext(FromRequest(req))

		// Extract the token from the request
		var token string
		if header := req.Header.Get("Authorization"); len(header) > 0 {
			// Extract the auth token from the request
			if strings.HasPrefix(header, BearerScheme) {
				token = header[len(BearerScheme):]
			}
		}

		// Get the account using the token, some are unauthenticated, so the lack of an
		// account doesn't necessarily mean a forbidden request
		acc, err := Inspect(token)
		if err == nil {
			// inject into the context
			ctx := ContextWithAccount(req.Context(), acc)
			*req = *req.Clone(ctx)
		}

		// construct the resource name, e.g. home => foo.api.home
		resName := endpoint.Name
		resEndpoint := endpoint.Method

		logger.Debugf("Resolving %v %v", resName, resEndpoint)

		// Perform the verification check to see if the account has access to
		// the resource they're requesting
		res := &Resource{Type: "service", Name: resName, Endpoint: resEndpoint}
		if err := Verify(acc, res); err == nil {
			// The account has the necessary permissions to access the resource
			c.Next()
			return
		} else if err != ErrForbidden {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		// The account is set, but they don't have enough permissions, hence
		// we return a forbidden error.
		if acc != nil {
			http.Error(c.Writer, "Forbidden request", http.StatusForbidden)
			return
		}

		http.Error(c.Writer, "unauthorized request", http.StatusUnauthorized)
	}
}

// Verify an account has access to a resource using the rules
func Verify(acc *Account, res *Resource, opts ...VerifyOption) error {
	return nil
}

// Inspect a token
func Inspect(token string) (*Account, error) {
	return &Account{}, nil
}

type accountKey struct{}

// AccountFromContext gets the account from the context, which
// is set by the auth wrapper at the start of a call. If the account
// is not set, a nil account will be returned. The error is only returned
// when there was a problem retrieving an account
func AccountFromContext(ctx context.Context) (*Account, bool) {
	acc, ok := ctx.Value(accountKey{}).(*Account)
	return acc, ok
}

// ContextWithAccount sets the account in the context
func ContextWithAccount(ctx context.Context, account *Account) context.Context {
	return context.WithValue(ctx, accountKey{}, account)
}

func FromRequest(r *http.Request) context.Context {
	ctx := r.Context()
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(metadata.Metadata)
	}
	for k, v := range r.Header {
		md[textproto.CanonicalMIMEHeaderKey(k)] = strings.Join(v, ",")
	}
	// pass http host
	md["Host"] = r.Host
	// pass http method
	md["Method"] = r.Method
	if r.URL != nil {
		md["URL"] = r.URL.String()
	}
	return metadata.NewContext(ctx, md)
}
