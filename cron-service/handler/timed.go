package handler

import (
	"comm/consts"
	"comm/logger"
	"comm/timed"
	pb "proto/cron"
	"time"
)

var (
	req  = &pb.TimedReq{}
	resp = &pb.TimedResp{}
)

func (h *Handler) TimedLog() {
	timedV1 := timed.NewTimed()
	ctx, lock, cancel, err := timedV1.TimedInit(
		timed.SetLimitEnv("test"),
		timed.SetFuncName("TimedLog"),
		timed.SetDuration(10*time.Second),
		timed.SetTimeout(2*time.Hour),
		timed.SetLockKey(consts.TimedLog),
		timed.SetRedisClient(h.RedisClient))
	logger := logger.Extract(ctx)
	if err != nil {
		logger.Debugf("%s fail. [%s]", consts.TimedLog, err.Error())
		return
	}
	defer cancel()
	defer lock.Release(ctx)
	h.Log(ctx, req, resp)
}
