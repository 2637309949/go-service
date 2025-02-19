package main

import (
	"auth/handler"
	"comm/service"
	pbAuth "proto/auth"
)

func main() {
	service := service.NewService(
		service.Name("auth"),
	)
	pbAuth.RegisterAuthServiceHandler(service.Server(), handler.NewHandler())
	service.Run()
}
