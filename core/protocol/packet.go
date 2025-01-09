package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// RawPacket represents a data packet with a header and content
// received from Nitro Client.
type RawPacket struct {
	header  uint16 // header is the identifier of the packet defined by Nitro messaging.
	content []byte // content serves as raw and readable bytes from composed packets.
	offset  int    // offset indicates the actual point of buffer reading.
}

// Packet defines an already composed packet from raw protocol.
type Packet interface {

	// Id provides identifier for packet.
	Id() uint16

	// Serialize provides a raw packet version to be sent.
	Serialize() RawPacket

	// Rate provides the maximum of packets per second a connection can receive.
	Rate() (uint16, uint16)
}

// GetHeader obtains the header of the packet.
func (p *RawPacket) GetHeader() uint16 {
	return p.header
}

// GetContent obtains the content of the packet.
func (p *RawPacket) GetContent() []byte {
	return p.content
}

// AddInt adds a 4-byte integer to the packet content.
func (p *RawPacket) AddInt(value int32) {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, value)
	p.content = append(p.content, buf.Bytes()...)
}

// AddShort adds a 2-byte short integer to the packet content.
func (p *RawPacket) AddShort(value int16) {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, value)
	p.content = append(p.content, buf.Bytes()...)
}

// AddBoolean adds a 1-byte boolean value (true or false) to the packet content.
func (p *RawPacket) AddBoolean(value bool) {
	var b byte
	if value {
		b = 1
	} else {
		b = 0
	}
	p.content = append(p.content, b)
}

// AddString adds a UTF-8 string to the packet content, preceded by its length in bytes.
func (p *RawPacket) AddString(value string) {
	length := int16(len(value))
	p.AddShort(length)
	p.content = append(p.content, []byte(value)...)
}

// ResetOffset restarts from the beginning the reading offset.
func (p *RawPacket) ResetOffset() {
	p.offset = 0
}

// ToBytes converts the packet to a byte slice, including the header and content.
func (p *RawPacket) ToBytes() []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, int32(len(p.content)+2)) // Length of the packet
	_ = binary.Write(buf, binary.BigEndian, p.header)                // Header
	buf.Write(p.content)                                             // Content
	return buf.Bytes()
}

// ReadInt reads a 4-byte integer from the packet content.
func (p *RawPacket) ReadInt() (int32, error) {
	if p.offset+4 > len(p.content) {
		return 0, errors.New("not enough bytes to read int")
	}
	var value int32
	buf := bytes.NewReader(p.content[p.offset : p.offset+4])
	_ = binary.Read(buf, binary.BigEndian, &value)
	p.offset += 4
	return value, nil
}

// ReadShort reads a 2-byte short integer from the packet content.
func (p *RawPacket) ReadShort() (int16, error) {
	if p.offset+2 > len(p.content) {
		return 0, errors.New("not enough bytes to read short")
	}
	var value int16
	buf := bytes.NewReader(p.content[p.offset : p.offset+2])
	_ = binary.Read(buf, binary.BigEndian, &value)
	p.offset += 2
	return value, nil
}

// ReadBoolean reads a 1-byte boolean from the packet content.
func (p *RawPacket) ReadBoolean() (bool, error) {
	if p.offset+1 > len(p.content) {
		return false, errors.New("not enough bytes to read boolean")
	}
	value := p.content[p.offset]
	p.offset++
	return value == 1, nil
}

// ReadString reads a string from the packet content. It expects the string to be preceded by a short indicating its length.
func (p *RawPacket) ReadString() (string, error) {
	length, err := p.ReadShort()
	if err != nil {
		return "", err
	}
	if p.offset+int(length) > len(p.content) {
		return "", errors.New("not enough bytes to read string")
	}
	value := string(p.content[p.offset : p.offset+int(length)])
	p.offset += int(length)
	return value, nil
}

// FromBytes converts a byte slice into a RawPacket.
func FromBytes(data []byte) (*RawPacket, error) {
	if len(data) < 6 {
		return nil, errors.New("data too short to be a valid packet")
	}
	buf := bytes.NewReader(data)
	var length int32
	_ = binary.Read(buf, binary.BigEndian, &length)
	if int(length)+4 != len(data) {
		return nil, errors.New("length mismatch in packet")
	}
	var header uint16
	_ = binary.Read(buf, binary.BigEndian, &header)
	content := data[6:]
	return &RawPacket{header: header, content: content, offset: 0}, nil
}

// NewPacket creates from scratch a packet.
func NewPacket(header uint16) RawPacket {
	return RawPacket{
		header: header,
		offset: 0,
	}
}
