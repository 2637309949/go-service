package service

import (
	"context"

	"comm/logger"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go-micro.dev/v5/server"
)

func LoggerWrapper(l logger.Logger) server.HandlerWrapper {
	if l == nil {
		l = logger.DefaultLogger
	}
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			_, ok := logger.FromContext(ctx)
			if !ok {
				span := opentracing.SpanFromContext(ctx)
				traceID := ""
				if span != nil {
					traceID = span.Context().(jaeger.SpanContext).TraceID().String()
				}
				l = l.Fields(map[string]interface{}{"traceid": traceID})
				ctx = logger.NewContext(ctx, l)
			}
			if err := h(ctx, req, rsp); err != nil {
				return err
			}
			return nil
		}
	}
}
