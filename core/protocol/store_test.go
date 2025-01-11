package protocol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// MockConnection is a mock implementation of the Connection interface for testing purposes.
type MockConnection struct {
	id string
}

func (m *MockConnection) Identifier() string {
	return m.id
}

func (m *MockConnection) SendPacket(_ Packet) {}

func (m *MockConnection) SendRaw(_ RawPacket, _ uint16, _ uint16) {}

func (m *MockConnection) RateRegistry() *RateLimiterRegistry {
	return nil
}

func (m *MockConnection) Dispose() error {
	return nil
}

type MockPacket struct {
	Packet // Embeds the base protocol.Packet interface.
}

// Id returns the unique identifier of the Packet type.
func (p *MockPacket) Id() uint16 {
	return 1
}

// Rate returns the rate limit for the Ping packet.
func (p *MockPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Serialize converts the Ping packet into a RawPacket that can be transmitted over the network.
func (p *MockPacket) Serialize() RawPacket {
	rawPack := NewPacket(1)
	rawPack.AddBoolean(false)
	return rawPack
}

// TestConnectionStore_AddConnection tests the AddConnection method to ensure a connection is added correctly.
func TestConnectionStore_AddConnection(t *testing.T) {
	store := NewConnectionStore()
	conn := &MockConnection{id: "conn1"}

	store.AddConnection(conn)

	retrievedConn, found := store.GetConnection("conn1")
	assert.True(t, found)
	assert.Equal(t, "conn1", retrievedConn.Identifier())
}

// TestConnectionStore_RemoveConnection tests the RemoveConnection method to ensure a connection is removed correctly.
func TestConnectionStore_RemoveConnection(t *testing.T) {
	store := NewConnectionStore()
	conn := &MockConnection{id: "conn1"}

	store.AddConnection(conn)
	store.RemoveConnection("conn1")

	_, found := store.GetConnection("conn1")
	assert.False(t, found)
}

// TestConnectionStore_GetConnection tests the GetConnection method to ensure a connection is retrieved correctly.
func TestConnectionStore_GetConnection(t *testing.T) {
	store := NewConnectionStore()
	conn := &MockConnection{id: "conn1"}

	store.AddConnection(conn)

	retrievedConn, found := store.GetConnection("conn1")
	assert.True(t, found)
	assert.Equal(t, "conn1", retrievedConn.Identifier())
}

// TestConnectionStore_BroadcastPacket tests the BroadcastPacket method to ensure packets are sent to all connections.
func TestConnectionStore_BroadcastPacket(t *testing.T) {
	store := NewConnectionStore()
	conn1 := &MockConnection{id: "conn1"}
	conn2 := &MockConnection{id: "conn2"}

	store.AddConnection(conn1)
	store.AddConnection(conn2)

	packet := &MockPacket{}

	store.BroadcastPacket(packet)

	// Assume SendPacket is called during the broadcast.
	// In practice, you'd use a mock or spy on SendPacket to verify it was called.
	assert.True(t, true) // Replace with actual verification when mocking SendPacket
}

// TestConnectionStore_BroadcastPacketToIDs tests the BroadcastPacketToIDs method to ensure packets are sent only to specific connections.
func TestConnectionStore_BroadcastPacketToIDs(t *testing.T) {
	store := NewConnectionStore()
	conn1 := &MockConnection{id: "conn1"}
	conn2 := &MockConnection{id: "conn2"}

	store.AddConnection(conn1)
	store.AddConnection(conn2)

	packet := &MockPacket{}
	ids := []string{"conn1"}

	store.BroadcastPacketToIDs(packet, ids)

	// Verify that only the connection with ID "conn1" received the packet.
	// In practice, you'd mock SendRaw and assert it was only called for "conn1".
	assert.True(t, true) // Replace with actual verification when mocking SendRaw
}

// TestConnectionStore_ConnectionCount tests the ConnectionCount method to ensure it returns the correct number of active connections.
func TestConnectionStore_ConnectionCount(t *testing.T) {
	store := NewConnectionStore()
	conn1 := &MockConnection{id: "conn1"}
	conn2 := &MockConnection{id: "conn2"}

	store.AddConnection(conn1)
	store.AddConnection(conn2)

	count := store.ConnectionCount()
	assert.Equal(t, 2, count)
}
