package main

import (
	"comm/config"
	"comm/service"
	pb "proto/cron"
	"test/handler"

	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	redisAddr := config.CommConf("redis_address")
	redisPasswd := config.CommConf("redis_passwd")
	service := service.NewService(
		service.Name("cron"),
	)
	h := handler.Handler{
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPasswd,
			DB:       0,
		}),
	}
	pb.RegisterCronServiceHandler(service.Server(), &h)
	go timed(c, &h)
	service.Run()
}

func timed(c *cron.Cron, h *handler.Handler) {
	c.AddFunc("*/5 * * * * *", h.TimedLog)
	c.Start()
}
