package main

import (
	"apigate/handler"
	"apigate/handler/rpc"
	"apigate/handler/web"
	"apigate/router"
	"apigate/router/registry"
	"comm/config"
	wb "comm/service/web"
	"context"
	"net/http"
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
		wb.WrapHandler(func(h wb.HandlerFunc) wb.HandlerFunc {
			set := func(w http.ResponseWriter, k, v string) {
				if v := w.Header().Get(k); len(v) > 0 {
					return
				}
				w.Header().Set(k, v)
			}
			return func(w http.ResponseWriter, r *http.Request) {
				if origin := r.Header.Get("Origin"); len(origin) > 0 {
					set(w, "Access-Control-Allow-Origin", origin)
				} else {
					set(w, "Access-Control-Allow-Origin", "*")
				}
				set(w, "Access-Control-Allow-Credentials", "true")
				set(w, "Access-Control-Allow-Methods", "POST, PATCH, GET, OPTIONS, PUT, DELETE")
				set(w, "Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Micro-Namespace")

				if r.Method == "OPTIONS" {
					return
				}
				h(w, r)
			}
		}),
		wb.WrapHandler(func(h wb.HandlerFunc) wb.HandlerFunc {
			resolver := registry.NewResolver()
			return func(w http.ResponseWriter, r *http.Request) {
				endpoint := resolver.Resolve(apiBase, r)
				ctx := r.Context()
				ctx = context.WithValue(ctx, registry.Endpoint{}, *endpoint)
				*r = *r.Clone(ctx)
				h(w, r)
			}
		}),
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
	service.Handle(fmt.Sprintf("%s/", apiBase), hd, wb.Seen(true))

	// register handler
	pb.RegisterApigateServiceHandler(service, hd)
	service.Run()
}
