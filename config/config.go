package config

import "os"

type Config struct {
	Port    string
	Version string
}

func New() *Config {
	return &Config{
		Port:    getenv("PORT", "8081"),
		Version: getenv("VERSION", "1.0"),
	}
}

func getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
