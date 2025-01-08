package ping

import "pixels-emulator/core/protocol"

const ClientPacketCode = 4000

type ClientPacket struct {
	protocol.Packet
	Version string
}

func (p ClientPacket) GetId() uint16 {
	return ClientPacketCode
}

func NewPingPacket(packet protocol.RawPacket) (*ClientPacket, error) {
	ver, err := packet.ReadString()
	if err != nil {
		return nil, err
	}

	return &ClientPacket{
		Version: ver,
	}, nil
}
