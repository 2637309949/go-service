package db

import (
	"comm/consts"
	"context"
	"reflect"

	"github.com/jinzhu/gorm"
)

type LimitOffset interface {
	GetOffset() int32
	GetLimit() int32
}

type LimitPage interface {
	GetPageSize() int32
	GetPageNo() int32
}

func InitPage(ctx context.Context, itf LimitPage) {
	pageSize := itf.GetPageSize()
	pageNo := itf.GetPageNo()

	if pageNo == 0 {
		pageNo = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	fieldPageNo := reflect.ValueOf(itf).Elem().FieldByName("PageNo")
	fieldPageSize := reflect.ValueOf(itf).Elem().FieldByName("PageSize")
	fieldPageNo.SetInt(int64(pageNo))
	fieldPageSize.SetInt(int64(pageSize))
}

func SetConns(typ string, opts []Options) {
	defaultDB.setConns(typ, opts)
}

func InitDb(ctx context.Context, keys ...uint64) (*gorm.DB, error) {
	return defaultDB.initDb(ctx, keys...)
}

func InitDbAssDb(ctx context.Context, dbname string, keys ...uint64) (*gorm.DB, error) {
	return defaultDB.initDbAssDb(ctx, dbname, keys...)
}

func SetLimit(ctx context.Context, db *gorm.DB, obj interface{}) *gorm.DB {
	if l, ok1 := obj.(LimitOffset); ok1 {
		if l.GetOffset() > 0 {
			db = db.Offset(l.GetOffset())
		}
		if l.GetLimit() > 0 {
			db = db.Limit(l.GetLimit())
		}
	} else if l, ok2 := obj.(LimitPage); ok2 {
		InitPage(ctx, l)
		db = db.Limit(l.GetPageSize())
		db = db.Offset(l.GetPageSize() * (l.GetPageNo() - 1))
	} else {
		db = db.Limit(20)
	}

	return db
}

func SetOrder(ctx context.Context, db *gorm.DB, o order, tb ...string) *gorm.DB {
	strOrder := o.GetOrderCol()
	if len(strOrder) > 0 {
		if len(tb) > 0 {
			strOrder = tb[0] + "." + strOrder
		}
		switch o.GetOrderType() {
		case consts.ORDER_ASC:
			strOrder += " ASC"
		case consts.ORDER_DESC:
			strOrder += " DESC"
		default:
			strOrder += " ASC"
		}

		db = db.Order(strOrder)
	}

	return db
}
