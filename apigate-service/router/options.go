package router

import (
	pb "proto/apigate"

	"go-micro.dev/v5/registry"
)

type Options struct {
	Registry registry.Registry
	Auth     pb.AuthService
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	var options Options

	for _, o := range opts {
		o(&options)
	}

	return options
}

func WithRegistry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

func WithAuth(auth pb.AuthService) Option {
	return func(o *Options) {
		o.Auth = auth
	}
}
