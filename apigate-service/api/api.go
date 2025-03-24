package api

import (
	"go-micro.dev/v5/registry"
)

// Endpoint is a mapping between an RPC method and HTTP endpoint
type Endpoint struct {
	// RPC Method e.g. Greeter.Hello
	Name string
	// Domain
	Domain string
	// Description e.g what's this endpoint for
	Description string
	// API Handler e.g rpc, proxy
	Handler string
	// HTTP Host e.g example.com
	Host []string
	// HTTP Methods e.g GET, POST
	Method []string
	// HTTP Path e.g /greeter. Expect POSIX regex
	Path string
	// Body destination
	// "*" or "" - top level message value
	// "string" - inner message value
	Body string
	// Stream flag
	Stream bool
}

// Service represents an API service
type Service struct {
	// Name of service
	Name string
	// Domain of the service
	Domain string
	// The endpoint for this service
	Endpoint *Endpoint
	// Versions of this service
	Services []*registry.Service
}
