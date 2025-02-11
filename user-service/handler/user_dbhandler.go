package handler

import (
	"comm/logger"
	"context"
	"errors"
	"user/models/model"

	"github.com/jinzhu/gorm"
)

func (s *Handler) QueryUserDB(ctx context.Context, session *gorm.DB, where *model.User, list *[]model.User, count ...*int32) error {
	logger := logger.Extract(ctx)
	session = session.Table(where.TableName()).Where(where).Find(list)
	if len(count) > 0 {
		session = session.Offset(0).Count(count[0])
	}

	if errs := session.GetErrors(); len(errs) != 0 {
		logger.Errorf("QueryUserDB failed. [%v]", errs)
		return errors.New("查询User失败")
	}
	return nil
}

func (s *Handler) QueryUserDetailDB(ctx context.Context, session *gorm.DB, where *model.User, data *model.User) error {
	var err error
	logger := logger.Extract(ctx)
	var lst []model.User
	err = s.QueryUserDB(ctx, session.Limit(1), where, &lst)
	if err != nil {
		logger.Errorf("QueryUserDB failed. [%s]", err.Error())
		return err
	}

	if len(lst) == 0 {
		logger.Warn("QueryUserDB empty.")
		return errors.New("查询User为空")
	}

	*data = lst[0]
	return nil
}

func (s *Handler) InsertUserDB(ctx context.Context, session *gorm.DB, data *model.User) error {
	var err error
	logger := logger.Extract(ctx)
	err = session.Create(data).Error
	if err != nil {
		logger.Errorf("InsertUserDB failed. [%s]", err.Error())
		return errors.New("新增User失败")
	}
	return nil
}

func (s *Handler) UpdateUserDB(ctx context.Context, session *gorm.DB, data *model.User) error {
	var err error
	logger := logger.Extract(ctx)
	err = session.Table(data.TableName()).Model(&data).Updates(&data).Error
	if err != nil {
		logger.Errorf("UpdateUserDB failed. [%s]", err.Error())
		return errors.New("更新User失败")
	}
	return nil
}

func (s *Handler) SaveUserDB(ctx context.Context, session *gorm.DB, data *model.User) error {
	var err error
	logger := logger.Extract(ctx)
	err = session.Save(data).Error
	if err != nil {
		logger.Errorf("SaveUserDB failed. [%s]", err.Error())
		return errors.New("保存User失败")
	}
	return nil
}

func (s *Handler) DeleteUserDB(ctx context.Context, session *gorm.DB, data *model.User) error {
	var err error
	logger := logger.Extract(ctx)
	err = session.Where(data).Delete(&data).Error
	if err != nil {
		logger.Errorf("DeleteUserDB failed. [%s]", err.Error())
		return errors.New("删除User失败")
	}
	return nil
}
