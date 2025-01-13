package protocol_test

import (
	"github.com/stretchr/testify/mock"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/util"
)

// MockConnection is a mock of the Connection interface.
type MockConnection struct {
	mock.Mock
	util.Disposable
}

// Identifier provides an unique identifier of this connection.
func (m *MockConnection) Identifier() string {
	args := m.Called()
	return args.String(0)
}

// GrantIdentifier provides a new identifier for connection.
func (m *MockConnection) GrantIdentifier(identifier string) {
	m.Called(identifier)
}

// SendPacket pings an outgoing packet.
func (m *MockConnection) SendPacket(packet protocol.Packet) {
	m.Called(packet)
}

// SendRaw pings a raw packet with custom restriction.
func (m *MockConnection) SendRaw(packet protocol.RawPacket, period uint16, rate uint16) {
	m.Called(packet, period, rate)
}

// RateRegistry limit rates outgoing packets.
func (m *MockConnection) RateRegistry() protocol.RateLimiter {
	args := m.Called()
	// Returning the actual object (make sure it doesn't cause an issue due to the mutex)
	return args.Get(0).(protocol.RateLimiter)
}

// Dispose releases resources or closes connections associated with the object.
func (m *MockConnection) Dispose() error {
	args := m.Called()
	return args.Error(0)
}
