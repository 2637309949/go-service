package main

import (
	"comm/define"
	"comm/logger"
	"comm/service"
	"comm/store"
	"helloworld-service/handler"
	"proto/cache"
	"proto/helloworld"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	srv := service.New(service.Name("helloworld"))
	hdl := handler.Handler{
		CacheStore: store.CacheStore(cache.NewCacheService("cache", srv.Client())),
	}
	helloworld.RegisterHelloworldHandler(srv.Server(), &hdl)
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
