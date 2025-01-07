package socket

import (
	"github.com/gofiber/websocket/v2"
	"pixels-emulator/core/protocol"
)

type WebConnection struct {
	protocol.Connection

	// Socket wraps the websocket connection.
	Socket *websocket.Conn
}

func (w *WebConnection) Dispose() error {
	return w.Socket.Close()
}

func (w *WebConnection) Identifier() string {
	return "" // TODO: Identify.
}
