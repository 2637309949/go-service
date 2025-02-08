package db

import (
	"comm/logger"
	"context"
	"errors"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	defaultDB = &DB{dbs: make(map[string][]*gorm.DB)}
)

type order interface {
	GetOrderType() int32
	GetOrderCol() string
}

type DB struct {
	initdbs sync.Once
	dbs     map[string][]*gorm.DB
	od      []string
	conns   []ConnDb
	typ     string
}

type Options struct {
	User   string
	Passwd string
	Host   string
	Port   uint16
	Db     string
}

type ConnDb struct {
	Conn string
	Db   string
}

func (s *DB) setConns(typ string, opts []Options) {
	s.typ = typ
	for _, o := range opts {
		var c ConnDb
		if s.typ == "mysql" {
			c.Conn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", o.User, o.Passwd, o.Host, o.Port, o.Db)
			c.Db = o.Db
			s.conns = append(s.conns, c)
		} else if s.typ == "postgres" {
			c.Conn = fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", o.User, o.Passwd, o.Host, o.Port, o.Db)
			c.Db = o.Db
			s.conns = append(s.conns, c)
		} else {
			logger.Fatal("Connction type err")
		}
	}
}

func (s *DB) init() {
	for _, conn := range s.conns {
		db, err := gorm.Open(s.typ, conn.Conn)
		if err != nil {
			logger.Errorf("DB::Open failed %+v", err)
			return
		}
		db = db.Set("gorm:table_Options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
		db.DB().SetMaxIdleConns(20)
		db.DB().SetMaxOpenConns(20)
		db.LogMode(true)

		if _, ok := s.dbs[conn.Db]; !ok {
			s.od = append(s.od, conn.Db)
		}

		s.dbs[conn.Db] = append(s.dbs[conn.Db], db)

		logger.Infof("Init success, name:[%s] [%p]", conn.Db, db)
	}
}

func (s *DB) initDb(ctx context.Context, keys ...uint64) (*gorm.DB, error) {
	var key uint64
	if len(keys) > 0 {
		key = keys[0]
	}

	s.initdbs.Do(s.init)
	if len(s.od) == 0 {
		errmessage := "init err, dbs nil"
		logger.Error("err:" + errmessage)
		return nil, errors.New(errmessage)
	}

	dbname := s.od[0]
	gormDb := s.dbs[dbname]
	index := int(key % uint64(len(gormDb)))
	db := gormDb[index]
	db = setlogger(ctx, db)

	return db, nil
}

func (s *DB) initDbAssDb(ctx context.Context, dbname string, keys ...uint64) (*gorm.DB, error) {
	var key uint64
	if len(keys) > 0 {
		key = keys[0]
	}

	s.initdbs.Do(s.init)

	var db *gorm.DB
	gormDb, ok := s.dbs[dbname]
	if ok {
		index := int(key % uint64(len(gormDb)))
		db = gormDb[index]
	} else {
		errmessage := fmt.Sprintf("there has no this db:[%s]", dbname)
		logger.Error("err." + errmessage)
		return nil, errors.New(errmessage)
	}

	db = setlogger(ctx, db)
	return db, nil
}

func setlogger(ctx context.Context, db *gorm.DB) *gorm.DB {
	logger := logger.Extract(ctx)
	db.SetLogger(logger)
	return db
}
