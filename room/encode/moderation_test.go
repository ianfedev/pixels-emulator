package encode

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"testing"
)

// moderationRights is a simple mock
var moderationRights = &ModerationRights{
	Mute: None,
	Kick: Rights,
	Ban:  Administrator,
}

// TestModerationRights_Encode check if encode is done correctly.
func TestModerationRights_Encode(t *testing.T) {

	pck := protocol.NewPacket(100)
	pck.AddBoolean(true)
	moderationRights.Encode(&pck)
	raw := pck.ToBytes()

	dec, err := protocol.FromBytes(raw)
	assert.NoError(t, err, "Raw decoding must not have an error")

	_, err = dec.ReadBoolean()
	decRights := &ModerationRights{}
	err = decRights.Decode(dec)
	assert.NoError(t, err, "Decoding must not have an error")
	assert.Equal(t, moderationRights.Mute, decRights.Mute)
	assert.Equal(t, moderationRights.Kick, decRights.Kick)
	assert.Equal(t, moderationRights.Ban, decRights.Ban)

}
