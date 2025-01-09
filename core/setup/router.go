package setup

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"
	"pixels-emulator/core/registry"
	"pixels-emulator/core/socket"
)

// Router defines the gin healthcheck for API and
// WebSocket engine as primary communication with
// Nitro Client.
func Router(
	logger *zap.Logger,
	registry *registry.ProcessorRegistry,
	handlerRegistry *registry.Registry) (*fiber.App, error) {

	app := fiber.New(fiber.Config{
		ServerHeader:          "Pixels Emulator",
		CaseSensitive:         false,
		DisableStartupMessage: true,
	})

	app.Use(fiberzap.New(fiberzap.Config{Logger: logger}))
	app.Get("/", websocket.New(socket.Handle(logger, registry, handlerRegistry)))

	return app, nil

}
