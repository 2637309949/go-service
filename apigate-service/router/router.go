package router

import (
	"apigate/api"
	"net/http"
)

// Router is used to determine an endpoint for a request
type Router interface {
	// Returns options
	Options() Options
	// Stop the router
	Close() error
	// Route returns an api.Service route
	Route(string, *http.Request) (*api.Service, error)
	// Init initializes options
	Init(...Option)
}
