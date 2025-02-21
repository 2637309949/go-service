package registry

import (
	"apigate/api"
	"apigate/router"
	"apigate/router/auth"
	"errors"
	"net/http"
	"path"

	"go-micro.dev/v5/registry"
	"go-micro.dev/v5/registry/cache"
)

// router is the default router
type registryRouter struct {
	exit chan bool
	opts router.Options
	// registry cache
	rc cache.Cache

	resolver *apiResolver
}

func (r *registryRouter) isClosed() bool {
	select {
	case <-r.exit:
		return true
	default:
		return false
	}
}

func (r *registryRouter) Options() router.Options {
	return r.opts
}

func (r *registryRouter) Close() error {
	select {
	case <-r.exit:
		return nil
	default:
		close(r.exit)
		r.rc.Stop()
	}
	return nil
}

func (r *registryRouter) Register(ep *api.Endpoint) error {
	return nil
}

func (r *registryRouter) Deregister(ep *api.Endpoint) error {
	return nil
}

func (r *registryRouter) Route(apiBase string, req *http.Request) (*api.Service, error) {
	if r.isClosed() {
		return nil, errors.New("router closed")
	}
	// resolve service
	rp := r.resolver.Resolve(apiBase, req)
	if len(rp.Name) == 0 {
		return nil, errors.New("error during resolve: service not resolved")
	}

	// get service
	services, err := r.rc.GetService(rp.Name, registry.GetDomain(rp.Domain))
	if err != nil {
		return nil, err
	}

	// route match
	sv, err := r.match(rp, services)
	if err != nil {
		return nil, err
	}

	// auth match
	err = r.auth(req, rp)
	if err != nil {
		return nil, err
	}

	return sv, err
}

func (r *registryRouter) auth(req *http.Request, rp *Endpoint) error {
	cx := req.Context()

	if rp.Authorization {
		if r.opts.Auth == nil {
			return errors.New("no auth policy found")
		}
		acc, err := r.opts.Auth.Inspect(rp.Token, auth.WithContext(cx), auth.Host(req.Host))
		if err != nil {
			return err
		}
		if len(rp.Scope) > 0 {
			err = r.opts.Auth.Verify(acc, rp.Scope, auth.WithContext(cx), auth.Host(req.Host))
			if err != nil {
				return err
			}
		}
		// set context
	}

	return nil
}

func (r *registryRouter) match(rp *Endpoint, services []*registry.Service) (*api.Service, error) {
	sv := api.Service{
		Name:     rp.Name,
		Endpoint: &api.Endpoint{},
		Services: services,
	}

	authorization := false
	scope := ""

	for i := range services {
		service := services[i]
		for j := range service.Endpoints {
			endpoint := service.Endpoints[j]
			endpoint.Path = path.Clean(endpoint.Path)

			if len(endpoint.Path) == 0 {
				continue
			}

			if endpoint.Method != rp.Method || endpoint.Path != rp.Path {
				continue
			}

			if endpoint.Authorization {
				authorization = true
			}

			if len(scope) == 0 {
				scope = endpoint.Scope
			}

			if len(sv.Endpoint.Name) == 0 {
				sv.Endpoint.Name = endpoint.Name
			}

			if len(sv.Endpoint.Handler) == 0 {
				sv.Endpoint.Handler = endpoint.Handler
			}

			sv.Services = append(sv.Services, service)
		}
	}

	rp.Authorization = authorization
	rp.Scope = scope

	if len(sv.Services) == 0 || len(sv.Endpoint.Name) == 0 {
		return nil, errors.New("rpc: can't find service Endpoint")
	}

	// construct api service
	return &sv, nil
}

func newRouter(opts ...router.Option) *registryRouter {
	options := router.NewOptions(opts...)
	r := &registryRouter{
		exit:     make(chan bool),
		opts:     options,
		rc:       cache.New(options.Registry),
		resolver: NewResolver(),
	}
	return r
}

// NewRouter returns the default router
func NewRouter(opts ...router.Option) router.Router {
	return newRouter(opts...)
}
