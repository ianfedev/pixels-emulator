package message

import (
	"pixels-emulator/core/protocol"
	"pixels-emulator/navigator/encode"
)

// NavigatorSearchResultCode is the unique identifier for the packet.
const NavigatorSearchResultCode = 2690

// NavigatorSearchResultPacket represents a response containing multiple search result compounds.
type NavigatorSearchResultPacket struct {
	protocol.Packet

	// SearchCode is the identifier used by the query and filter.
	SearchCode string

	// SearchQuery defines the input search term or query filter.
	SearchQuery string

	// Results contains the list of SearchResultCompounds to be sent.
	Results []*encode.SearchResultCompound
}

// Id returns the unique identifier of the Packet type.
func (p *NavigatorSearchResultPacket) Id() uint16 {
	return NavigatorSearchResultCode
}

// Rate returns the rate limit for the packet.
func (p *NavigatorSearchResultPacket) Rate() (uint16, uint16) {
	return 10, 2
}

// Serialize writes the packet data into a RawPacket.
func (p *NavigatorSearchResultPacket) Serialize(pck *protocol.RawPacket) {
	pck.AddString(p.SearchCode)
	pck.AddString(p.SearchQuery)

	pck.AddInt(int32(len(p.Results))) // Number of result blocks

	for _, result := range p.Results {
		result.Encode(pck) // Encode each SearchResultCompound
	}
}

// ComposeNavigatorSearchResult creates a new NavigatorSearchResultPacket.
func ComposeNavigatorSearchResult(searchCode string, searchQuery string, results []*encode.SearchResultCompound) *NavigatorSearchResultPacket {
	return &NavigatorSearchResultPacket{
		SearchCode:  searchCode,
		SearchQuery: searchQuery,
		Results:     results,
	}
}
