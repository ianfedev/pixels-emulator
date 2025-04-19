package unit

import (
	"errors"
	"pixels-emulator/core/protocol"
	"pixels-emulator/room/encode"
)

// DetailCode is the unique identifier for the packet
const DetailCode = 374

// DetailPacket defines a packet to broadcast all the details
// of specific units.
type DetailPacket struct {
	Units        []*encode.UnitDetail
	PlayerDetail []*encode.PlayerDetail
	PetDetail    []*encode.PetDetail
	BotDetail    []*encode.RentableBotDetail
	protocol.Packet
}

// Id returns the unique identifier of the Packet type.
func (p *DetailPacket) Id() uint16 {
	return DetailCode
}

// Rate returns the rate limit for the packet.
func (p *DetailPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *DetailPacket) Deadline() uint {
	return 0
}

// Serialize transforms packet in byte.
func (p *DetailPacket) Serialize(unitType encode.UnitType) (*protocol.RawPacket, error) {

	pck := protocol.NewPacket(DetailCode)
	pck.AddInt(int32(len(p.Units)))

	switch unitType {
	case encode.User:
		if len(p.Units) != len(p.PlayerDetail) {
			return nil, errors.New("serialization type mismatch")
		}
		for i := 0; i < len(p.Units); i++ {
			p.Units[i].Encode(&pck)
			p.PlayerDetail[i].Encode(&pck)
		}
		break
	case encode.Pet:
		if len(p.Units) != len(p.PetDetail) {
			return nil, errors.New("serialization type mismatch")
		}
		// TODO: Pets
		break
	case encode.Bot:
	case encode.Rentable:
		if len(p.Units) != len(p.BotDetail) {
			return nil, errors.New("serialization type mismatch")
		}
		// TODO: Bots
		break
	default:
		return nil, errors.New("no serializable type provided")
	}

	return &pck, nil

}
