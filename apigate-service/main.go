package main

import (
	"apigate/handler"
	"apigate/handler/rpc"
	"apigate/handler/web"
	"apigate/router"
	"apigate/router/registry"
	"apigate/util"
	"comm/config"
	wb "comm/service/web"
	"fmt"
	pb "proto/apigate"
)

var (
	apiBase = "/api"
)

func main() {
	service := wb.NewService(
		wb.Name("apigate"),
		wb.Address(config.String("addr")),
		wb.WrapHandler(util.WrapCors()),
	)

	// register router
	opts := []handler.Option{}
	r := registry.NewRouter(router.WithAuth(router.NewJWT()))
	opts = append(opts, handler.WithApiBase(apiBase))
	opts = append(opts, handler.WithRouter(r))
	opts = append(opts, handler.WithHandler(rpc.NewHandler(opts...)))
	opts = append(opts, handler.WithHandler(web.NewHandler(opts...)))
	hd := handler.NewHandler(opts...)
	r.Init(router.WithRegistry(service.Options().Registry))
	service.Handle(fmt.Sprintf("%s/", hd.Options().ApiBase), hd, wb.NoSeen(true))

	// register handler
	pb.RegisterApigateServiceHandler(service, hd)
	service.Run()
}
