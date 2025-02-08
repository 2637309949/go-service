package service

import (
	"comm/logger"
	"context"
	"time"

	"comm/config"

	"github.com/micro/plugins/v5/registry/consul"
	"github.com/micro/plugins/v5/wrapper/trace/opentracing"
	"go-micro.dev/v5"
	"go-micro.dev/v5/registry"
	"go-micro.dev/v5/server"
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
	ctx := context.Background()
	registryAddress := config.CommConf("registry_address")
	opentracingAddress := config.CommConf("opentracing_address")
	registry := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			registryAddress,
		}
	})
	opts = append(opts, micro.WrapClient(wrapClient))
	opts = append(opts, micro.Registry(registry))
	opts = append(opts, micro.WrapClient(opentracing.NewClientWrapper(nil)))
	opts = append(opts, micro.WrapHandler(opentracing.NewHandlerWrapper(nil)))
	opts = append(opts, micro.WrapSubscriber(opentracing.NewSubscriberWrapper(nil)))
	opts = append(opts, micro.WrapHandler(loggerWrapper(logger.DefaultLogger)))
	opts = append(opts, micro.RegisterTTL(time.Second*90))
	opts = append(opts, micro.RegisterInterval(time.Second*30))
	opts = append(opts, micro.Version("latest"))
	opts = append(opts, micro.Context(ctx))
	service := &service{micro.NewService(
		opts...,
	)}
	service.Init()
	service.Server().Init(
		server.Wait(nil),
	)
	initJaegerTracer(service.Name(), opentracingAddress)

	setupService(service)
	return service
}
