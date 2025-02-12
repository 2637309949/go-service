package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"go-micro.dev/v5/logger"
)

func initJaegerTracer(serviceName string) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	tracer, _, err := cfg.NewTracer()
	if err != nil {
		logger.Fatalf("Error tracer: %v", err)
	}
	opentracing.SetGlobalTracer(tracer)
}
