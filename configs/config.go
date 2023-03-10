package configs

import (
	"os"
	"strings"
)

type config struct {
	port      string
	version   string
	masterKey string
	allowList []string
}

func Default() *config {
	return &config{
		port:      getEnvWithFallback("PORT", "8082"),
		version:   getEnvWithFallback("VERSION", "1.0"),
		masterKey: getEnv("MASTER_KEY"),
		allowList: getEnvAsListWithFallback("ALLOW_LIST", "*"),
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func getEnvWithFallback(key string, fallback string) string {
	value := getEnv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getEnvAsList(key string) []string {
	return strings.Split(getEnv(key), ",")
}

func getEnvAsListWithFallback(key string, fallback string) []string {
	value := getEnvAsList(key)
	if len(value) <= 1 {
		return []string{ fallback }
	}
	return value
}

func (_config *config) GetPort() string {
	return _config.port
}

func (_config *config) GetVersion() string {
	return _config.version
}

func (_config *config) GetMasterKey() string {
	return _config.masterKey
}

func (_config *config) GetAllowList() []string {
	return _config.allowList
}
