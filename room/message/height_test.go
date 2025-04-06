package message

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"testing"
)

var hPck = &HeightMapRequestPacket{
	Width: 3,
	Total: 9,
	Heights: []int16{
		1, 2, 3, // Y=0, X=0-2
		4, 5, 6, // Y=1, X=0-2
		7, 8, 9, // Y=2, X=0-2
	},
}

// parseHeight reads from raw packet to transform into heightmap request packet.
func parseHeight(pck protocol.RawPacket) (*HeightMapRequestPacket, error) {

	w, err := pck.ReadInt()
	t, err := pck.ReadInt()

	hm := make([]int16, t)
	for i := 0; i < int(t); i++ {
		h, err := pck.ReadShort()
		if err != nil {
			hm[i] = 0
		} else {
			hm[i] = h
		}
	}

	return &HeightMapRequestPacket{
		Width:   w,
		Total:   t,
		Heights: hm,
	}, err

}

// TestHeightMapRequestPacket_Serialize checks if serialization is made correctly.
func TestHeightMapRequestPacket_Serialize(t *testing.T) {
	raw := hPck.Serialize()
	bytes := raw.ToBytes()
	pck, err := protocol.FromBytes(bytes)
	assert.NoError(t, err)
	heightPck, err := parseHeight(*pck)
	assert.NoError(t, err)
	assert.Equal(t, heightPck.Width, hPck.Width)
	assert.Equal(t, heightPck.Total, hPck.Total)

	for i := range heightPck.Heights {
		assert.Equal(t, heightPck.Heights[i], hPck.Heights[i])
	}
}

// TestHeightMapRequestPacket check packet integrity.
func TestHeightMapRequestPacket(t *testing.T) {
	assert.Equal(t, hPck.Id(), uint16(HeightMapRequestPacketCode))
	assert.Equal(t, hPck.Deadline(), uint(0))
	mn, mx := hPck.Rate()
	if mn != 0 || mx != 0 {
		t.Errorf("Packet.Rate() = (%d, %d), expected (0, 0)", mn, mx)
	}
}
