package config

import "go-micro.dev/v5"

var (
	DefaultService micro.Service
)

// SetupService configures the service
func SetupService(service micro.Service) {
	DefaultService = service
}

// The service name
func serviceName() string {
	return DefaultService.Name()
}

func Conf(path ...string) string {
	lst := append([]string{"micro", "config", "service", serviceName()}, path...)
	return String(lst...)
}

func ConfMap(path ...string) map[string]string {
	lst := append([]string{"micro", "config", "service", serviceName()}, path...)
	return StringMap(lst...)
}

func ConfSlice(path ...string) []string {
	lst := append([]string{"micro", "config", "service", serviceName()}, path...)
	return StringSlice(lst...)
}

func ConfBytes(path ...string) []byte {
	lst := append([]string{"micro", "config", "service", serviceName()}, path...)
	return Bytes(lst...)
}

func CommConf(path ...string) string {
	lst := append([]string{"micro", "config", "comm"}, path...)
	return String(lst...)
}
