package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
)

// TestNewNavigatorMetaDataPacket verifies that NewNavigatorMetaDataPacket returns a valid instance.
func TestNewNavigatorMetaDataPacket(t *testing.T) {
	contexts := []string{"My world", "Popular rooms", "Events", "Categories"}
	packet := NewNavigatorMetaDataPacket(contexts...)
	assert.NotNil(t, packet, "NewNavigatorMetaDataPacket returned nil, expected a valid instance")
	assert.Equal(t, contexts, packet.Contexts, "NewNavigatorMetaDataPacket.Contexts mismatch")
}

// TestNavigatorMetaDataPacketId verifies that the Id method returns the correct code.
func TestNavigatorMetaDataPacketId(t *testing.T) {
	packet := NewNavigatorMetaDataPacket()
	assert.Equal(t, uint16(NavigatorMetaDataCode), packet.Id(), "NavigatorMetaDataPacket.Id() mismatch")
}

// TestNavigatorMetaDataPacketRate verifies that the Rate method returns (0, 0).
func TestNavigatorMetaDataPacketRate(t *testing.T) {
	packet := NewNavigatorMetaDataPacket()
	mn, mx := packet.Rate()
	assert.Equal(t, uint16(0), mn, "NavigatorMetaDataPacket.Rate() min mismatch")
	assert.Equal(t, uint16(0), mx, "NavigatorMetaDataPacket.Rate() max mismatch")
}

// TestNavigatorMetaDataPacketSerialize verifies that the Serialize method correctly encodes the packet.
func TestNavigatorMetaDataPacketSerialize(t *testing.T) {
	contexts := []string{"My world", "Popular rooms", "Events", "Categories"}
	packet := NewNavigatorMetaDataPacket(contexts...)
	serialized := packet.Serialize()

	expected := protocol.NewPacket(NavigatorMetaDataCode)
	expected.AddInt(int32(len(contexts)))
	for _, ctx := range contexts {
		expected.AddString(ctx)
		expected.AddInt(0)
	}

	assert.Equal(t, expected, serialized, "NavigatorMetaDataPacket.Serialize() mismatch")
}
