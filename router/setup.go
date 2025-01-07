package router

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"
	"pixels-emulator/core/socket"
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
	app.Get("/", websocket.New(socket.Handle(logger)))

	return app, nil

}
