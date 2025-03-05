// Package main
package main

import (
	"context"

	hello "github.com/go-micro/examples/greeter/srv/proto/hello"
	"github.com/go-micro/plugins/v5/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/util/log"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Log("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {

	service := micro.NewService(
		micro.Server(grpc.NewServer()),
		micro.Name("greeter"),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	hello.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
