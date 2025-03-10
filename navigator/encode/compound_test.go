package encode_test

import (
	"pixels-emulator/navigator/encode"
	"testing"

	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	roomEncode "pixels-emulator/room/encode"
)

func TestSearchResultCompoundEncodeDecode(t *testing.T) {
	r := encode.SearchResultCompound{
		Code:       "test_code",
		Query:      "test_query",
		Collapsed:  true,
		Actionable: true,
		Thumbnails: true,
		Rooms: []*roomEncode.RoomData{
			{
				ID:        1,
				Name:      "Room 1",
				OwnerID:   100,
				OwnerName: "Owner1",
			},
			{
				ID:        2,
				Name:      "Room 2",
				OwnerID:   101,
				OwnerName: "Owner2",
			},
		},
	}

	pck := protocol.NewPacket(1010)
	r.Encode(&pck)

	decodedResult := encode.SearchResultCompound{}
	err := decodedResult.Decode(&pck)
	assert.NoError(t, err, "Decode should not return an error")
	assert.Equal(t, r.Code, decodedResult.Code, "Code should match")
	assert.Equal(t, r.Query, decodedResult.Query, "Query should match")
	assert.Equal(t, r.Collapsed, decodedResult.Collapsed, "Collapsed should match")
	assert.Equal(t, r.Actionable, decodedResult.Actionable, "Actionable should match")
	assert.Equal(t, r.Thumbnails, decodedResult.Thumbnails, "Thumbnails should match")
	assert.Equal(t, len(r.Rooms), len(decodedResult.Rooms), "Number of rooms should match")
}
