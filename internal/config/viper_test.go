package config

import (
	"testing"

	"net/http/httptest"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestNewViper_LoadConfig(t *testing.T) {
	v := NewViper()

	assert.NotEmpty(t, v.GetString("APP_NAME"))
	assert.Equal(t, 3000, v.GetInt("PORT"))
	assert.NotEmpty(t, v.GetString("UPSTREAM_HOST"))
	assert.Equal(t, "curl", v.GetString("CURL_BINARY"))
}

func TestNewViper_EnvOverride(t *testing.T) {
	t.Setenv("PORT", "4000")
	t.Setenv("UPSTREAM_HOST", "upstream.internal")
	t.Setenv("SOCKS_TARGETS", "proxy1.local:1080,proxy2.local:1080")

	v := NewViper()

	assert.Equal(t, 4000, v.GetInt("PORT"))
	assert.Equal(t, "upstream.internal", v.GetString("UPSTREAM_HOST"))
	assert.Equal(t, "proxy1.local:1080,proxy2.local:1080", v.GetString("SOCKS_TARGETS"))
}

func TestNewLoggerAndFiber(t *testing.T) {
	t.Setenv("APP_NAME", "Telescope Test App")

	v := NewViper()
	log := NewLogger(v)
	app := NewFiber(v)

	assert.NotNil(t, log)
	assert.NotNil(t, app)
	assert.Equal(t, "Telescope Test App", app.Config().AppName)
}

func TestBootstrapAndErrorHandler(t *testing.T) {
	v := NewViper()
	log := NewLogger(v)
	app := NewFiber(v)

	Bootstrap(&BootstrapConfig{
		App:    app,
		Log:    log,
		Config: v,
	})

	// Simulate route that return fiber.Error and regular error
	app.Get("/err-fiber", func(c fiber.Ctx) error {
		return fiber.ErrBadRequest
	})

	app.Get("/err-generic", func(c fiber.Ctx) error {
		return assert.AnError
	})

	req1 := httptest.NewRequest("GET", "/err-fiber", nil)
	res1, err := app.Test(req1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, res1.StatusCode)

	req2 := httptest.NewRequest("GET", "/err-generic", nil)
	res2, err := app.Test(req2)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, res2.StatusCode)
}
