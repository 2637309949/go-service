package main

import (
	"apigate/handler"
	"apigate/router"
	"apigate/router/registry"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/micro/plugins/v5/registry/consul"
	"github.com/micro/plugins/v5/wrapper/trace/opentracing"
	"go-micro.dev/v5/client"
	"go-micro.dev/v5/config"
	"go-micro.dev/v5/logger"
	regi "go-micro.dev/v5/registry"
)

var (
	serviceName = "apigate"
	apiBase     = "/api"
)

func main() {
	logger.Info("Starting server")
	initJaegerTracer(serviceName)
	cli := client.DefaultClient
	cli.Init(client.WrapCall(opentracing.NewCallWrapper(nil)))
	cli.Init(client.WrapCall(modifyRsp))

	consulAddress := config.Get("consul").String("")
	consulRegistry := consul.NewRegistry(func(op *regi.Options) {
		op.Addrs = []string{
			consulAddress,
		}
	})

	opts := []handler.Option{
		handler.WithClient(cli),
		handler.WithRouter(registry.NewRouter(
			router.WithApiBase(apiBase),
			router.WithRegistry(consulRegistry),
		)),
	}
	hd := handler.NewHandler(opts...)

	gin.DefaultWriter = io.Discard
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "PATCH", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:    []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Micro-Namespace"},
	}))
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": "5.0",
		})
	})
	r.GET("/favicon.ico", func(c *gin.Context) {})
	r.Use(auth(apiBase))
	r.NoRoute(func(c *gin.Context) {
		hd.ServeHTTP(c.Writer, c.Request)
	})
	r.Run(":8080")
	logger.Info("Stopping server")
}
