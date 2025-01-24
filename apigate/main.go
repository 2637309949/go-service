package main

import (
	"apigate/api/handler"
	"apigate/api/router"
	"apigate/api/router/registry"
	"io"

	"github.com/micro/plugins/v5/registry/consul"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v5/client"
	"go-micro.dev/v5/logger"
)

func main() {
	logger.Info("Starting server")

	consulRegistry := consul.NewRegistry()
	opts := []handler.Option{
		handler.WithClient(client.DefaultClient),
		handler.WithRouter(registry.NewRouter(router.WithRegistry(consulRegistry))),
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
		c.JSON(200, gin.H{
			"version": "1.0",
		})
	})
	r.GET("/favicon.ico", func(c *gin.Context) {
	})
	r.Use(auth())
	r.NoRoute(func(c *gin.Context) {
		hd.ServeHTTP(c.Writer, c.Request)
	})
	r.Run(":8080")
	logger.Info("Stopping server")
}
