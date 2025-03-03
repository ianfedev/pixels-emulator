package message

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"testing"
)

// TestComposeNavigatorInit verifies that ComposeNavigatorInit returns a valid instance.
func TestComposeNavigatorInit(t *testing.T) {
	packet := ComposeNavigatorInit(protocol.RawPacket{})
	assert.NotNil(t, packet, "ComposeNavigatorInit returned nil, expected a valid instance")
}

// TestNavigatorInitPacketId verifies that the Id method returns the correct code.
func TestNavigatorInitPacketId(t *testing.T) {
	packet := ComposeNavigatorInit(protocol.RawPacket{})
	assert.Equal(t, uint16(NavigatorInitCode), packet.Id(), "NavigatorInitPacket.Id() mismatch")
}

// TestNavigatorInitPacketRate verifies that the Rate method returns (0, 0).
func TestNavigatorInitPacketRate(t *testing.T) {
	packet := ComposeNavigatorInit(protocol.RawPacket{})
	mn, mx := packet.Rate()
	assert.Equal(t, uint16(0), mn, "NavigatorInitPacket.Rate() min mismatch")
	assert.Equal(t, uint16(0), mx, "NavigatorInitPacket.Rate() max mismatch")
}
