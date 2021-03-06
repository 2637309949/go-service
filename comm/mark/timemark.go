package mark

import (
	"comm/logger"
	"context"
	"fmt"
	"time"
)

type TimeMark struct {
	ctx         context.Context
	start, last time.Time
}

func (t *TimeMark) Mark(format string, data ...interface{}) {
	duration := time.Since(t.last)
	desc := fmt.Sprintf(format, data...)
	logger.Init(logger.WithCallerSkipCount(3))
	logger.Infof(t.ctx, "%s duration:[%s]", desc, duration)
	logger.Init(logger.WithCallerSkipCount(2))
	t.last = time.Now()
}

func (t *TimeMark) Init(ctx context.Context, format string, data ...interface{}) func() {
	t.ctx = ctx
	t.start = time.Now()
	t.last = time.Now()
	return func() {
		desc := fmt.Sprintf(format, data...)
		duration := time.Since(t.start)
		logger.Init(logger.WithCallerSkipCount(3))
		logger.Infof(t.ctx, "%s total duration:[%v]", desc, duration)
		logger.Init(logger.WithCallerSkipCount(2))
	}
}
