package main

import (
	"comm/service"
	pbUser "proto/user"
	"user/handler"
)

func main() {
	service := service.NewService(
		service.Name("user"),
	)
	pbUser.RegisterEndpointHandler(service.Server(), new(handler.Handler))
	service.Run()
}
