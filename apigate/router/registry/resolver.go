package registry

import (
	"net/http"

	"go-micro.dev/v5/registry"
)

// default resolver for legacy purposes
// it uses proxy routing to resolve names
// /foo becomes namespace.foo
// /v1/foo becomes namespace.v1.foo
type apiResolver struct {
	apiBase string
}

func (r *apiResolver) Resolve(req *http.Request) *Endpoint {
	path := req.URL.Path
	method := req.Method
	// get route
	service, endpoint := apiRoute(r.apiBase, path)

	// check for the namespace in the request header, this can be set by the client or injected
	// by the auth wrapper if an auth token was provided. The headr takes priority over any domain
	// passed as a default
	domain := registry.DefaultDomain
	if dom := req.Header.Get("Micro-Namespace"); len(dom) > 0 {
		domain = dom
	}

	return &Endpoint{
		Name:   service,
		Domain: domain,
		Path:   endpoint,
		Method: method,
	}
}

type Endpoint struct {
	Name   string
	Method string
	Domain string
	Path   string
}

func NewResolver(apiBase string) *apiResolver {
	return &apiResolver{apiBase}
}
