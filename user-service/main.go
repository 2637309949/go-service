package main

import (
	"comm/logger"
	"comm/service"
	pbUser "proto/user"
	"user/handler"
)

func main() {
	logger.Info("Starting server")

	service := service.NewService(
		service.Name("user"),
	)

	pbUser.RegisterHandlerHandler(service.Server(), new(handler.Handler))
	if err := service.Run(); err != nil {
		logger.Fatalf("Error configuring store table option: %v", err)
	}

	logger.Info("Stopping server")
}
