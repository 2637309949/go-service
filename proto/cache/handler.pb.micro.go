// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/cache/handler.proto

package cache

import (
	fmt "fmt"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "go-micro.dev/v5/client"
	registry "go-micro.dev/v5/registry"
	server "go-micro.dev/v5/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ registry.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CacheService service

func NewCacheServiceEndpoints() []*registry.Endpoint {
	return []*registry.Endpoint{
		{
			Authorization: false,
			Name:          "CacheService.Get",
			Path:          "/Get",
			Method:        "GET",
			Handler:       "rpc",
		},
	}
}

// Client API for CacheService service

type CacheService interface {
	// Get gets a cached value by key.
	Get(ctx context.Context, in *CacheFilter, opts ...client.CallOption) (*Cache, error)
	// Put stores a key-value pair into cache.
	Put(ctx context.Context, in *Cache, opts ...client.CallOption) (*Cache, error)
	// Delete removes a key from cache.
	Delete(ctx context.Context, in *CacheFilter, opts ...client.CallOption) (*Cache, error)
}

type cacheService struct {
	c    client.Client
	name string
}

func NewCacheService(name string, c client.Client) CacheService {
	return &cacheService{
		c:    c,
		name: name,
	}
}

func (c *cacheService) Get(ctx context.Context, in *CacheFilter, opts ...client.CallOption) (*Cache, error) {
	req := c.c.NewRequest(c.name, "CacheService.Get", in)
	out := new(Cache)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheService) Put(ctx context.Context, in *Cache, opts ...client.CallOption) (*Cache, error) {
	req := c.c.NewRequest(c.name, "CacheService.Put", in)
	out := new(Cache)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheService) Delete(ctx context.Context, in *CacheFilter, opts ...client.CallOption) (*Cache, error) {
	req := c.c.NewRequest(c.name, "CacheService.Delete", in)
	out := new(Cache)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CacheService service

type CacheServiceHandler interface {
	// Get gets a cached value by key.
	Get(context.Context, *CacheFilter, *Cache) error
	// Put stores a key-value pair into cache.
	Put(context.Context, *Cache, *Cache) error
	// Delete removes a key from cache.
	Delete(context.Context, *CacheFilter, *Cache) error
}

func RegisterCacheServiceHandler(s server.Server, hdlr CacheServiceHandler, opts ...server.HandlerOption) error {
	type cacheService interface {
		Get(ctx context.Context, in *CacheFilter, out *Cache) error
		Put(ctx context.Context, in *Cache, out *Cache) error
		Delete(ctx context.Context, in *CacheFilter, out *Cache) error
	}
	type CacheService struct {
		cacheService
	}
	h := &cacheServiceHandler{hdlr}
	opts = append(opts, server.WithEndpoint(&registry.Endpoint{
		Authorization: false,
		Name:          "CacheService.Get",
		Path:          "/Get",
		Method:        "GET",
		Handler:       "rpc",
	}))
	return s.Handle(s.NewHandler(&CacheService{h}, opts...))
}

type cacheServiceHandler struct {
	CacheServiceHandler
}

func (h *cacheServiceHandler) Get(ctx context.Context, in *CacheFilter, out *Cache) error {
	return h.CacheServiceHandler.Get(ctx, in, out)
}

func (h *cacheServiceHandler) Put(ctx context.Context, in *Cache, out *Cache) error {
	return h.CacheServiceHandler.Put(ctx, in, out)
}

func (h *cacheServiceHandler) Delete(ctx context.Context, in *CacheFilter, out *Cache) error {
	return h.CacheServiceHandler.Delete(ctx, in, out)
}
