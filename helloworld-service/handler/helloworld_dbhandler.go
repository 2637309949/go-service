package handler

import (
	"comm/errors"
	"comm/logger"
	"comm/service"
	"context"
	"helloworld-service/model"

	"gorm.io/gorm"
)

// QueryInfoDB defined TODO
func (h *Handler) QueryInfoDB(ctx context.Context, session *gorm.DB, where *model.Info, list *[]*model.Info, count ...*int64) error {
	session = session.Table(where.TableName()).Where(where).Find(list)
	if len(count) > 0 {
		session = session.Offset(0).Count(count[0])
	}
	if err := session.Error; err != nil {
		logger.Errorf(ctx, "QueryInfoDB failed. [%v]", err)
		return errors.InternalServerError(service.GetName(), "QueryInfoDB fail. [%v]", err)
	}
	return nil
}

// QueryInfoDetailDB defined TODO
func (h *Handler) QueryInfoDetailDB(ctx context.Context, session *gorm.DB, where *model.Info, data *model.Info) error {
	var lst []*model.Info
	err := h.QueryInfoDB(ctx, session, where, &lst)
	if err != nil {
		logger.Errorf(ctx, "QueryInfoDetailDB failed. [%s]", err.Error())
		return err
	}
	if len(lst) == 0 {
		logger.Warn(ctx, "QueryInfoDetailDB empty.")
		return errors.InternalServerError(service.GetName(), "QueryInfoDetailDB empty.")
	}
	*data = *lst[0]
	return nil
}

// InsertInfoDB defined TODO
func (h *Handler) InsertInfoDB(ctx context.Context, session *gorm.DB, data *model.Info) error {
	err := session.Create(data).Error
	if err != nil {
		logger.Errorf(ctx, "InsertInfoDB failed. [%s]", err.Error())
		return errors.InternalServerError(service.GetName(), "InsertInfoDB fail. [%v]", err)
	}
	return nil
}

// UpdateInfoDB defined TODO
func (h *Handler) UpdateInfoDB(ctx context.Context, session *gorm.DB, data *model.Info) error {
	err := session.Table(data.TableName()).Model(&data).Updates(&data).Error
	if err != nil {
		logger.Errorf(ctx, "UpdateInfoDB failed. [%s]", err.Error())
		return errors.InternalServerError(service.GetName(), "UpdateInfoDB fail. [%v]", err)
	}
	return nil
}

// SaveInfoDB defined TODO
func (h *Handler) SaveInfoDB(ctx context.Context, session *gorm.DB, data *model.Info) error {
	err := session.Save(data).Error
	if err != nil {
		logger.Errorf(ctx, "SaveInfoDB failed. [%s]", err.Error())
		return errors.InternalServerError(service.GetName(), "SaveInfoDB fail. [%v]", err)
	}
	return nil
}

// DeleteInfoDB defined TODO
func (h *Handler) DeleteInfoDB(ctx context.Context, session *gorm.DB, data *model.Info) error {
	err := session.Where(data).Delete(&data).Error
	if err != nil {
		logger.Errorf(ctx, "DeleteInfoDB failed. [%s]", err.Error())
		return errors.InternalServerError(service.GetName(), "DeleteInfoDB fail. [%v]", err)
	}
	return nil
}
