package router

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// SetupRouter defines the gin server for API and
// WebSocket engine as primary communication with
// Nitro Client.
func SetupRouter(logger *zap.Logger) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		ServerHeader:          "Pixels Emulator",
		CaseSensitive:         false,
		DisableStartupMessage: true,
	})
	app.Use(fiberzap.New(fiberzap.Config{Logger: logger}))
	return app, nil
}
