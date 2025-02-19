package web

import (
	"comm/config"

	"go-micro.dev/v5/web"
)

var (
	DefaultService web.Service
)

// setupService configures the service
func setupService(service web.Service) {
	DefaultService = service
	config.SetupService(service)
}
