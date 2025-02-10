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
	pbCache.RegisterCacheServiceHandler(service.Server(), new(handler.Handler))
	service.Run()
}
