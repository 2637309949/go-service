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
	h := handler.Handler{}
	pbUser.RegisterUserServiceHandler(service.Server(), &h)
	service.Run()
}
