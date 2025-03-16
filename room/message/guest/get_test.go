package guest

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"testing"
)

// guestPacket is the mocked packet.
var guestPacket = &GetRoomPacket{
	RoomId:  1,
	Enter:   true,
	Forward: true,
}

// encodeGuestNavigator generates a decoding helper function to test.
func encodeGuestNavigator(dec *GetRoomPacket) *protocol.RawPacket {
	pck := protocol.NewPacket(GetGuestRoomCode)
	pck.AddInt(dec.RoomId)
	pck.AddBoolean(dec.Enter)
	pck.AddBoolean(dec.Forward)
	return &pck
}

// TestComposeGuestRoomPacket check if packet is encoded and decoded.
func TestComposeGuestRoomPacket(t *testing.T) {

	enc := encodeGuestNavigator(guestPacket)
	bytes := enc.ToBytes()

	newPck, err := protocol.FromBytes(bytes)
	assert.NoError(t, err, "Protocol byte parsing must not have error")

	dec, err := ComposeGuestRoomPacket(*newPck)
	assert.NoError(t, err, "Protocol decoding must not have error")
	assert.Equal(t, guestPacket.RoomId, dec.RoomId)
	assert.Equal(t, guestPacket.Enter, dec.Enter)
	assert.Equal(t, guestPacket.Forward, dec.Forward)
	assert.Equal(t, int(guestPacket.Id()), GetGuestRoomCode)

}

// TestComposeGuestRoomPacketPacketRate verifies that the Rate method returns (0, 0).
func TestComposeGuestRoomPacketPacketRate(t *testing.T) {
	mn, mx := guestPacket.Rate()
	assert.Equal(t, uint16(0), mn, "NavigatorInitPacket.Rate() min mismatch")
	assert.Equal(t, uint16(0), mx, "NavigatorInitPacket.Rate() max mismatch")
}
