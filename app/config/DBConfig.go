package config

import (
	"github.com/spf13/viper"
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
	EIS       DataBaseConfig
	DebugMode bool
}

func New() *Config {
	return &Config{
		Abit: DataBaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Name:     viper.GetString("DB_NAME"),
			UserName: viper.GetString("DB_USERNAME"),
			Password: viper.GetString("DB_PASSWORD"),
		},
		EIS: DataBaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Name:     viper.GetString("DB_NAME_EIS"),
			UserName: viper.GetString("DB_USERNAME"),
			Password: viper.GetString("DB_PASSWORD"),
		},
		DebugMode: getEnvAsBool("DEBUG_MODE", true),
	}
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := viper.GetString(name)
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}
