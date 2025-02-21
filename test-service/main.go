package main

import (
	"comm/service"
	pbCache "proto/cache"
	pbProxy "proto/proxy"
	pbTest "proto/test"
	"test/handler"
)

func main() {
	service := service.NewService(
		service.Name("test"),
	)
	h := handler.Handler{
		CacheService: pbCache.NewCacheService("cache", service.Client()),
		MeiZuService: pbProxy.NewMeiZuService("proxy", service.Client()),
	}
	pbTest.RegisterTestServiceHandler(service.Server(), &h)
	service.Run()
}
