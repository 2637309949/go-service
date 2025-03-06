package timed

import (
	"comm/config"
	"comm/logger"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	trace "github.com/micro/plugins/v5/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v5/metadata"
)

type timed struct {
	LimitEnv    map[string]bool
	LimitEnvStr string
	FuncName    string
	Dur         time.Duration
	Timeout     time.Duration
	LockKey     string
	RedisClient *redis.Client
}

func NewTimed() *timed {
	return &timed{
		LimitEnv: make(map[string]bool),
	}
}

type TimeInitOption func(t *timed)

func SetFuncName(funcName string) TimeInitOption {
	return func(t *timed) {
		t.FuncName = funcName
	}
}

func SetRedisClient(redisClient *redis.Client) TimeInitOption {
	return func(t *timed) {
		t.RedisClient = redisClient
	}
}

func SetLimitEnv(env ...string) TimeInitOption {
	return func(t *timed) {
		for _, v := range env {
			t.LimitEnv[v] = true
			t.LimitEnvStr += v + ","
		}

		t.LimitEnvStr = strings.Trim(t.LimitEnvStr, ",")
	}
}
func SetDuration(dur time.Duration) TimeInitOption {
	return func(t *timed) {
		t.Dur = dur
	}
}

func SetTimeout(timeout time.Duration) TimeInitOption {
	return func(t *timed) {
		t.Timeout = timeout
	}
}

func SetLockKey(lockKeY string) TimeInitOption {
	return func(t *timed) {
		t.LockKey = lockKeY
	}
}

func (t *timed) TimedInit(options ...TimeInitOption) (context.Context, *redislock.Lock, context.CancelFunc, error) {

	for _, opt := range options {
		opt(t)
	}
	if t.FuncName == "" {
		return nil, nil, nil, errors.New("funcName is empty")
	}
	if t.LockKey == "" {
		return nil, nil, nil, errors.New("lockKey is empty")
	}
	if t.Dur == 0 {
		return nil, nil, nil, errors.New("dur is empty")
	}
	if t.Timeout == 0 {
		t.Timeout = t.Dur
	}
	if t.RedisClient == nil {
		return nil, nil, nil, errors.New("redisClient is empty")
	}

	funcname := t.FuncName
	dur := t.Dur
	timeout := t.Timeout
	lockKey := t.LockKey
	redisClient := t.RedisClient
	deadline := time.Now().Add(timeout)

	ot := opentracing.GlobalTracer()

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	ctx, span, err := trace.StartSpanFromContext(ctx, ot, funcname)
	if err != nil {
		return nil, nil, cancel, err
	}
	defer span.Finish()

	if md, ok := metadata.FromContext(ctx); ok {
		if traceID, exists := md["Uber-Trace-Id"]; exists {
			traceID = strings.Split(traceID, ":")[0]
			l := logger.DefaultLogger.Fields(map[string]interface{}{"traceid": traceID, "timed": funcname})
			ctx = logger.NewContext(ctx, l)
		}
	}

	env := config.CommConf("env")
	if _, ok := t.LimitEnv[env]; !ok && len(t.LimitEnv) > 0 {
		err = fmt.Errorf("timed env err. cur [%s] limit [%s]", env, t.LimitEnvStr)
		return ctx, nil, cancel, err
	}

	locker := redislock.New(redisClient)
	lock, err := locker.Obtain(ctx, lockKey, dur, nil)
	if err == redislock.ErrNotObtained {
		return ctx, nil, cancel, err
	} else if err != nil {
		return ctx, nil, cancel, err
	}

	return ctx, lock, cancel, err
}

func TimedClose(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
