package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unicode/utf8"
)

// RawPacket represents a data packet with a header and content
// received from Nitro Client.
type RawPacket struct {
	header  uint16 // header is the identifier of the packet defined by Nitro messaging.
	content []byte // content serves as raw and readable bytes from composed packets.
	offset  int    // offset indicates the actual point of buffer reading.
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
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		fmt.Println("Error writing int:", err)
		return
	}
	p.content = append(p.content, buf.Bytes()...)
}

// AddShort adds a 2-byte short integer to the packet content.
func (p *RawPacket) AddShort(value int16) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		fmt.Println("Error writing short:", err)
		return
	}
	p.content = append(p.content, buf.Bytes()...)
}

// ResetOffset restarts from beginning the reading offset.
func (p *RawPacket) ResetOffset() {
	p.offset = 0
}

// AddBoolean adds a 1-byte boolean value (true or false) to the packet content.
func (p *RawPacket) AddBoolean(value bool) {
	var byteValue byte
	if value {
		byteValue = 1
	} else {
		byteValue = 0
	}
	p.content = append(p.content, byteValue)
}

// AddString adds a UTF-8 string to the packet content, preceded by its length in bytes.
func (p *RawPacket) AddString(value string) {
	// Get the length of the string in bytes
	length := int16(utf8.RuneCountInString(value))

	// Convert the length to big endian short
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, length)
	if err != nil {
		fmt.Println("Error writing string length:", err)
		return
	}

	// Convert the string to bytes
	p.content = append(p.content, buf.Bytes()...)
	p.content = append(p.content, []byte(value)...)
}

// ToBytes converts the packet to a byte slice, including the header and content.
func (p *RawPacket) ToBytes() []byte {

	// Create the packet by first adding the length (excluding the first 4 bytes)
	packetLength := int32(4 + len(p.content))
	buf := new(bytes.Buffer)

	// Write packet length in Big Endian
	err := binary.Write(buf, binary.BigEndian, packetLength)
	if err != nil {
		fmt.Println("Error writing packet length:", err)
		return nil
	}

	// Write header in Big Endian
	err = binary.Write(buf, binary.BigEndian, p.header)
	if err != nil {
		fmt.Println("Error writing header:", err)
		return nil
	}

	buf.Write(p.content)

	return buf.Bytes()
}

// ReadInt reads a 4-byte integer from the packet content.
func (p *RawPacket) ReadInt() (int32, error) {
	if p.offset+4 > len(p.content) {
		return 0, fmt.Errorf("not enough data to read int")
	}
	value := binary.BigEndian.Uint32(p.content[p.offset : p.offset+4])
	p.offset += 4
	return int32(value), nil
}

// ReadShort reads a 2-byte short integer from the packet content.
func (p *RawPacket) ReadShort() (int16, error) {
	if p.offset+2 > len(p.content) {
		return 0, fmt.Errorf("not enough data to read short")
	}
	value := binary.BigEndian.Uint16(p.content[p.offset : p.offset+2])
	p.offset += 2
	return int16(value), nil
}

// ReadBoolean reads a 1-byte boolean from the packet content.
func (p *RawPacket) ReadBoolean() (bool, error) {
	if p.offset+1 > len(p.content) {
		return false, fmt.Errorf("not enough data to read boolean")
	}
	value := p.content[p.offset]
	p.offset += 1
	return value == 1, nil
}

// ReadString reads a string from the packet content. It expects the string to be preceded by a short indicating its length.
func (p *RawPacket) ReadString() (string, error) {
	if p.offset+2 > len(p.content) {
		return "", fmt.Errorf("not enough data to read string length")
	}
	length := binary.BigEndian.Uint16(p.content[p.offset : p.offset+2])
	p.offset += 2
	if p.offset+int(length) > len(p.content) {
		return "", fmt.Errorf("not enough data to read string")
	}
	value := string(p.content[p.offset : p.offset+int(length)])
	p.offset += int(length)
	return value, nil
}

// FromBytes converts a byte slice into a RawPacket.
func FromBytes(data []byte) (*RawPacket, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("not enough data to read packet")
	}

	packetLength := binary.BigEndian.Uint32(data[:4])
	if len(data) < int(packetLength) {
		return nil, fmt.Errorf("packet size mismatch")
	}

	header := binary.BigEndian.Uint16(data[4:6])

	// Ensure that content is only read if there is any.
	var content []byte
	if packetLength > 6 {
		content = data[6:int(packetLength)]
	}

	return &RawPacket{
		header:  header,
		content: content,
		offset:  0,
	}, nil
}

// NewPacket creates from scratch a packet.
func NewPacket(header uint16) *RawPacket {
	return &RawPacket{
		header: header,
		offset: 0,
	}
}
