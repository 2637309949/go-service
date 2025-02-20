package web

import (
	"comm/config"
	"comm/logger"
	"context"

	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5/registry"
	"go-micro.dev/v5/web"
)

var (
	ServiceName    string
	defaultService web.Service
)

type service struct {
	web.Service
}

func (s *service) Run() error {
	if err := s.Service.Run(); err != nil {
		logger.Fatalf("Error configuring store table option: %v", err)
		return err
	}
	return nil
}

// NewService creates and returns a new Service based on the packages within.
func NewService(opts ...web.Option) web.Service {
	ctx := context.Background()
	registryAddress := config.CommConf("registry_address")
	opentracingAddress := config.CommConf("opentracing_address")
	registry := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			registryAddress,
		}
	})
	opts = append(opts, Registry(registry))
	opts = append(opts, WrapHandler(NewTracerWrapper(nil)))
	opts = append(opts, Version("latest"))
	opts = append(opts, Context(ctx))
	opts = append(opts, WrapHandler(loggerWrapper(logger.DefaultLogger)))
	opts = append(opts, WrapHandler(loggerHandler))
	service := &service{web.NewService(
		opts...,
	)}

	defaultService = service
	ServiceName = service.Name()

	// init tracer
	initJaegerTracer(service.Name(), opentracingAddress)
	// init conf
	setupService(service)

	return service
}
