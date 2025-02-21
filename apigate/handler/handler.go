package handler

import (
	"apigate/api"
	"apigate/util"
	"net/http"

	"go-micro.dev/v5/errors"
)

var (
	defaultHandler = "rpc"
)

// Handler represents a HTTP handler that manages a request
type Handler interface {
	// standard http handler
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	// name of handler
	String() string
}

type h struct {
	opts Options
}

func (a *h) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var service *api.Service

	if a.opts.Service != nil {
		service = a.opts.Service
	} else if a.opts.Router != nil {
		s, err := a.opts.Router.Route(a.opts.ApiBase, r)
		if err != nil {
			util.WriteError(w, r, errors.InternalServerError("go.micro.api", "%s", err.Error()))
			return
		}
		service = s
	} else {
		// we have no way of routing the request
		util.WriteError(w, r, errors.InternalServerError("go.micro.api", "no route found"))
		return
	}

	handler := defaultHandler
	if len(service.Endpoint.Handler) > 0 {
		handler = service.Endpoint.Handler
	}

	if h, ok := a.opts.Handlers[handler]; ok {
		h.ServeHTTP(w, r)
		return
	}

	util.WriteError(w, r, errors.InternalServerError("go.micro.api", "no handler found"))
}

func (a *h) String() string {
	return "handler"
}

func NewHandler(opts ...Option) Handler {
	options := NewOptions(opts...)
	return &h{
		opts: options,
	}
}
