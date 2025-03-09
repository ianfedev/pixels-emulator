package protocol

// Encodable represents a data structure that can be written into and read from a RawPacket.
type Encodable interface {
	// Encode writes the Encodable structure into the given RawPacket.
	Encode(pck *RawPacket)

	// Decode reads the Encodable structure from the given RawPacket.
	Decode(pck *RawPacket) error
}
