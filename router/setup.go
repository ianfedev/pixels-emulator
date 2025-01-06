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
				logger.Error("Error flushing out websocket, be aware", zap.Error(err))
			}
		}()

		for {

			_, msg, err := c.ReadMessage()
			if err != nil {
				logger.Error("Failed to read message from socket", zap.Error(err))
			}

			func() {

				defer func() {
					if r := recover(); r != nil {
						logger.Error("Recovered from fatal error deserialization", zap.Any("panic", r))
					}
				}()

				fmt.Println(msg)
				pack, packErr := packet.FromBytes(msg)
				logger.Debug("Packet received:", zap.Uint16("header", pack.GetHeader()))

				if packErr != nil {
					logger.Warn("Received unserializable packet", zap.Error(packErr))
				} else {
					// Game logic
				}

			}()
		}

	}))

	return app, nil
}
