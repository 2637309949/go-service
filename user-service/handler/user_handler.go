package handler

import (
	"comm/db"
	"comm/logger"
	"comm/mark"
	"context"
	pbUser "proto/user"
	"user/models/model"
)

func (h *Handler) QueryUserDetail(ctx context.Context, req *pbUser.UserFilter, rsp *pbUser.User) error {
	var err error
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "QueryUserDetail")()

	session, _ := db.InitDb(ctx, uint64(req.GetUserId()))
	timemark.Mark("InitDb")

	var data model.User
	where := model.User{
		UserId: req.GetUserId(),
	}

	err = h.QueryUserDetailDB(ctx, session, &where, &data)
	timemark.Mark("QueryUserDetailDB")
	if err != nil {
		logger.Debugf("QueryUserDetailDB failed. [%s]", err.Error())
		return err
	}

	data.Marshal(rsp)
	return nil
}

func (h *Handler) QueryUser(ctx context.Context, req *pbUser.UserFilter, rsp *pbUser.UserList) error {
	var err error
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "QueryUserDetail")()

	session, _ := db.InitDb(ctx, uint64(req.GetUserId()))
	session = db.SetLimit(ctx, session, req)
	session = db.SetOrder(ctx, session, req)
	timemark.Mark("InitDb")

	var totalCount int32
	var lst []model.User
	where := model.User{
		UserId: req.GetUserId(),
	}

	h.QueryUserDB(ctx, session, &where, &lst, &totalCount)
	timemark.Mark("QueryUserDB")
	if err != nil {
		logger.Errorf("QueryUserDB failed. [%s]", err.Error())
		return err
	}

	model.UserMarshalLst(lst, &rsp.Data)

	rsp.TotalCount = totalCount
	rsp.CurPage = req.GetPageNo()
	rsp.TotalPage = totalCount / req.GetPageSize()
	if totalCount%req.GetPageSize() != 0 {
		rsp.TotalPage += 1
	}

	return nil
}

func (h *Handler) InsertUser(ctx context.Context, req *pbUser.User, rsp *pbUser.User) error {
	var err error
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "QueryUserDetail")()

	session, _ := db.InitDb(ctx, uint64(req.GetUserId()))
	timemark.Mark("InitDb")

	var inUser model.User
	inUser.Unmarshal(req)
	where := model.User{
		UserId: req.GetUserId(),
	}

	err = h.QueryUserDetailDB(ctx, session, &where, &inUser)
	timemark.Mark("QueryUserDetailDB")
	if err == nil {
		logger.Info("记录已存在.")
		goto resp
	}

	err = h.InsertUserDB(ctx, session, &inUser)
	timemark.Mark("InsertUserDB")
	if err != nil {
		logger.Errorf("InsertUserDB failed. [%v]", err)
		return err
	}

resp:
	inUser.Marshal(rsp)

	return nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *pbUser.User, rsp *pbUser.User) error {
	var err error
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "QueryUserDetail")()

	session, _ := db.InitDb(ctx, uint64(req.GetUserId()))
	timemark.Mark("InitDb")

	var inUser model.User
	inUser.Unmarshal(req)

	err = h.UpdateUserDB(ctx, session, &inUser)
	timemark.Mark("UpdateUserDB")
	if err != nil {
		logger.Errorf("UpdateUserDB failed. [%v]", err)
		return err
	}

	inUser.Marshal(rsp)

	return nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *pbUser.User, rsp *pbUser.User) error {
	var err error
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "QueryUserDetail")()

	session, _ := db.InitDb(ctx, uint64(req.GetUserId()))
	timemark.Mark("InitDb")

	where := model.User{
		UserId: req.GetUserId(),
	}

	err = h.DeleteUserDB(ctx, session, &where)
	timemark.Mark("DeleteUserDB")
	if err != nil {
		logger.Errorf("DeleteUserDB failed. [%v]", err)
		return err
	}

	return nil
}

func (h *Handler) Test(ctx context.Context, req *pbUser.User, rsp *pbUser.User) error {
	logger := logger.Extract(ctx)
	logger.Error("Test ok.")
	return nil
}
