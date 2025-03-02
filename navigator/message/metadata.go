package message

import "pixels-emulator/core/protocol"

// NavigatorMetaDataCode is the unique identifier for the packet
const NavigatorMetaDataCode = 3052

// NavigatorMetaDataPacket represents a crucial packet for Nitro
// navigator.
//
// It must provide the top level contextual categories.
//
// INVESTIGATION: At the moment, I have not seen any non-static implementation, but it can be
// worked out in the future, for now, it will just return the default 4 base categories.
type NavigatorMetaDataPacket struct {
	protocol.Packet

	Contexts []string // Contexts sets a list of the top view contexts provided.

}

// Id returns the unique identifier of the Packet type.
func (p *NavigatorMetaDataPacket) Id() uint16 {
	return NavigatorMetaDataCode
}

// Rate returns the rate limit for the packet.
func (p *NavigatorMetaDataPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Serialize converts the Auth OK packet into a RawPacket that can be transmitted over the network.
func (p *NavigatorMetaDataPacket) Serialize() protocol.RawPacket {

	pck := protocol.NewPacket(NavigatorMetaDataCode)
	pck.AddInt(int32(len(p.Contexts)))
	for _, navCtx := range p.Contexts {
		pck.AddString(navCtx)
		pck.AddInt(0)
	}

	return pck

}

// NewNavigatorMetaDataPacket creates the packet from a set of contextual strings (e.g "My world").
func NewNavigatorMetaDataPacket(contexts ...string) *NavigatorMetaDataPacket {
	return &NavigatorMetaDataPacket{
		Contexts: contexts,
	}
}
