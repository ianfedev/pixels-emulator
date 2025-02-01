package message

import (
	"reflect"
	"testing"

	"pixels-emulator/core/protocol"
)

// TestNewAuthOkPacket verifies that NewAuthOkPacket returns a valid instance.
func TestNewAuthOkPacket(t *testing.T) {
	packet := NewAuthOkPacket()
	if packet == nil {
		t.Fatal("NewAuthOkPacket returned nil, expected a valid instance")
	}
}

// TestAuthOkPacketId verifies that the Id method returns the correct code.
func TestAuthOkPacketId(t *testing.T) {
	packet := NewAuthOkPacket()
	if got, want := packet.Id(), uint16(AuthOkCode); got != want {
		t.Errorf("AuthOkPacket.Id() = %d, expected %d", got, want)
	}
}

// TestAuthOkPacketRate verifies that the Rate method returns (0, 0).
func TestAuthOkPacketRate(t *testing.T) {
	packet := NewAuthOkPacket()
	mn, mx := packet.Rate()
	if mn != 0 || mx != 0 {
		t.Errorf("AuthOkPacket.Rate() = (%d, %d), expected (0, 0)", mn, mx)
	}
}

// TestAuthOkPacketSerialize verifies that the Serialize method returns a RawPacket with the correct code.
func TestAuthOkPacketSerialize(t *testing.T) {
	packet := NewAuthOkPacket()
	serialized := packet.Serialize()
	expected := protocol.NewPacket(AuthOkCode)
	if !reflect.DeepEqual(serialized, expected) {
		t.Errorf("AuthOkPacket.Serialize() = %+v, expected %+v", serialized, expected)
	}
}
