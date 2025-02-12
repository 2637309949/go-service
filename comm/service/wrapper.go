package service

import (
	"context"
	"strings"

	"comm/logger"

	"go-micro.dev/v5/metadata"
	"go-micro.dev/v5/server"
)

func loggerHandler(h server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		logger := logger.Extract(ctx)
		logger.WithFields(map[string]interface{}{"invoke": req.Endpoint()}).Info(">>>>>>>>>>>>>>>>>>>>>")
		return h(ctx, req, rsp)
	}
}

func loggerWrapper(l logger.Logger) server.HandlerWrapper {
	if l == nil {
		l = logger.DefaultLogger
	}
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			_, ok := logger.FromContext(ctx)
			if !ok {
				if md, ok := metadata.FromContext(ctx); ok {
					if traceID, exists := md["Uber-Trace-Id"]; exists {
						traceID = strings.Split(traceID, ":")[0]
						l = l.Fields(map[string]interface{}{"traceid": traceID})
						ctx = logger.NewContext(ctx, l)
					}
				}
			}
			if err := h(ctx, req, rsp); err != nil {
				return err
			}
			return nil
		}
	}
}
