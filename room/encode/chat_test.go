package encode

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"testing"
)

// TestRoomChatSettings_Encode verifies the encoding process.
func TestRoomChatSettings_Encode(t *testing.T) {
	rcs := &RoomChatSettings{
		Mode:       ChatModeFreeFlow,
		Weight:     ChatBubbleWidthNormal,
		Speed:      ChatScrollSpeedNormal,
		Distance:   10,
		Protection: FloodFilterNormal,
	}

	pck := protocol.NewPacket(200)
	pck.AddBoolean(true)
	rcs.Encode(&pck)
	raw := pck.ToBytes()

	dec, err := protocol.FromBytes(raw)
	assert.NoError(t, err, "Raw decoding must not have an error")

	_, err = dec.ReadBoolean()
	decRcs := &RoomChatSettings{}
	err = decRcs.Decode(dec)
	assert.NoError(t, err, "Decoding must not have an error")
	assert.Equal(t, rcs.Mode, decRcs.Mode)
	assert.Equal(t, rcs.Weight, decRcs.Weight)
	assert.Equal(t, rcs.Speed, decRcs.Speed)
	assert.Equal(t, rcs.Distance, decRcs.Distance)
	assert.Equal(t, rcs.Protection, decRcs.Protection)
}
