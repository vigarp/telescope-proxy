package main

import (
	"fmt"
	"telescope-proxy/internal/config"

	"github.com/gofiber/fiber/v3"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	app := config.NewFiber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		App:    app,
		Log:    log,
		Config: viperConfig,
	})

	webPort := viperConfig.GetInt("PORT")
	log.Infof("🔭 Teleskop Proxy listening on port %d", webPort)

	err := app.Listen(fmt.Sprintf(":%d", webPort), fiber.ListenConfig{
		EnablePrefork: viperConfig.GetBool("WEB_PREFORK"),
	})
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
