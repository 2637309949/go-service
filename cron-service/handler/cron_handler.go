package handler

import (
	"comm/logger"
	"comm/mark"
	"context"
	pb "proto/cron"
)

func (h *Handler) Log(ctx context.Context, req *pb.TimedReq, resp *pb.TimedResp) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Log")()
	logger.Infof("Received Handler.Log request: %+v", req)
	return nil
}
