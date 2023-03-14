package config

import (
	"os"
	"strconv"
)

type DataBaseConfig struct {
	Host     string
	Name     string
	UserName string
	Password string
}
type Config struct {
	Abit      DataBaseConfig
	DebugMode bool
}

func New() *Config {
	return &Config{
		Abit: DataBaseConfig{
			Host:     getEnv("DB_HOST", ""),
			Name:     getEnv("DB_NAME", ""),
			UserName: getEnv("DB_USERNAME", ""),
			Password: getEnv("DB_PASSWORD", ""),
		},
		DebugMode: getEnvAsBool("DEBUG_MODE", true),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
