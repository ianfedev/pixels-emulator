package protocol_test

import (
	"bytes"
	"encoding/binary"
	"pixels-emulator/core/protocol"
	"testing"
)

func TestAddInt(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	packet.AddInt(42)

	expected := make([]byte, 4)
	binary.BigEndian.PutUint32(expected, 42)

	if !bytes.Equal(packet.GetContent(), expected) {
		t.Errorf("AddInt failed. Expected %v, got %v", expected, packet.GetContent())
	}
}

func TestAddShort(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	packet.AddShort(300)

	expected := make([]byte, 2)
	binary.BigEndian.PutUint16(expected, 300)

	if !bytes.Equal(packet.GetContent(), expected) {
		t.Errorf("AddShort failed. Expected %v, got %v", expected, packet.GetContent())
	}
}

func TestAddBoolean(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	packet.AddBoolean(true)

	expected := []byte{1}

	if !bytes.Equal(packet.GetContent(), expected) {
		t.Errorf("AddBoolean failed. Expected %v, got %v", expected, packet.GetContent())
	}
}

func TestAddString(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	packet.AddString("hello")

	expected := new(bytes.Buffer)
	_ = binary.Write(expected, binary.BigEndian, int16(len("hello")))
	expected.WriteString("hello")

	if !bytes.Equal(packet.GetContent(), expected.Bytes()) {
		t.Errorf("AddString failed. Expected %v, got %v", expected.Bytes(), packet.GetContent())
	}
}

func TestToBytes(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	packet.AddInt(42)

	result := packet.ToBytes()

	expected := new(bytes.Buffer)
	_ = binary.Write(expected, binary.BigEndian, int32(6))
	_ = binary.Write(expected, binary.BigEndian, uint16(0x1234))
	_ = binary.Write(expected, binary.BigEndian, int32(42))

	if !bytes.Equal(result, expected.Bytes()) {
		t.Errorf("ToBytes failed. Expected %v, got %v", expected.Bytes(), result)
	}
}

func TestReadInt(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	packet.AddInt(42)

	value, err := packet.ReadInt()
	if err != nil {
		t.Errorf("ReadInt returned an error: %v", err)
	}

	if value != 42 {
		t.Errorf("ReadInt failed. Expected 42, got %d", value)
	}
}

func TestReadShort(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	packet.AddShort(300)

	value, err := packet.ReadShort()
	if err != nil {
		t.Errorf("ReadShort returned an error: %v", err)
	}

	if value != 300 {
		t.Errorf("ReadShort failed. Expected 300, got %d", value)
	}
}

func TestReadBoolean(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	packet.AddBoolean(true)

	value, err := packet.ReadBoolean()
	if err != nil {
		t.Errorf("ReadBoolean returned an error: %v", err)
	}

	if !value {
		t.Errorf("ReadBoolean failed. Expected true, got %v", value)
	}
}

func TestReadString(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	packet.AddString("hello")

	value, err := packet.ReadString()
	if err != nil {
		t.Errorf("ReadString returned an error: %v", err)
	}

	if value != "hello" {
		t.Errorf("ReadString failed. Expected 'hello', got '%s'", value)
	}
}

func TestFromBytes(t *testing.T) {
	data := new(bytes.Buffer)
	_ = binary.Write(data, binary.BigEndian, int32(6))
	_ = binary.Write(data, binary.BigEndian, uint16(0x1234))
	_ = binary.Write(data, binary.BigEndian, int32(42))

	packet, err := protocol.FromBytes(data.Bytes())
	if err != nil {
		t.Errorf("FromBytes returned an error: %v", err)
	}

	if packet.GetHeader() != 0x1234 {
		t.Errorf("FromBytes failed. Expected header 0x1234, got 0x%x", packet.GetHeader())
	}

	value, err := packet.ReadInt()
	if err != nil {
		t.Errorf("ReadInt returned an error: %v", err)
	}

	if value != 42 {
		t.Errorf("ReadInt failed. Expected 42, got %d", value)
	}
}

func TestAddBooleanFalse(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	packet.AddBoolean(false)

	expected := []byte{0}

	if !bytes.Equal(packet.GetContent(), expected) {
		t.Errorf("AddBoolean(false) failed. Expected %v, got %v", expected, packet.GetContent())
	}
}

func TestResetOffset(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	packet.AddInt(42)

	_, _ = packet.ReadInt()
	packet.ResetOffset()

	value, err := packet.ReadInt()
	if err != nil {
		t.Errorf("ReadInt after ResetOffset returned an error: %v", err)
	}

	if value != 42 {
		t.Errorf("ReadInt after ResetOffset failed. Expected 42, got %d", value)
	}
}

func TestReadIntNotEnoughBytes(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	_, err := packet.ReadInt()

	if err == nil || err.Error() != "not enough bytes to read int" {
		t.Errorf("ReadInt should have returned error 'not enough bytes to read int', got %v", err)
	}
}

func TestReadShortNotEnoughBytes(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	_, err := packet.ReadShort()

	if err == nil || err.Error() != "not enough bytes to read short" {
		t.Errorf("ReadShort should have returned error 'not enough bytes to read short', got %v", err)
	}
}

func TestReadBooleanNotEnoughBytes(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	_, err := packet.ReadBoolean()

	if err == nil || err.Error() != "not enough bytes to read boolean" {
		t.Errorf("ReadBoolean should have returned error 'not enough bytes to read boolean', got %v", err)
	}
}

func TestReadStringNotEnoughBytes(t *testing.T) {
	packet := protocol.NewPacket(0x1234)
	packet.AddShort(10) // Indica que el string tiene 10 bytes, pero no lo añade.

	_, err := packet.ReadString()
	if err == nil || err.Error() != "not enough bytes to read string" {
		t.Errorf("ReadString should have returned error 'not enough bytes to read string', got %v", err)
	}
}

func TestFromBytesTooShort(t *testing.T) {
	data := []byte{0x00, 0x01}
	_, err := protocol.FromBytes(data)

	if err == nil || err.Error() != "data too short to be a valid packet" {
		t.Errorf("FromBytes should have returned error 'data too short to be a valid packet', got %v", err)
	}
}

func TestFromBytesLengthMismatch(t *testing.T) {
	data := new(bytes.Buffer)
	_ = binary.Write(data, binary.BigEndian, int32(10)) // Declara longitud incorrecta
	_ = binary.Write(data, binary.BigEndian, uint16(0x1234))

	_, err := protocol.FromBytes(data.Bytes())
	if err == nil || err.Error() != "length mismatch in packet" {
		t.Errorf("FromBytes should have returned error 'length mismatch in packet', got %v", err)
	}
}

func TestReadString_ErrorOnReadShort(t *testing.T) {
	// Crea un paquete con contenido insuficiente para un short
	packet := protocol.NewPacket(0x1234)
	packet.AddBoolean(true) // Agrega solo un byte, que no es suficiente para un short (2 bytes)

	// Intenta leer una string del paquete
	_, err := packet.ReadString()
	if err == nil || err.Error() != "not enough bytes to read short" {
		t.Errorf("Expected error 'not enough bytes to read short', got %v", err)
	}
}
