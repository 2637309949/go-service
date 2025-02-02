package service

import (
	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5"
)

// NewService creates and returns a new Service based on the packages within.
func NewService(opts ...micro.Option) micro.Service {
	registry := consul.NewRegistry()
	opts = append(opts, micro.Registry(registry))
	service := micro.NewService(
		opts...,
	)
	service.Init()
	return service
}
