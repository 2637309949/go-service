package service

import (
	"comm/logger"

	"github.com/micro/plugins/v5/registry/consul"
	sourceConsul "github.com/micro/plugins/v5/source/consul"
	"github.com/micro/plugins/v5/wrapper/trace/opentracing"
	"go-micro.dev/v5"
	"go-micro.dev/v5/config"
	"go-micro.dev/v5/config/source/env"
	"go-micro.dev/v5/registry"
)

type service struct {
	micro.Service
}

func (s *service) Run() error {
	if err := s.Service.Run(); err != nil {
		logger.Fatalf("Error configuring store table option: %v", err)
		return err
	}
	return nil
}

// NewService creates and returns a new Service based on the packages within.
func NewService(opts ...micro.Option) micro.Service {
	err := config.Load(env.NewSource(
		env.WithPath("./.env"),
	))
	if err != nil {
		logger.Fatalf("Error source load: %v", err)
	}
	consulAddress := config.Get("consul").String("")
	err = config.Load(sourceConsul.NewSource(
		sourceConsul.WithAddress(consulAddress),
		sourceConsul.WithPrefix("/micro/config"),
		sourceConsul.StripPrefix(false),
	))
	if err != nil {
		logger.Fatalf("Error source load: %v", err)
	}
	registryAddress := config.Get("micro", "config", "comm", "registry_address").String("")
	opentracingAddress := config.Get("micro", "config", "comm", "opentracing_address").String("")
	_ = opentracingAddress
	registry := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			registryAddress,
		}
	})
	opts = append(opts, micro.Registry(registry))
	opts = append(opts, micro.WrapClient(opentracing.NewClientWrapper(nil)))
	opts = append(opts, micro.WrapHandler(opentracing.NewHandlerWrapper(nil)))
	opts = append(opts, micro.WrapSubscriber(opentracing.NewSubscriberWrapper(nil)))
	opts = append(opts, micro.WrapHandler(LoggerWrapper(logger.DefaultLogger)))
	service := &service{micro.NewService(
		opts...,
	)}
	service.Init()
	initJaegerTracer(service.Name(), "")
	return service
}
