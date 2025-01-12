package socket

import (
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"sync"
)

// WebConnection wraps a websocket connection and adds additional functionalities such as unique identification and packet handling.
type WebConnection struct {
	protocol.Connection

	// Socket wraps the websocket connection.
	Socket *websocket.Conn

	// Id provides a unique id for the connection.
	Id string

	// logger to write messages.
	logger *zap.Logger

	// limiter provides rate limiter for outgoing packets.
	limiter *protocol.RateLimiterRegistry

	// writeMutex ensures thread-safe writes to the websocket.
	writeMutex sync.Mutex
}

// Dispose closes the websocket connection.
func (w *WebConnection) Dispose() error {
	return w.Socket.Close()
}

// Identifier returns the unique id for the WebConnection.
func (w *WebConnection) Identifier() string {
	return w.Id
}

// GrantIdentifier provides a new identifier for connection.
func (w *WebConnection) GrantIdentifier(id string) {
	if w.Id == "processing" {
		w.Id = id
	}
}

// SendPacket serializes the provided packet and sends it over the websocket connection.
// Logs an error if the sending process fails.
func (w *WebConnection) SendPacket(packet protocol.Packet) {
	period, rate := packet.Rate()
	w.SendRaw(packet.Serialize(), period, rate)
}

// SendRaw sends a packet over the websocket connection.
// Logs an error if the sending process fails.
func (w *WebConnection) SendRaw(packet protocol.RawPacket, period uint16, rate uint16) {
	conLog := w.logger.With(zap.Uint16("header", packet.GetHeader()), zap.String("identifier", w.Identifier()))

	if rate > 0 {
		limiter := w.limiter.GetLimiter(packet.GetHeader(), period, rate)

		if !limiter.Allow() {
			w.logger.Debug("rate limit exceeded on connection")
			return
		}
	}

	w.writeMutex.Lock()
	defer w.writeMutex.Unlock()

	err := w.Socket.WriteMessage(2, packet.ToBytes())
	if err != nil {
		conLog.Error("Error while processing packet for send", zap.Error(err))
		return
	}

	conLog.Debug("Packet sent")
}

// NewWeb creates a new WebConnection wrapper for a given websocket connection, unique id, and logger.
func NewWeb(socket *websocket.Conn, id string, limiter *protocol.RateLimiterRegistry, logger *zap.Logger) protocol.Connection {
	return &WebConnection{
		Socket:     socket,
		Id:         id,
		logger:     logger,
		limiter:    limiter,
		writeMutex: sync.Mutex{},
	}
}
