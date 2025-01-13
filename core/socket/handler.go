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
	pReg registry.ProcessorRegistry,
	hReg registry.HandlerRegistry,
	conStore protocol.ConnectionManager) func(*websocket.Conn) {
	return func(c *websocket.Conn) {

		rReg := protocol.NewRateLimiter()
		wCon := NewWeb(c, "processing", rReg, logger)
		conStore.AddConnection(wCon)
		logger.Debug("Added new connection", zap.Int("active", conStore.ConnectionCount()))

		defer func() {
			if err := recover(); err != nil {
				logger.Error("Recovered from fatal error in WebSocket registry", zap.Any("panic", err), zap.String("id", wCon.Identifier()))
			}
		}()
		defer func() {
			if err := wCon.Dispose(); err != nil {
				logger.Error("Error closing WebSocket connection", zap.Error(err), zap.String("id", wCon.Identifier()))
			}
		}()

		for {
			if err := handleMessage(c, wCon, pReg, hReg, rReg, logger); err != nil {
				if websocket2.IsUnexpectedCloseError(err) {
					if websocket2.IsUnexpectedCloseError(err, websocket2.CloseGoingAway, websocket.CloseNormalClosure) {
						logger.Warn("WebSocket connection closed unexpectedly", zap.Error(err), zap.String("id", wCon.Identifier()))
					}
					logger.Debug("WebSocket connection closed", zap.Error(err), zap.String("id", wCon.Identifier()))
					conStore.RemoveConnection(wCon.Identifier())
					break
				} else {
					logger.Error("Error handling WebSocket message", zap.Error(err), zap.String("id", wCon.Identifier()))
				}
			}
		}
	}
}

// handleMessage processes a single WebSocket message.
func handleMessage(
	c *websocket.Conn,
	wCon protocol.Connection,
	pReg registry.ProcessorRegistry,
	hReg registry.HandlerRegistry,
	rReg protocol.RateLimiter,
	logger *zap.Logger) error {

	_, msg, err := c.ReadMessage()
	if err != nil {
		return err
	}

	conLog := logger.With(zap.String("id", wCon.Identifier()))

	if err := processPacket(msg, wCon, pReg, hReg, rReg, conLog); err != nil {
		logger.Warn("Error processing packet", zap.Error(err))
	}

	return nil
}

// processPacket deserializes and processes a packet.
func processPacket(
	msg []byte,
	wCon protocol.Connection,
	pReg registry.ProcessorRegistry,
	hReg registry.HandlerRegistry,
	rReg protocol.RateLimiter,
	logger *zap.Logger) error {

	defer func() {
		if r := recover(); r != nil {
			logger.Error("Recovered from panic during packet processing", zap.Any("panic", r), zap.String("id", wCon.Identifier()))
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
			logger.Debug("rate limit exceeded on connection", zap.Uint16("header", pack.GetHeader()), zap.String("id", wCon.Identifier()))
			return nil
		}
	}

	err = hReg.Handle(comPacket, wCon)
	if err != nil {
		return err
	}

	return nil
}
