package main

import (
	"apigate/handler"
	"apigate/handler/rpc"
	"apigate/handler/web"
	"apigate/router"
	"apigate/router/registry"
	"comm/config"
	wb "comm/service/web"
	"fmt"
	pb "proto/apigate"
)

func main() {
	apiBase := "/api"
	service := wb.NewService(
		wb.Name("apigate"),
		wb.Address(config.String("addr")),
	)

	// register router
	opts := []handler.Option{}
	// authService := pb.NewAuthService("auth", service.Options().Service.Client())
	// r := registry.NewRouter(router.WithAuth(authService))
	r := registry.NewRouter(
		router.WithAuth(router.NewJWT()),
		router.WithRegistry(service.Options().Registry),
	)
	opts = append(opts, handler.WithApiBase(apiBase))
	opts = append(opts, handler.WithRouter(r))
	opts = append(opts, handler.WithHandler(rpc.NewHandler(opts...)))
	opts = append(opts, handler.WithHandler(web.NewHandler(opts...)))
	hd := handler.NewHandler(opts...)

	// r.Init(router.WithRegistry(service.Options().Registry))
	service.NoSeen(fmt.Sprintf("%s/", apiBase), hd)

	// register handler
	pb.RegisterApigateServiceHandler(service, hd)
	service.Run()
}
