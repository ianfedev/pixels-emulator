package encode

import (
	"fmt"
	"pixels-emulator/core/protocol"
	"pixels-emulator/room/unit"
	"strconv"
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
	pck.AddString(strconv.Itoa(int(e.Z)))
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
	if err != nil {
		return err
	}

	e.X, err = pck.ReadInt()
	if err != nil {
		return err
	}

	e.Y, err = pck.ReadInt()
	if err != nil {
		return err
	}

	z, err := pck.ReadString()
	if err != nil {
		return err
	}

	pz, err := strconv.ParseInt(z, 10, 32)
	if err != nil {
		return err
	}
	e.Z = int32(pz)

	e.Head, err = pck.ReadInt()
	if err != nil {
		return err
	}

	e.Body, err = pck.ReadInt()
	if err != nil {
		return err
	}

	statusStr, err := pck.ReadString()
	if err != nil {
		return err
	}

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
