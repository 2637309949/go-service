package main

import (
	"comm/service"
	pbCache "proto/cache"
	pbUser "proto/user"
	"user/handler"
)

func main() {
	service := service.NewService(
		service.Name("user"),
	)
	h := handler.Handler{
		CacheService: pbCache.NewCacheService("cache", service.Client()),
	}
	pbUser.RegisterUserServiceHandler(service.Server(), &h)
	service.Run()
}
