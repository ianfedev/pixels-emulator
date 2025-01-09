package socket

import (
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
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
}

// Dispose closes the websocket connection.
func (w *WebConnection) Dispose() error {
	return w.Socket.Close()
}

// Identifier returns the unique id for the WebConnection.
func (w *WebConnection) Identifier() string {
	return w.Id
}

// SendPacket serializes the provided packet and sends it over the websocket connection.
// Logs an error if the sending process fails.
func (w *WebConnection) SendPacket(packet protocol.Packet) {
	sPacket := packet.Serialize()
	err := w.Socket.WriteMessage(2, sPacket.ToBytes())
	if err != nil {
		w.logger.Error("Error while processing packet for send", zap.Error(err))
		return
	}
	w.logger.Debug("Packet sent", zap.Uint16("header", packet.Id()), zap.String("identifier", w.Identifier()))
}

// NewWeb creates a new WebConnection wrapper for a given websocket connection, unique id, and logger.
func NewWeb(socket *websocket.Conn, id string, logger *zap.Logger) *WebConnection {
	return &WebConnection{
		Socket: socket,
		Id:     id,
		logger: logger,
	}
}
