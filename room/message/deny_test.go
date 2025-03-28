package message

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"testing"
)

var testPck = &DenyRoomConnectionPacket{
	Type:        Banned,
	QueryHolder: "test",
}

// parse reads from raw packet to transform into denying packet.
func parse(pck protocol.RawPacket) (*DenyRoomConnectionPacket, error) {
	t, err := pck.ReadInt()
	qh, err := pck.ReadString()
	return &DenyRoomConnectionPacket{
		Type:        ReasonType(t),
		QueryHolder: qh,
	}, err
}

// TestDenyRoomConnectionPacket_Serialize checks if serialization is made correctly.
func TestDenyRoomConnectionPacket_Serialize(t *testing.T) {
	raw := testPck.Serialize()
	bytes := raw.ToBytes()
	pck, err := protocol.FromBytes(bytes)
	assert.NoError(t, err)
	denyPck, err := parse(*pck)
	assert.NoError(t, err)
	assert.Equal(t, denyPck.Id(), uint16(DenyRoomConnectionCode))
	assert.Equal(t, denyPck.Type, testPck.Type)
	assert.Equal(t, denyPck.QueryHolder, testPck.QueryHolder)
}

// TestDenyRoomConnectionPacket check packet attributes.
func TestDenyRoomConnectionPacket(t *testing.T) {
	assert.Equal(t, testPck.Id(), uint16(DenyRoomConnectionCode))
	assert.Equal(t, testPck.Deadline(), uint(0))
	mn, mx := testPck.Rate()
	if mn != 0 || mx != 0 {
		t.Errorf("AuthOkPacket.Rate() = (%d, %d), expected (0, 0)", mn, mx)
	}
}
