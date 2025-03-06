package util

import (
	"apigate/router/registry"
	"context"
	"fmt"
	"time"

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

var defaultLogFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}
	return fmt.Sprintf("%v |%s %3d %s| %13v | %15s |%s %-7s %s %s\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

// LoggerWithConfig instance a Logger middleware with config.
func LoggerWithConfig(conf gin.LoggerConfig) gin.HandlerFunc {
	formatter := conf.Formatter
	if formatter == nil {
		formatter = defaultLogFormatter
	}

	notlogged := conf.SkipPaths

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when it is not being skipped
		if _, ok := skip[path]; ok || (conf.Skip != nil && conf.Skip(c)) {
			return
		}

		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		fmt.Print(formatter(param))
	}
}
