package handler

import (
	"comm/logger"
	"comm/mark"
	"context"
	pbDtm "proto/dtm"
)

func (h *Handler) TransOut(ctx context.Context, req *pbDtm.BusiReq, rsp *pbDtm.BusiReply) error {
	var err error
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "TransOut")()
	logger.Infof("%+v", err)
	return nil
}

func (h *Handler) TransOutRevert(ctx context.Context, req *pbDtm.BusiReq, rsp *pbDtm.BusiReply) error {
	var err error
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "TransOutRevert")()
	logger.Infof("%+v", err)
	return nil
}

func (h *Handler) TransIn(ctx context.Context, req *pbDtm.BusiReq, rsp *pbDtm.BusiReply) error {
	var err error
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "TransIn")()
	logger.Infof("%+v", err)
	return nil
}

func (h *Handler) TransInRevert(ctx context.Context, req *pbDtm.BusiReq, rsp *pbDtm.BusiReply) error {
	var err error
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "TransInRevert")()
	logger.Infof("%+v", err)
	return nil
}

func (h *Handler) TransResume(ctx context.Context, req *pbDtm.BusiReq, rsp *pbDtm.BusiReply) error {
	var err error
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "TransResume")()
	logger.Infof("%+v", err)
	return nil
}
