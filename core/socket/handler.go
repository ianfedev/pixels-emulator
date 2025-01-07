package socket

import (
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
)

// Handle manages the basic message reception from websocket.
func Handle(logger *zap.Logger) func(*websocket.Conn) {
	return func(c *websocket.Conn) {
		wCon := &WebConnection{Socket: c}

		connLogger := logger.With(zap.String("identifier", wCon.Identifier()))

		defer func() {
			if err := recover(); err != nil {
				connLogger.Error("Recovered from fatal error in WebSocket handler", zap.Any("panic", err))
			}
		}()
		defer func() {
			if err := wCon.Dispose(); err != nil {
				connLogger.Error("Error closing WebSocket connection", zap.Error(err))
			}
		}()

		for {
			if err := handleMessage(c, wCon, connLogger); err != nil {
				if websocket.IsUnexpectedCloseError(err) {
					connLogger.Warn("WebSocket connection closed unexpectedly", zap.Error(err))
					break
				}
				connLogger.Error("Error handling WebSocket message", zap.Error(err))
			}
		}
	}
}

// handleMessage processes a single WebSocket message.
func handleMessage(c *websocket.Conn, wCon *WebConnection, logger *zap.Logger) error {
	_, msg, err := c.ReadMessage()
	if err != nil {
		return err
	}

	if err := processPacket(msg, wCon, logger); err != nil {
		logger.Warn("Error processing packet", zap.Error(err))
	}

	return nil
}

// processPacket deserializes and processes a packet.
func processPacket(msg []byte, wCon *WebConnection, logger *zap.Logger) error {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Recovered from panic during packet processing", zap.Any("panic", r))
		}
	}()

	pack, err := protocol.FromBytes(msg)
	if err != nil {
		return err
	}

	logger.Debug("Packet received", zap.Uint16("header", pack.GetHeader()))
	// Additional handling logic here.
	return nil
}
