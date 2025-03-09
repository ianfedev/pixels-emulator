package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestData represents a structure used for testing Encodable functionality.
type TestData struct {
	ID    int32
	Name  string
	Valid bool
}

// Encode writes TestData into a RawPacket.
func (t *TestData) Encode(pck *RawPacket) {
	pck.AddInt(t.ID)
	pck.AddString(t.Name)
	pck.AddBoolean(t.Valid)
}

// Decode reads TestData from a RawPacket.
func (t *TestData) Decode(pck *RawPacket) error {
	var err error
	t.ID, err = pck.ReadInt()
	if err != nil {
		return err
	}

	t.Name, err = pck.ReadString()
	if err != nil {
		return err
	}

	t.Valid, err = pck.ReadBoolean()
	return err
}

// TestEncodable ensures that a structure implementing Encodable can be serialized and deserialized correctly.
func TestEncodable(t *testing.T) {
	original := TestData{
		ID:    12345,
		Name:  "Test Room",
		Valid: true,
	}

	pck := NewPacket(9999)
	pck.AddShort(42) // Extra data before Encodable
	original.Encode(&pck)
	pck.AddInt(9876) // Extra data after Encodable

	rawBytes := pck.ToBytes()

	// Deserialize the packet
	receivedPck, err := FromBytes(rawBytes)
	require.NoError(t, err, "Failed to parse RawPacket")

	beforeEncodable, err := receivedPck.ReadShort()
	require.NoError(t, err, "Failed to read beforeEncodable")
	assert.Equal(t, int16(42), beforeEncodable, "Incorrect beforeEncodable value")

	var decoded TestData
	require.NoError(t, decoded.Decode(receivedPck), "Failed to decode TestData")

	// Validate decoded data
	assert.Equal(t, original.ID, decoded.ID, "ID mismatch")
	assert.Equal(t, original.Name, decoded.Name, "Name mismatch")
	assert.Equal(t, original.Valid, decoded.Valid, "Valid flag mismatch")

	// Validate extra data after Encodable
	afterEncodable, err := receivedPck.ReadInt()
	require.NoError(t, err, "Failed to read afterEncodable")
	assert.Equal(t, int32(9876), afterEncodable, "Incorrect afterEncodable value")
}
