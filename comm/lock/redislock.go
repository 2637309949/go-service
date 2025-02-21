package redislock

import (
	"comm/logger"
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
)

type RedisLocker struct {
	lockName string
	code     string
	timeOut  time.Duration
	expire   time.Duration
	client   *redis.Client
}

func New(client *redis.Client, lockName string, acquireTimeout, expire time.Duration) *RedisLocker {
	var locker RedisLocker
	locker.lockName = lockName
	locker.timeOut = acquireTimeout
	locker.expire = expire
	locker.client = client

	return &locker
}

func (l *RedisLocker) Lock(ctx context.Context) error {
	code := uuid.NewV4().String()
	endTime := time.Now().Add(l.timeOut).UnixNano()
	logger := logger.Extract(ctx)
	logger.Infof("now:%d end:%d", time.Now().UnixNano(), endTime)
	for time.Now().UnixNano() <= endTime {
		if success, err := l.client.SetNX(ctx, l.lockName, code, l.expire).Result(); err != nil && err != redis.Nil {
			return err
		} else if success {
			l.code = code
			return nil
		} else if l.client.TTL(ctx, l.lockName).Val() == -1 { //-2:失效；-1：无过期；
			l.client.Expire(ctx, l.lockName, l.expire)
		}
		time.Sleep(time.Millisecond * 10)
	}
	return errors.New("timeout")
}

// var count = 0  // test assist
func (l *RedisLocker) Release(ctx context.Context) bool {
	txf := func(tx *redis.Tx) error {
		if v, err := tx.Get(ctx, l.lockName).Result(); err != nil && err != redis.Nil {
			return err
		} else if v == l.code {
			_, err := tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
				//count++
				//fmt.Println(count)
				pipe.Del(ctx, l.lockName)
				return nil
			})
			return err
		}
		return nil
	}

	for {
		if err := l.client.Watch(ctx, txf, l.lockName); err == nil {
			return true
		} else if err == redis.TxFailedErr {
			logger.Errorf("watch key is modified, retry to release lock. err:%v", err.Error())
		} else {
			logger.Errorf("err:%v", err.Error())
			return false
		}
	}
}
