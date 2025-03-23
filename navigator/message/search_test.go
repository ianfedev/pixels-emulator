package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
)

// Test_NavigatorSearchPacket verifies the packet ID, rate, and parsing.
func Test_NavigatorSearchPacket(t *testing.T) {
	pck := protocol.NewPacket(NavigatorSearchCode)
	pck.AddString("hotel_view")
	pck.AddString("popular_rooms")

	navSearchPacket, err := ComposeNavigatorSearch(pck)

	assert.NoError(t, err)
	assert.NotNil(t, navSearchPacket)
	assert.Equal(t, uint16(NavigatorSearchCode), uint16(navSearchPacket.Id()))
	assert.Equal(t, "hotel_view", navSearchPacket.View)
	assert.Equal(t, "popular_rooms", navSearchPacket.Query)

	rateLimit, rateInterval := navSearchPacket.Rate()
	assert.Equal(t, uint16(5), rateLimit)
	assert.Equal(t, uint16(5), rateInterval)
}

// Test_ComposeNavigatorSearch_InvalidPacket ensures error handling for incomplete packets.
func Test_ComposeNavigatorSearch_InvalidPacket(t *testing.T) {
	// Case 1: Packet with only one string
	pck := protocol.NewPacket(NavigatorSearchCode)
	pck.AddString("only_one_value")

	_, err := ComposeNavigatorSearch(pck)
	assert.Error(t, err)

	// Case 2: Completely empty packet (to trigger first ReadString error)
	emptyPck := protocol.NewPacket(NavigatorSearchCode)
	_, err = ComposeNavigatorSearch(emptyPck)
	assert.Error(t, err)
}
