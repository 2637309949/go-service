package handler

import (
	"apigate/api"
	"apigate/router"

	"go-micro.dev/v5/client"
)

var (
	DefaultMaxRecvSize int64 = 1024 * 1024 * 100 // 10Mb
)

type Options struct {
	MaxRecvSize int64
	Namespace   string
	Router      router.Router
	Client      client.Client
	Service     *api.Service
	Handlers    map[string]Handler
	ApiBase     string
}

type Option func(o *Options)

// NewOptions fills in the blanks
func NewOptions(opts ...Option) Options {
	var options Options
	options.Handlers = map[string]Handler{}
	for _, o := range opts {
		o(&options)
	}

	if options.Client == nil {
		WithClient(client.DefaultClient)(&options)
	}

	// set namespace if blank
	if len(options.Namespace) == 0 {
		WithNamespace("micro")(&options)
	}

	if options.MaxRecvSize == 0 {
		options.MaxRecvSize = DefaultMaxRecvSize
	}

	return options
}

// WithNamespace specifies the namespace for the handler
func WithNamespace(s string) Option {
	return func(o *Options) {
		o.Namespace = s
	}
}

// WithRouter specifies a router to be used by the handler
func WithRouter(r router.Router) Option {
	return func(o *Options) {
		o.Router = r
	}
}

func WithClient(c client.Client) Option {
	return func(o *Options) {
		o.Client = c
	}
}

// WithMaxRecvSize specifies max body size
func WithMaxRecvSize(size int64) Option {
	return func(o *Options) {
		o.MaxRecvSize = size
	}
}

func WithWrapCall(cw ...client.CallWrapper) Option {
	return func(o *Options) {
		if o.Client == nil {
			WithClient(client.DefaultClient)(o)
		}
		o.Client.Init(client.WrapCall(cw...))
	}
}

func WithService(s *api.Service) Option {
	return func(o *Options) {
		o.Service = s
	}
}

func WithHandler(h Handler) Option {
	return func(o *Options) {
		o.Handlers[h.String()] = h
	}
}

func WithApiBase(base string) Option {
	return func(o *Options) {
		o.ApiBase = base
	}
}
