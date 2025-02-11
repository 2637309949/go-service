package main

import (
	"cache/handler"
	"comm/service"
	pbCache "proto/cache"
)

func main() {
	service := service.NewService(
		service.Name("cache"),
	)
	pbCache.RegisterCacheServiceHandler(service.Server(), handler.NewHandler())
	service.Run()
}
