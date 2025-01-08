package socket

import (
	websocket2 "github.com/fasthttp/websocket"
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"
	"pixels-emulator/core/handler"
	"pixels-emulator/core/protocol"
)

// Handle manages the basic message reception from websocket.
func Handle(
	logger *zap.Logger,
	pReg *handler.ProcessorRegistry,
	hReg *handler.Registry) func(*websocket.Conn) {
	return func(c *websocket.Conn) {

		wCon := &WebConnection{Socket: c, Identifier: "authenticating"}
		connLogger := logger.With(zap.String("identifier", wCon.GetIdentifier()))

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
			if err := handleMessage(
				c,
				wCon,
				pReg,
				hReg, connLogger); err != nil {
				if websocket2.IsUnexpectedCloseError(err) {
					if websocket2.IsUnexpectedCloseError(err, websocket2.CloseGoingAway, websocket.CloseNormalClosure) {
						connLogger.Warn("WebSocket connection closed unexpectedly", zap.Error(err))
					}
					connLogger.Debug("WebSocket connection closed", zap.Error(err))
					break
				} else {
					connLogger.Error("Error handling WebSocket message", zap.Error(err))
				}
			}
		}
	}
}

// handleMessage processes a single WebSocket message.
func handleMessage(
	c *websocket.Conn,
	wCon *WebConnection,
	pReg *handler.ProcessorRegistry,
	hReg *handler.Registry,
	logger *zap.Logger) error {
	_, msg, err := c.ReadMessage()
	if err != nil {
		return err
	}

	if err := processPacket(msg, wCon, pReg, hReg, logger); err != nil {
		logger.Warn("Error processing packet", zap.Error(err))
	}

	return nil
}

// processPacket deserializes and processes a packet.
func processPacket(
	msg []byte,
	wCon *WebConnection,
	pReg *handler.ProcessorRegistry,
	hReg *handler.Registry,
	logger *zap.Logger) error {
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
	comPacket, err := pReg.Handle(*pack, wCon)
	if err != nil {
		return err
	}

	err = hReg.Handle(comPacket, wCon)
	if err != nil {
		return err
	}

	return nil
}
