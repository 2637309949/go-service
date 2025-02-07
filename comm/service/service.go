package service

import (
	"comm/config"

	"go-micro.dev/v5"
	"go-micro.dev/v5/client"
	"go-micro.dev/v5/server"
)

var (
	DefaultService micro.Service
)

// setupService configures the service
func setupService(service micro.Service) {
	DefaultService = service
	config.SetupService(service)
}

// The service name
func GetName() string {
	return DefaultService.Name()
}

// Options returns the current options
func GetOptions() micro.Options {
	return DefaultService.Options()
}

// Client is used to call services
func GetClient() client.Client {
	return DefaultService.Client()
}

// Server is for handling requests and events
func GetServer() server.Server {
	return DefaultService.Server()
}

func GetString() string {
	return DefaultService.String()
}
