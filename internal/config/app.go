package config

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	App    *fiber.App
	Log    *logrus.Logger
	Config *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	config.Log.Infof("Bootstrapping %s...", config.Config.GetString("APP_NAME"))
}
