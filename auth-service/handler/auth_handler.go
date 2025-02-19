package handler

import (
	"comm/mark"
	"context"
	pbAuth "proto/auth"

	"go-micro.dev/v5/logger"
)

func (h *Handler) Generate(ctx context.Context, req *pbAuth.Account, rsp *pbAuth.Token) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Generate")()
	logger.Infof("Received Handler.Get request: %+v", req)

	return nil
}

func (h *Handler) Inspect(ctx context.Context, req *pbAuth.Token, rsp *pbAuth.Account) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Inspect")()
	logger.Infof("Received Handler.Get request: %+v", req)

	return nil
}

func (h *Handler) Verify(ctx context.Context, req *pbAuth.Credential, rsp *pbAuth.Empty) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Verify")()
	logger.Infof("Received Handler.Get request: %+v", req)

	return nil
}

func (h *Handler) Refresh(ctx context.Context, req *pbAuth.Token, rsp *pbAuth.Token) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Refresh")()
	logger.Infof("Received Handler.Get request: %+v", req)

	return nil
}
