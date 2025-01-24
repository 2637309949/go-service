package main

import (
	"context"
	pb "hello/proto"
	"log"

	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(registry),
	)
	service.Init()
	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatalf("Failed to run service: %v", err)
	}
}
