package message

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"testing"
)

// enterPacket is a demo room packet.
var enterPacket = RoomEnterPacket{
	RoomId:   1,
	Password: "yes",
}

func encodeEnterPacket(dec *RoomEnterPacket) *protocol.RawPacket {
	pck := protocol.NewPacket(RoomEnterCode)
	pck.AddInt(dec.RoomId)
	pck.AddString(dec.Password)
	return &pck
}

func testPacketEncoding(t *testing.T, basePck *RoomEnterPacket) *RoomEnterPacket {
	enc := encodeEnterPacket(basePck)
	bytes := enc.ToBytes()
	pck, err := protocol.FromBytes(bytes)
	assert.NoError(t, err, "Raw decoding must not have error")
	roomPck, err := ComposeRoomEnterPacket(*pck)
	assert.NoError(t, err, "Decoding must not have error")
	assert.Equal(t, basePck.RoomId, roomPck.RoomId)
	assert.Equal(t, basePck.Password, roomPck.Password)
	return roomPck
}

func TestComposeRoomEnterPacket(t *testing.T) {
	testPacketEncoding(t, &enterPacket)
}

func TestComposeRoomEnterPacket_EmptyPassword(t *testing.T) {
	pck := testPacketEncoding(t, &RoomEnterPacket{
		RoomId:   30,
		Password: "",
	})
	assert.Equal(t, pck.Password, "")
}
