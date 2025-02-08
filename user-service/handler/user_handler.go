package handler

import (
	"comm/logger"
	"comm/mark"
	"context"
	pbUser "proto/user"
)

func (h *Handler) QueryUserDetail(ctx context.Context, req *pbUser.UserFilter, rsp *pbUser.User) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "QueryUserDetail")()

	logger.Infof("Greeting Hello: %v", req.Name)
	rsp.Greeting = "Hello " + req.Name
	timemark.Mark("Greeting")

	return nil
}

func (h *Handler) QueryUser(ctx context.Context, req *pbUser.User, rsp *pbUser.UserList) error {
	return nil
}

func (h *Handler) InsertUser(ctx context.Context, req *pbUser.User, rsp *pbUser.User) error {
	return nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *pbUser.User, rsp *pbUser.User) error {
	return nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *pbUser.User, rsp *pbUser.User) error {
	return nil
}
