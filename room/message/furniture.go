package message

import "pixels-emulator/core/protocol"

// RoomFurnitureAliasCode is the unique identifier for the packet
const RoomFurnitureAliasCode = 3898

// RoomFurnitureAliasPacket is a packet sent by Nitro when the heightmap id is provided.
// INVESTIGATION: Difference of this and Entry.
type RoomFurnitureAliasPacket struct {
	protocol.Packet
}

// Id returns the unique identifier of the Packet type.
func (p *RoomFurnitureAliasPacket) Id() uint16 {
	return RoomFurnitureAliasCode
}

// Rate returns the rate limit for the Ping packet.
func (p *RoomFurnitureAliasPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *RoomFurnitureAliasPacket) Deadline() uint {
	return 500
}

// ComposeRoomFurnitureAliasPacket composes a new instance of the packet.
func ComposeRoomFurnitureAliasPacket(_ protocol.RawPacket) (*RoomFurnitureAliasPacket, error) {
	return &RoomFurnitureAliasPacket{}, nil

}
