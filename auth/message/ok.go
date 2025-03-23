package message

import "pixels-emulator/core/protocol"

// AuthOkCode is the unique identifier for the Auth OK message.
const AuthOkCode = 2491

// AuthOkPacket represents a Ping packet used to verify the status of a connection.
// This type embeds the base protocol.Packet interface.
type AuthOkPacket struct {
	protocol.Packet // Embeds the base protocol.Packet interface.
}

// Id returns the unique identifier of the Packet type.
func (p *AuthOkPacket) Id() uint16 {
	return AuthOkCode
}

// Rate returns the rate limit for the Ping packet.
func (p *AuthOkPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *AuthOkPacket) Deadline() uint {
	return 10
}

// Serialize converts the Auth OK packet into a RawPacket that can be transmitted over the network.
func (p *AuthOkPacket) Serialize() protocol.RawPacket {
	return protocol.NewPacket(AuthOkCode)
}

// NewAuthOkPacket creates a new instance of Auth OK packet.
func NewAuthOkPacket() *AuthOkPacket {
	return &AuthOkPacket{}
}
