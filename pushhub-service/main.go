package main

import (
	"comm/service"
	pb "proto/pushhub"
	"test/handler"
)

func main() {
	service := service.NewService(
		service.Name("pushhub"),
	)
	h := handler.Handler{}
	pb.RegisterPushhubServiceHandler(service.Server(), &h)
	service.Run()
}
