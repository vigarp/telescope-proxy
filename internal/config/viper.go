package config

import (
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetDefault("APP_NAME", "Telescope Proxy")
	config.SetDefault("WEB_PREFORK", false)
	config.SetDefault("PORT", 3000)
	config.SetDefault("LOG_LEVEL", 6)
	config.SetDefault("UPSTREAM_HOST", "example.com")
	config.SetDefault("SOCKS_TARGETS", "")
	config.SetDefault("CURL_BINARY", "curl")
	config.SetDefault("CONNECT_TIMEOUT_SEC", 3)
	config.SetDefault("EXEC_TIMEOUT_MS", 15000)

	config.SetConfigName(".env")
	config.SetConfigType("env")
	config.AddConfigPath("./../../")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")

	_ = config.ReadInConfig()

	config.AutomaticEnv()

	return config
}
