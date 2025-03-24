package main

import (
	"auth/handler"
	"comm/service"
	pbApigate "proto/apigate"
)

func main() {
	service := service.NewService(
		service.Name("auth"),
	)
	pbApigate.RegisterAuthServiceHandler(service.Server(), handler.NewHandler())
	service.Run()
}
