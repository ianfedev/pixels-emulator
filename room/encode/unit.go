package encode

import (
	"fmt"
	"pixels-emulator/core/protocol"
	"pixels-emulator/room/unit"
	"strings"
)

// UnitMessage acts as an encodable wrapper for the
// movement position of RoomUnits.
type UnitMessage struct {
	protocol.Encodable
	Id                  int32
	X, Y, Z, Head, Body int32
	Status              map[unit.Status]string
}

// Encode adds the RoomChatSettings data to a packet
func (e *UnitMessage) Encode(pck *protocol.RawPacket) {

	pck.AddInt(e.Id)
	pck.AddInt(e.X)
	pck.AddInt(e.Y)
	pck.AddInt(e.Z)
	pck.AddInt(e.Head)
	pck.AddInt(e.Body)

	var builder strings.Builder
	builder.WriteString("/")

	for k, v := range e.Status {
		builder.WriteString(fmt.Sprintf("%s %s/", k, v))
	}

	pck.AddString(builder.String())

}

func (e *UnitMessage) Decode(pck *protocol.RawPacket) error {

	var err error

	e.Id, err = pck.ReadInt()
	e.X, err = pck.ReadInt()
	e.Y, err = pck.ReadInt()
	e.Z, err = pck.ReadInt()
	e.Head, err = pck.ReadInt()
	e.Body, err = pck.ReadInt()

	statusStr, err := pck.ReadString()

	e.Status = make(map[unit.Status]string)

	parts := strings.Split(strings.Trim(statusStr, "/"), "/")
	for _, part := range parts {
		if part == "" {
			continue
		}
		kv := strings.SplitN(part, " ", 2)
		if len(kv) != 2 {
			continue
		}
		e.Status[unit.Status(kv[0])] = kv[1]
	}

	return err
}
