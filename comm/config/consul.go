package config

import (
	"comm/logger"
	"time"

	sourceConsul "github.com/micro/plugins/v5/source/consul"
	"go-micro.dev/v5/config"
	"go-micro.dev/v5/config/source/env"
)

func Bool(path ...string) bool {
	return config.Get(path...).Bool(false)
}

func Int(path ...string) int {
	return config.Get(path...).Int(0)
}
func String(path ...string) string {
	return config.Get(path...).String("")
}

func Float64(path ...string) float64 {
	return config.Get(path...).Float64(0)
}

func Duration(path ...string) time.Duration {
	return config.Get(path...).Duration(0)
}

func StringSlice(path ...string) []string {
	return config.Get(path...).StringSlice([]string{})
}

func StringMap(path ...string) map[string]string {
	return config.Get(path...).StringMap(map[string]string{})
}

func Bytes(path ...string) []byte {
	return config.Get(path...).Bytes()
}

func init() {
	err := config.Load(env.NewSource(
		env.WithPath("./.env"),
	))
	if err != nil {
		logger.Fatalf("Error source load: %v", err)
	}
	consulAddress := String("consul")
	if len(consulAddress) == 0 {
		logger.Fatal("Canot find consulAddress")
	}
	err = config.Load(sourceConsul.NewSource(
		sourceConsul.WithAddress(consulAddress),
		sourceConsul.WithPrefix("/micro/config"),
		sourceConsul.StripPrefix(false),
	))
	if err != nil {
		logger.Fatalf("Error source load: %v", err)
	}
}
