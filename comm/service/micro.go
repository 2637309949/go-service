package service

import (
	"comm/db"
	"comm/logger"
	"context"
	"encoding/base64"
	"encoding/json"
	"time"

	"comm/config"

	"github.com/forgoer/openssl"
	"github.com/micro/plugins/v5/registry/consul"
	"github.com/micro/plugins/v5/wrapper/trace/opentracing"
	"github.com/micro/plugins/v5/wrapper/validator"
	"go-micro.dev/v5"
	"go-micro.dev/v5/registry"
	"go-micro.dev/v5/server"
)

var (
	dbdriver       = "mysql"
	ServiceName    string
	defaultService micro.Service
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
	opts = append(opts, micro.Registry(registry))
	opts = append(opts, micro.WrapHandler(validator.NewHandlerWrapper()))
	opts = append(opts, micro.WrapClient(opentracing.NewClientWrapper(nil)))
	opts = append(opts, micro.WrapHandler(opentracing.NewHandlerWrapper(nil)))
	opts = append(opts, micro.WrapSubscriber(opentracing.NewSubscriberWrapper(nil)))
	opts = append(opts, micro.RegisterTTL(time.Second*90))
	opts = append(opts, micro.RegisterInterval(time.Second*30))
	opts = append(opts, micro.Version("latest"))
	opts = append(opts, micro.Context(ctx))
	opts = append(opts, micro.WrapHandler(loggerWrapper(logger.DefaultLogger)))
	opts = append(opts, micro.WrapHandler(loggerHandler))
	service := &service{micro.NewService(
		opts...,
	)}
	service.Init()
	service.Server().Init(
		server.Wait(nil),
	)
	defaultService = service
	ServiceName = service.Name()

	// init tracer
	initJaegerTracer(service.Name(), opentracingAddress)
	// init conf
	setupService(service)
	// init db
	var dbOpts []db.Options
	json.Unmarshal(config.ConfBytes("db"), &dbOpts)
	key := []byte("Uu9sbdsduYUSudhs18dh/w==")
	for i, opt := range dbOpts {
		tmp, _ := base64.StdEncoding.DecodeString(opt.Passwd)
		var temp []byte
		temp, _ = openssl.Des3ECBDecrypt(tmp, key, openssl.PKCS5_PADDING)
		dbOpts[i].Passwd = string(temp)
	}
	db.SetConns(dbdriver, dbOpts)

	return service
}
