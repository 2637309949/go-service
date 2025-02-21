package util

import (
	"apigate/router/registry"
	"context"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/micro/plugins/v5/wrapper/trace/opentracing"
	ogo "github.com/opentracing/opentracing-go"
)

var CorsMiddle = cors.New(cors.Config{
	AllowAllOrigins: true,
	AllowMethods:    []string{"POST", "PATCH", "GET", "OPTIONS", "PUT", "DELETE"},
	AllowHeaders:    []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Micro-Namespace"},
})

func ResolverMiddle(apiBase string) gin.HandlerFunc {
	resolver := registry.NewResolver()
	return func(cx *gin.Context) {
		req := cx.Request
		endpoint := resolver.Resolve(apiBase, req)

		ctx := req.Context()
		ctx = context.WithValue(ctx, registry.Endpoint{}, *endpoint)

		*req = *req.Clone(ctx)
		cx.Next()
	}
}

func TracerMiddle(tracer ogo.Tracer) gin.HandlerFunc {
	return func(cx *gin.Context) {
		name := "apigate"
		v := cx.Request.Context().Value(registry.Endpoint{})
		if v != nil {
			endpoint := v.(registry.Endpoint)
			name = fmt.Sprintf("%s.%s", endpoint.Name, endpoint.Path)
		}
		ctx, span, err := opentracing.StartSpanFromContext(cx.Request.Context(), tracer, name)
		if err != nil {
			return
		}
		defer span.Finish()
		*cx.Request = *cx.Request.Clone(ctx)
		cx.Next()
	}
}
