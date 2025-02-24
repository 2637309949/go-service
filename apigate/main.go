package main

import (
	"apigate/handler"
	"apigate/handler/rpc"
	"apigate/handler/web"
	"apigate/router"
	"apigate/router/registry"
	"apigate/util"
	"comm/config"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/plugins/v5/registry/consul"
	"github.com/micro/plugins/v5/wrapper/trace/opentracing"
	"go-micro.dev/v5/logger"
	regi "go-micro.dev/v5/registry"
)

var (
	serviceName = "apigate"
	apiBase     = "/api"
	tracer      = util.InitJaegerTracer(serviceName)
)

func main() {
	addr := config.String("addr")
	consulAddress := config.String("consul")
	consulRegistry := consul.NewRegistry(func(op *regi.Options) {
		op.Addrs = []string{
			consulAddress,
		}
	})
	opts := []handler.Option{}
	opts = append(opts, handler.WithApiBase(apiBase))
	opts = append(opts, handler.WithRouter(
		registry.NewRouter(
			router.WithAuth(router.NewJWT()),
			router.WithRegistry(consulRegistry),
		),
	))
	opts = append(opts, handler.WithWrapCall(opentracing.NewCallWrapper(tracer)))
	opts = append(opts, handler.WithHandler(rpc.NewHandler(opts...)))
	opts = append(opts, handler.WithHandler(web.NewHandler(opts...)))
	hd := handler.NewHandler(opts...)
	gin.DefaultWriter = io.Discard
	logger.Infof("Starting server %s", addr)
	r := gin.Default()
	r.Use(util.LoggerWithConfig(gin.LoggerConfig{}))
	r.Use(util.CorsMiddle)
	r.Any("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": "5+",
		})
	})
	r.GET("/favicon.ico", func(c *gin.Context) {})
	r.Use(util.ResolverMiddle(apiBase))
	r.Use(util.TracerMiddle(tracer))
	r.NoRoute(func(c *gin.Context) {
		hd.ServeHTTP(c.Writer, c.Request)
	})
	r.Run(addr)
	logger.Info("Stopping server")
}
