package handler

import (
	"github.com/go-redis/redis/v8"
)

type Handler struct {
	RedisClient *redis.Client
}
