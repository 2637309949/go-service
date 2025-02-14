package auth

import (
	"context"
	"time"
)

type Options struct {
	Secret  string
	Host    string
	Context context.Context
	Expiry  time.Duration
}

// Option sets values in Options
type Option func(o *Options)

// WithContext sets the stores context, for any extra configuration
func WithContext(c context.Context) Option {
	return func(o *Options) {
		o.Context = c
	}
}

func Host(host string) Option {
	return func(o *Options) {
		o.Host = host
	}
}

func Secret(secret string) Option {
	return func(o *Options) {
		o.Secret = secret
	}
}

func Expiry(expiry time.Duration) Option {
	return func(o *Options) {
		o.Expiry = expiry
	}
}
