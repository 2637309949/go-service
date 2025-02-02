package logger

import (
	"context"
	"testing"

	"go-micro.dev/v5/logger"
)

func TestLogger(t *testing.T) {
	l := logger.NewLogger(logger.WithLevel(logger.TraceLevel), logger.WithCallerSkipCount(2))

	h1 := logger.NewHelper(l).WithFields(map[string]interface{}{"key1": "val1"})
	h1.Log(logger.TraceLevel, "simple log before trace_msg1")
	h1.Trace("trace_msg1")
	h1.Log(logger.TraceLevel, "simple log after trace_msg1")
	h1.Warn("warn_msg1")

	h2 := logger.NewHelper(l).WithFields(map[string]interface{}{"key2": "val2"})
	h2.Logf(logger.TraceLevel, "formatted log before trace_msg%s", "2")
	h2.Trace("trace_msg2")
	h2.Logf(logger.TraceLevel, "formatted log after trace_msg%s", "2")
	h2.Warn("warn_msg2")

	l = logger.NewLogger(logger.WithLevel(logger.TraceLevel), logger.WithCallerSkipCount(1))
	l.Fields(map[string]interface{}{"key3": "val4"}).Log(logger.InfoLevel, "test_msg")
}

func TestExtract(t *testing.T) {
	l := logger.NewLogger(logger.WithLevel(logger.TraceLevel), logger.WithCallerSkipCount(2)).Fields(map[string]interface{}{"requestID": "req-1"})

	ctx := logger.NewContext(context.Background(), l)

	logger.Info("info message without request ID")
	logger.Extract(ctx).Info("info message with request ID")
}
