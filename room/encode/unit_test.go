package encode

import (
	"pixels-emulator/room/unit"
	"testing"

	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
)

// mockUnitMessage returns a sample UnitMessage for test cases.
func mockUnitMessage() *UnitMessage {
	return &UnitMessage{
		Id:   1,
		X:    2,
		Y:    3,
		Z:    4,
		Head: 2,
		Body: 4,
		Status: map[unit.Status]string{
			"mv":  "2",
			"sit": "1",
		},
	}
}

// TestEncodeDecodeUnitMessage validates the encode/decode cycle of UnitMessage.
func TestEncodeDecodeUnitMessage(t *testing.T) {
	original := mockUnitMessage()
	packet := protocol.NewPacket(1)
	original.Encode(&packet)

	decoded := &UnitMessage{}
	err := decoded.Decode(&packet)

	assert.NoError(t, err, "decoding should not return error")
	assert.Equal(t, original.Id, decoded.Id)
	assert.Equal(t, original.X, decoded.X)
	assert.Equal(t, original.Y, decoded.Y)
	assert.Equal(t, original.Z, decoded.Z)
	assert.Equal(t, original.Head, decoded.Head)
	assert.Equal(t, original.Body, decoded.Body)
	assert.Equal(t, len(original.Status), len(decoded.Status))

	for k, v := range original.Status {
		assert.Equal(t, v, decoded.Status[k], "status value mismatch for key %s", k)
	}
}

// TestEmptyStatus checks decoding behavior when the status map is empty.
func TestEmptyStatus(t *testing.T) {
	msg := &UnitMessage{
		Id:     5,
		X:      6,
		Y:      7,
		Z:      8,
		Head:   1,
		Body:   3,
		Status: map[unit.Status]string{},
	}

	packet := protocol.NewPacket(1)
	msg.Encode(&packet)

	decoded := &UnitMessage{}
	err := decoded.Decode(&packet)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(decoded.Status), "expected status map to be empty")
}
