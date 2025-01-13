package protocol_test

import (
	"github.com/stretchr/testify/mock"
	"pixels-emulator/core/protocol"
)

// MockConnectionManager is a mock of the ConnectionManager interface.
type MockConnectionManager struct {
	mock.Mock
}

// AddConnection adds a new connection to the manager.
func (m *MockConnectionManager) AddConnection(conn protocol.Connection) {
	m.Called(conn)
}

// RemoveConnection removes a connection from the manager by its identifier.
func (m *MockConnectionManager) RemoveConnection(identifier string) {
	m.Called(identifier)
}

// GetConnection retrieves a connection by its identifier.
func (m *MockConnectionManager) GetConnection(identifier string) (protocol.Connection, bool) {
	args := m.Called(identifier)
	return args.Get(0).(protocol.Connection), args.Bool(1)
}

// BroadcastPacket sends a packet to all active connections.
func (m *MockConnectionManager) BroadcastPacket(packet protocol.Packet) {
	m.Called(packet)
}

// BroadcastPacketToIDs sends a packet to a subset of connections identified by their IDs.
func (m *MockConnectionManager) BroadcastPacketToIDs(packet protocol.Packet, ids []string) {
	m.Called(packet, ids)
}

// ConnectionCount returns the total number of active connections.
func (m *MockConnectionManager) ConnectionCount() int {
	args := m.Called()
	return args.Int(0)
}

// CloseActive closes all active connections and clears the connection store.
func (m *MockConnectionManager) CloseActive() {
	m.Called()
}
