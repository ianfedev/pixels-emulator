package socket

import (
	"github.com/gofiber/websocket/v2"
	"pixels-emulator/core/protocol"
)

type WebConnection struct {
	protocol.Connection

	// Socket wraps the websocket connection.
	Socket *websocket.Conn

	// Identifier provides a unique id for the connection.
	Identifier string
}

func (w *WebConnection) Dispose() error {
	return w.Socket.Close()
}

func (w *WebConnection) GetIdentifier() string {
	return w.Identifier
}
