package main

import (
	"go-micro.dev/v5/config"
	"go-micro.dev/v5/config/source/env"
	"go-micro.dev/v5/logger"
)

func init() {
	err := config.Load(env.NewSource())
	if err != nil {
		logger.Fatalf("Error source load: %v", err)
	}
}
