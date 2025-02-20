package web

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"go-micro.dev/v5/logger"

	"github.com/urfave/cli/v2"
	"go-micro.dev/v5"
	"go-micro.dev/v5/registry"
	"go-micro.dev/v5/web"
)

// Name of Web.
func Name(n string) web.Option {
	return web.Name(n)
}

// Icon specifies an icon url to load in the UI.
func Icon(ico string) web.Option {
	return web.Icon(ico)
}

// Id for Unique server id.
func Id(id string) web.Option {
	return web.Id(id)
}

// Version of the service.
func Version(v string) web.Option {
	return web.Version(v)
}

// Metadata associated with the service.
func Metadata(md map[string]string) web.Option {
	return web.Metadata(md)
}

// Address to bind to - host:port.
func Address(a string) web.Option {
	return web.Address(a)
}

// Advertise The address to advertise for discovery - host:port.
func Advertise(a string) web.Option {
	return web.Advertise(a)
}

// Context specifies a context for the service.
// Can be used to signal shutdown of the service.
// Can be used for extra option values.
func Context(ctx context.Context) web.Option {
	return web.Context(ctx)
}

// Registry used for discovery.
func Registry(r registry.Registry) web.Option {
	return web.Registry(r)
}

// RegisterTTL Register the service with a TTL.
func RegisterTTL(t time.Duration) web.Option {
	return web.RegisterTTL(t)
}

// RegisterInterval Register the service with at interval.
func RegisterInterval(t time.Duration) web.Option {
	return web.RegisterInterval(t)
}

// Handler for custom handler.
func Handler(h http.Handler) web.Option {
	return web.Handler(h)
}

// Server for custom Server.
func Server(srv *http.Server) web.Option {
	return web.Server(srv)
}

// MicroService sets the micro.Service used internally.
func MicroService(s micro.Service) web.Option {
	return web.MicroService(s)
}

// Flags sets the command flags.
func Flags(flags ...cli.Flag) web.Option {
	return web.Flags(flags...)
}

// Action sets the command action.
func Action(a func(*cli.Context)) web.Option {
	return web.Action(a)
}

// BeforeStart is executed before the server starts.
func BeforeStart(fn func() error) web.Option {
	return web.BeforeStart(fn)
}

// BeforeStop is executed before the server stops.
func BeforeStop(fn func() error) web.Option {
	return web.BeforeStop(fn)
}

// AfterStart is executed after server start.
func AfterStart(fn func() error) web.Option {
	return web.AfterStart(fn)
}

// AfterStop is executed after server stop.
func AfterStop(fn func() error) web.Option {
	return web.AfterStop(fn)
}

// Secure Use secure communication.
// If TLSConfig is not specified we use InsecureSkipVerify and generate a self signed cert.
func Secure(b bool) web.Option {
	return web.Secure(b)
}

// TLSConfig to be used for the transport.
func TLSConfig(t *tls.Config) web.Option {
	return web.TLSConfig(t)
}

// StaticDir sets the static file directory. This defaults to ./html.
func StaticDir(d string) web.Option {
	return web.StaticDir(d)
}

// RegisterCheck run func before registry service.
func RegisterCheck(fn func(context.Context) error) web.Option {
	return web.RegisterCheck(fn)
}

// HandleSignal toggles automatic installation of the signal handler that
// traps TERM, INT, and QUIT.  Users of this feature to disable the signal
// handler, should control liveness of the service through the context.
func HandleSignal(b bool) web.Option {
	return web.HandleSignal(b)
}

// Logger sets the underline logger.
func Logger(l logger.Logger) web.Option {
	return web.Logger(l)
}

// Adds a handler Wrapper to a list of options passed into the server.
func WrapHandler(w web.HandlerWrapper) web.Option {
	return web.WrapHandler(w)
}
