package handler

import (
	"comm/db"
	"comm/logger"
	"comm/mark"
	"context"
	pbUser "proto/user"
)

func (h *Handler) QueryUserDetail(ctx context.Context, req *pbUser.UserFilter, rsp *pbUser.User) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "QueryUserDetail")()

	session, _ := db.InitDb(ctx)
	timemark.Mark("InitDb")
	type Result struct {
		Sum int `gorm:"column:sum"`
	}
	var result Result
	session.Raw("SELECT 1 + 1 as sum").Scan(&result)
	logger.Info(result.Sum) // 输出 2

	logger.Infof("Greeting Hello: %v", req.UserId)
	rsp.UserId = req.UserId
	timemark.Mark("Greeting")

	return nil
}

func (h *Handler) QueryUser(ctx context.Context, req *pbUser.UserFilter, rsp *pbUser.UserList) error {
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
