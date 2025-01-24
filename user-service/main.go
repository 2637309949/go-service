package main

import (
	pbUser "proto/user"
	 "user/handler"
	"log"

	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5"
)

func main() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("user"),
		micro.Registry(registry),
	)
	service.Init()
	pbUser.RegisterHandlerHandler(service.Server(), new(handler.Handler))

	if err := service.Run(); err != nil {
		log.Fatalf("Failed to run service: %v", err)
	}
}
