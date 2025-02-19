package main

import (
	"apigate/handler"
	"apigate/router"
	"apigate/router/registry"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/plugins/v5/registry/consul"
	"github.com/micro/plugins/v5/wrapper/trace/opentracing"
	"go-micro.dev/v5/config"
	"go-micro.dev/v5/config/source/env"
	"go-micro.dev/v5/logger"
	regi "go-micro.dev/v5/registry"
)

var (
	serviceName = "apigate"
	apiBase     = "/api"
)

func main() {
	logger.Info("Starting server")
	err := config.Load(env.NewSource())
	if err != nil {
		logger.Fatalf("Error source load: %v", err)
	}
	tracer := initJaegerTracer(serviceName)
	consulAddress := config.Get("consul").String("")
	consulRegistry := consul.NewRegistry(func(op *regi.Options) {
		op.Addrs = []string{
			consulAddress,
		}
	})

	opts := []handler.Option{}
	opts = append(opts, handler.WithApiBase(apiBase))
	opts = append(opts, handler.WithRouter(
		registry.NewRouter(
			router.WithAuth(new(jwt)),
			router.WithRegistry(consulRegistry),
		),
	))
	opts = append(opts, handler.WithWrapCall(opentracing.NewCallWrapper(tracer)))
	hd := NewHandler(opts...)

	gin.DefaultWriter = io.Discard
	r := gin.Default()
	r.Use(corsMiddle)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": "5.1",
		})
	})
	r.GET("/favicon.ico", func(c *gin.Context) {})
	r.Use(resolverMiddle(apiBase))
	r.Use(tracerMiddle(tracer))
	r.NoRoute(func(c *gin.Context) {
		hd.ServeHTTP(c.Writer, c.Request)
	})
	r.Run(":8080")
	logger.Info("Stopping server")
}
