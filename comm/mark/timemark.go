package mark

import (
	"comm/logger"
	"context"
	"fmt"
	"time"
)

type TimeMark struct {
	desc        string
	start, last time.Time
}

func (t *TimeMark) Mark(format string, data ...interface{}) {
	duration := time.Since(t.last)
	t.last = time.Now()
	desc := fmt.Sprintf(format, data...)
	t.desc += fmt.Sprintf(" %s[%s]", desc, duration)
}

func (t *TimeMark) Init(ctx context.Context, format string, data ...interface{}) func() {
	logger := logger.Extract(ctx)
	t.start = time.Now()
	t.last = time.Now()
	return func() {
		desc := fmt.Sprintf(format, data...)
		duration := time.Since(t.start)
		if len(t.desc) > 0 {
			logger.Infof("%s total duration:[%v] step%s", desc, duration, t.desc)
		} else {
			logger.Infof("%s total duration:[%v]", desc, duration)
		}
	}
}
