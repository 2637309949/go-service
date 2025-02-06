package server

import (
	"context"
	"strings"
	"go-micro.dev/v5/registry"
)

type HandlerOption func(*HandlerOptions)

type HandlerOptions struct {
	Metadata map[string]map[string]string
	Internal bool
}

type SubscriberOption func(*SubscriberOptions)

type SubscriberOptions struct {
	Context context.Context
	Queue   string
	// AutoAck defaults to true. When a handler returns
	// with a nil error the message is acked.
	AutoAck  bool
	Internal bool
}


// Encode encodes an endpoint to endpoint metadata
func Encode(e *registry.Endpoint) map[string]string {
	if e == nil {
		return nil
	}

	// endpoint map
	ep := make(map[string]string)

	// set vals only if they exist
	set := func(k, v string) {
		if len(v) == 0 {
			return
		}
		ep[k] = v
	}

	set("endpoint", e.Name)
	set("handler", e.Handler)
	set("method", strings.Join(e.Method, ","))
	set("path", e.Path)

	return ep
}

// Decode decodes endpoint metadata into an endpoint
func Decode(e map[string]string) *registry.Endpoint {
	if e == nil {
		return nil
	}

	return &registry.Endpoint{
		Name:        e["endpoint"],
		Method:      slice(e["method"]),
		Path:        e["path"],
		Handler:     e["handler"],
	}
}

func slice(s string) []string {
	var sl []string

	for _, p := range strings.Split(s, ",") {
		if str := strings.TrimSpace(p); len(str) > 0 {
			sl = append(sl, strings.TrimSpace(p))
		}
	}

	return sl
}

func WithEndpoint(e *registry.Endpoint) HandlerOption {
	return EndpointMetadata(e.Name, Encode(e))
}

// EndpointMetadata is a Handler option that allows metadata to be added to
// individual endpoints.
func EndpointMetadata(name string, md map[string]string) HandlerOption {
	return func(o *HandlerOptions) {
		o.Metadata[name] = md
	}
}

// Internal Handler options specifies that a handler is not advertised
// to the discovery system. In the future this may also limit request
// to the internal network or authorized user.
func InternalHandler(b bool) HandlerOption {
	return func(o *HandlerOptions) {
		o.Internal = b
	}
}

// Internal Subscriber options specifies that a subscriber is not advertised
// to the discovery system.
func InternalSubscriber(b bool) SubscriberOption {
	return func(o *SubscriberOptions) {
		o.Internal = b
	}
}
func NewSubscriberOptions(opts ...SubscriberOption) SubscriberOptions {
	opt := SubscriberOptions{
		AutoAck: true,
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// DisableAutoAck will disable auto acking of messages
// after they have been handled.
func DisableAutoAck() SubscriberOption {
	return func(o *SubscriberOptions) {
		o.AutoAck = false
	}
}

// Shared queue name distributed messages across subscribers.
func SubscriberQueue(n string) SubscriberOption {
	return func(o *SubscriberOptions) {
		o.Queue = n
	}
}

// SubscriberContext set context options to allow broker SubscriberOption passed.
func SubscriberContext(ctx context.Context) SubscriberOption {
	return func(o *SubscriberOptions) {
		o.Context = ctx
	}
}
