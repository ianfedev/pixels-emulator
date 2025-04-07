package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var aliasPacket = &RoomFurnitureAliasPacket{}

// TestRoomFurnitureAliasPacket_Integrity tests packet ID, deadline, and rate.
func TestRoomFurnitureAliasPacket_Integrity(t *testing.T) {
	assert.Equal(t, aliasPacket.Id(), uint16(RoomFurnitureAliasCode))
	assert.Equal(t, aliasPacket.Deadline(), uint(500))

	mn, mx := aliasPacket.Rate()
	assert.Equal(t, uint16(0), mn)
	assert.Equal(t, uint16(0), mx)
}
