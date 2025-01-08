package healthcheck

import "pixels-emulator/core/protocol"

const IncomingPingCode = 4000

type PingIncomingPacket struct {
	protocol.Packet
	Version string
}

func (p PingIncomingPacket) GetId() uint16 {
	return IncomingPingCode
}

func NewPingIncoming(packet protocol.RawPacket) (*PingIncomingPacket, error) {
	ver, err := packet.ReadString()
	if err != nil {
		return nil, err
	}

	return &PingIncomingPacket{
		Version: ver,
	}, nil
}
