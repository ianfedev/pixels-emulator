package socket

import (
	websocket2 "github.com/fasthttp/websocket"
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
)

// Handle manages the basic message reception from websocket.
func Handle(
	logger *zap.Logger,
	pReg *registry.ProcessorRegistry,
	hReg *registry.Registry) func(*websocket.Conn) {
	return func(c *websocket.Conn) {

		rReg := protocol.NewRateLimiterRegistry()
		wCon := NewWeb(c, "authenticating", rReg, logger)
		connLogger := logger.With(zap.String("identifier", wCon.Identifier()))

		defer func() {
			if err := recover(); err != nil {
				connLogger.Error("Recovered from fatal error in WebSocket registry", zap.Any("panic", err))
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
				hReg,
				rReg,
				connLogger); err != nil {
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
	pReg *registry.ProcessorRegistry,
	hReg *registry.Registry,
	rReg *protocol.RateLimiterRegistry,
	logger *zap.Logger) error {
	_, msg, err := c.ReadMessage()
	if err != nil {
		return err
	}

	if err := processPacket(msg, wCon, pReg, hReg, rReg, logger); err != nil {
		logger.Warn("Error processing packet", zap.Error(err))
	}

	return nil
}

// processPacket deserializes and processes a packet.
func processPacket(
	msg []byte,
	wCon *WebConnection,
	pReg *registry.ProcessorRegistry,
	hReg *registry.Registry,
	rReg *protocol.RateLimiterRegistry,
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

	period, rate := comPacket.Rate()

	if rate > 0 {
		limiter := rReg.GetLimiter(comPacket.Id(), period, rate)

		if !limiter.Allow() {
			logger.Debug("rate limit exceeded on connection", zap.Uint16("header", pack.GetHeader()))
			return nil
		}
	}

	err = hReg.Handle(comPacket, wCon)
	if err != nil {
		return err
	}

	return nil
}
