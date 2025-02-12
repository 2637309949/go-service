package registry

import (
	"apigate/api"
	"apigate/router"
	"errors"
	"net/http"
	"path"
	"regexp"

	util "apigate/util/router"

	"go-micro.dev/v5/registry"
	"go-micro.dev/v5/registry/cache"
	"go-micro.dev/v5/server"
)

var (
	errNotFound = errors.New("not found")
)

// endpoint struct, that holds compiled pcre
type endpoint struct {
	hostregs []*regexp.Regexp
	pathregs []util.Pattern
	pcreregs []*regexp.Regexp
}

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

func (r *registryRouter) Route(req *http.Request) (*api.Service, error) {
	if r.isClosed() {
		return nil, errors.New("router closed")
	}

	// resolve service
	rp := r.resolver.Resolve(req)
	name := rp.Name
	reqPath := rp.Path
	reqMethod := rp.Method
	if len(name) == 0 {
		return nil, errors.New("error during resolve: service not resolved")
	}

	// get service
	services, err := r.rc.GetService(name, registry.GetDomain(rp.Domain))
	if err != nil {
		return nil, err
	}

	// filter httpapi
	endpointName := ""
	httpServices := []*registry.Service{}
	for i := range services {
		service := services[i]
		for j := range service.Endpoints {
			endpoint := service.Endpoints[j]
			httpEndpoint := server.Decode(endpoint.Metadata)
			if httpEndpoint == nil {
				continue
			}
			httpEndpoint.Path = path.Clean(httpEndpoint.Path)
			if len(httpEndpoint.Path) == 0 {
				continue
			}

			if httpEndpoint.Method != reqMethod || httpEndpoint.Path != reqPath {
				continue
			}
			httpServices = append(httpServices, service)
			if len(endpointName) == 0 {
				endpointName = httpEndpoint.Name
			}
		}
	}
	if len(httpServices) == 0 || len(endpointName) == 0 {
		return nil, errors.New("rpc: can't find service Endpoint")
	}

	// construct api service
	return &api.Service{
		Name: name,
		Endpoint: &api.Endpoint{
			Name:    endpointName,
			Handler: "rpc",
		},
		Services: services,
	}, nil
}

func newRouter(opts ...router.Option) *registryRouter {
	options := router.NewOptions(opts...)
	r := &registryRouter{
		exit:     make(chan bool),
		opts:     options,
		rc:       cache.New(options.Registry),
		resolver: NewResolver(options.ApiBase),
	}
	return r
}

// NewRouter returns the default router
func NewRouter(opts ...router.Option) router.Router {
	return newRouter(opts...)
}
