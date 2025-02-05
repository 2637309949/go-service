package handler

import (
	"comm/logger"
	"comm/mark"
	"context"
	pbUser "proto/user"
)

func (h *Handler) Hello(ctx context.Context, req *pbUser.Request, rsp *pbUser.Response) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Hello")()

	logger.Infof("Greeting Hello: %v", req.Name)
	rsp.Greeting = "Hello " + req.Name
	timemark.Mark("Greeting")

	return nil
}
