package service

import (
	"comm/logger"

	"github.com/micro/plugins/v5/registry/consul"
	"github.com/micro/plugins/v5/wrapper/trace/opentracing"
	"go-micro.dev/v5"
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
	registry := consul.NewRegistry()
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
