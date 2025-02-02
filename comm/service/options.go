package service

import (
	"context"
	"time"

	"github.com/urfave/cli/v2"
	"go-micro.dev/v5"
	"go-micro.dev/v5/auth"
	"go-micro.dev/v5/broker"
	"go-micro.dev/v5/cache"
	"go-micro.dev/v5/client"
	"go-micro.dev/v5/cmd"
	"go-micro.dev/v5/config"
	"go-micro.dev/v5/debug/profile"
	"go-micro.dev/v5/debug/trace"
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/registry"
	"go-micro.dev/v5/selector"
	"go-micro.dev/v5/server"
	"go-micro.dev/v5/store"
	"go-micro.dev/v5/transport"
)

// Broker to be used for service.
func Broker(b broker.Broker) micro.Option {
	return func(o *micro.Options) {
		o.Broker = b
		// Update Client and Server
		o.Client.Init(client.Broker(b))
		o.Server.Init(server.Broker(b))
	}
}

func Cache(c cache.Cache) micro.Option {
	return func(o *micro.Options) {
		o.Cache = c
	}
}

func Cmd(c cmd.Cmd) micro.Option {
	return func(o *micro.Options) {
		o.Cmd = c
	}
}

// Client to be used for service.
func Client(c client.Client) micro.Option {
	return func(o *micro.Options) {
		o.Client = c
	}
}

// Context specifies a context for the service.
// Can be used to signal shutdown of the service and for extra option values.
func Context(ctx context.Context) micro.Option {
	return func(o *micro.Options) {
		o.Context = ctx
	}
}

// Handle will register a handler without any fuss
func Handle(v interface{}) micro.Option {
	return func(o *micro.Options) {
		o.Server.Handle(
			o.Server.NewHandler(v),
		)
	}
}

// HandleSignal toggles automatic installation of the signal handler that
// traps TERM, INT, and QUIT.  Users of this feature to disable the signal
// handler, should control liveness of the service through the context.
func HandleSignal(b bool) micro.Option {
	return func(o *micro.Options) {
		o.Signal = b
	}
}

// Profile to be used for debug profile.
func Profile(p profile.Profile) micro.Option {
	return func(o *micro.Options) {
		o.Profile = p
	}
}

// Server to be used for service.
func Server(s server.Server) micro.Option {
	return func(o *micro.Options) {
		o.Server = s
	}
}

// Store sets the store to use.
func Store(s store.Store) micro.Option {
	return func(o *micro.Options) {
		o.Store = s
	}
}

// Registry sets the registry for the service
// and the underlying components.
func Registry(r registry.Registry) micro.Option {
	return func(o *micro.Options) {
		o.Registry = r
		// Update Client and Server
		o.Client.Init(client.Registry(r))
		o.Server.Init(server.Registry(r))
		// Update Broker
		o.Broker.Init(broker.Registry(r))
	}
}

// Tracer sets the tracer for the service.
func Tracer(t trace.Tracer) micro.Option {
	return func(o *micro.Options) {
		o.Server.Init(server.Tracer(t))
	}
}

// Auth sets the auth for the service.
func Auth(a auth.Auth) micro.Option {
	return func(o *micro.Options) {
		o.Auth = a
	}
}

// Config sets the config for the service.
func Config(c config.Config) micro.Option {
	return func(o *micro.Options) {
		o.Config = c
	}
}

// Selector sets the selector for the service client.
func Selector(s selector.Selector) micro.Option {
	return func(o *micro.Options) {
		o.Client.Init(client.Selector(s))
	}
}

// Transport sets the transport for the service
// and the underlying components.
func Transport(t transport.Transport) micro.Option {
	return func(o *micro.Options) {
		o.Transport = t
		// Update Client and Server
		o.Client.Init(client.Transport(t))
		o.Server.Init(server.Transport(t))
	}
}

// Convenience options

// Address sets the address of the server.
func Address(addr string) micro.Option {
	return func(o *micro.Options) {
		o.Server.Init(server.Address(addr))
	}
}

// Name of the service.
func Name(n string) micro.Option {
	return func(o *micro.Options) {
		o.Server.Init(server.Name(n))
	}
}

// Version of the service.
func Version(v string) micro.Option {
	return func(o *micro.Options) {
		o.Server.Init(server.Version(v))
	}
}

// Metadata associated with the service.
func Metadata(md map[string]string) micro.Option {
	return func(o *micro.Options) {
		o.Server.Init(server.Metadata(md))
	}
}

// Flags that can be passed to service.
func Flags(flags ...cli.Flag) micro.Option {
	return func(o *micro.Options) {
		o.Cmd.App().Flags = append(o.Cmd.App().Flags, flags...)
	}
}

// Action can be used to parse user provided cli options.
func Action(a func(*cli.Context) error) micro.Option {
	return func(o *micro.Options) {
		o.Cmd.App().Action = a
	}
}

// RegisterTTL specifies the TTL to use when registering the service.
func RegisterTTL(t time.Duration) micro.Option {
	return func(o *micro.Options) {
		o.Server.Init(server.RegisterTTL(t))
	}
}

// RegisterInterval specifies the interval on which to re-register.
func RegisterInterval(t time.Duration) micro.Option {
	return func(o *micro.Options) {
		o.Server.Init(server.RegisterInterval(t))
	}
}

// WrapClient is a convenience method for wrapping a Client with
// some middleware component. A list of wrappers can be provided.
// Wrappers are applied in reverse order so the last is executed first.
func WrapClient(w ...client.Wrapper) micro.Option {
	return func(o *micro.Options) {
		// apply in reverse
		for i := len(w); i > 0; i-- {
			o.Client = w[i-1](o.Client)
		}
	}
}

// WrapCall is a convenience method for wrapping a Client CallFunc.
func WrapCall(w ...client.CallWrapper) micro.Option {
	return func(o *micro.Options) {
		o.Client.Init(client.WrapCall(w...))
	}
}

// WrapHandler adds a handler Wrapper to a list of options passed into the server.
func WrapHandler(w ...server.HandlerWrapper) micro.Option {
	return func(o *micro.Options) {
		var wrappers []server.Option

		for _, wrap := range w {
			wrappers = append(wrappers, server.WrapHandler(wrap))
		}

		// Init once
		o.Server.Init(wrappers...)
	}
}

// WrapSubscriber adds a subscriber Wrapper to a list of options passed into the server.
func WrapSubscriber(w ...server.SubscriberWrapper) micro.Option {
	return func(o *micro.Options) {
		var wrappers []server.Option

		for _, wrap := range w {
			wrappers = append(wrappers, server.WrapSubscriber(wrap))
		}

		// Init once
		o.Server.Init(wrappers...)
	}
}

// Add opt to server option.
func AddListenOption(option server.Option) micro.Option {
	return func(o *micro.Options) {
		o.Server.Init(option)
	}
}

// Before and Afters

// BeforeStart run funcs before service starts.
func BeforeStart(fn func() error) micro.Option {
	return func(o *micro.Options) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

// BeforeStop run funcs before service stops.
func BeforeStop(fn func() error) micro.Option {
	return func(o *micro.Options) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}

// AfterStart run funcs after service starts.
func AfterStart(fn func() error) micro.Option {
	return func(o *micro.Options) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

// AfterStop run funcs after service stops.
func AfterStop(fn func() error) micro.Option {
	return func(o *micro.Options) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}

// Logger sets the logger for the service.
func Logger(l logger.Logger) micro.Option {
	return func(o *micro.Options) {
		o.Logger = l
	}
}
