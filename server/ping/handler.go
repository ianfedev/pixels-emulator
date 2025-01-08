package ping

import (
	"fmt"
	"pixels-emulator/core/protocol/registry"
)

type ClientMessageHandler struct {
	registry.Handler[ClientPacket]
}

func (h ClientMessageHandler) Handle(packet ClientPacket) {
	fmt.Println(packet)
}

func NewHandler() registry.Handler[ClientPacket] {
	return ClientMessageHandler{}
}
