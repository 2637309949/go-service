package server

import (
	"reflect"

	"go-micro.dev/v5/registry"
)

type RpcHandler struct {
	handler   interface{}
	opts      HandlerOptions
	name      string
	endpoints []*registry.Endpoint
}

func NewRpcHandler(handler interface{}, opts ...HandlerOption) Handler {
	options := HandlerOptions{
		Metadata: make(map[string]map[string]string),
	}

	for _, o := range opts {
		o(&options)
	}

	typ := reflect.TypeOf(handler)
	hdlr := reflect.ValueOf(handler)
	name := reflect.Indirect(hdlr).Type().Name()

	var endpoints []*registry.Endpoint

	for m := 0; m < typ.NumMethod(); m++ {
		if e := extractEndpoint(typ.Method(m)); e != nil {
			e.Name = name + "." + e.Name
			meta := options.Metadata[e.Name]
			medp := Decode(meta)

			for k, v := range meta {
				e.Metadata[k] = v
			}

			// proto gen for http
			if meta != nil {
				e.Authorization = medp.Authorization
				e.Scope = medp.Scope
				e.Path = medp.Path
				e.Method = medp.Method
				e.Handler = medp.Handler
			}

			endpoints = append(endpoints, e)
		}
	}

	return &RpcHandler{
		name:      name,
		handler:   handler,
		endpoints: endpoints,
		opts:      options,
	}
}

func (r *RpcHandler) Name() string {
	return r.name
}

func (r *RpcHandler) Handler() interface{} {
	return r.handler
}

func (r *RpcHandler) Endpoints() []*registry.Endpoint {
	return r.endpoints
}

func (r *RpcHandler) Options() HandlerOptions {
	return r.opts
}
