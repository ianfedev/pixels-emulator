package router

import (
	"fmt"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"
	"pixels-emulator/packet"
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

	app.Get("/", websocket.New(func(c *websocket.Conn) {

		defer func() {
			if err := c.Close(); err != nil {
				logger.Error("Error closing WebSocket connection", zap.Error(err))
			}
		}()

		defer func() {
			if r := recover(); r != nil {
				logger.Error("Recovered from fatal error in WebSocket handler", zap.Any("panic", r))
			}
		}()

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				// Log the error and break the loop
				logger.Error("Failed to read message from WebSocket", zap.Error(err))
				break
			}

			func() {
				defer func() {
					if r := recover(); r != nil {
						logger.Error("Recovered from fatal error during packet processing", zap.Any("panic", r))
					}
				}()

				fmt.Println(msg)
				pack, packErr := packet.FromBytes(msg)

				if packErr != nil {
					logger.Warn("Received unserializable packet", zap.Error(packErr))
					return
				}

				logger.Debug("Packet received:", zap.Uint16("header", pack.GetHeader()))
				// Game logic here
			}()
		}

		logger.Info("WebSocket connection closed")
	}))

	return app, nil
}
