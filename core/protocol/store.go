package protocol

import (
	"sync"
)

// ConnectionStore centralizes all active connections and provides thread-safe access.
type ConnectionStore struct {
	connections sync.Map // Stores connections with their unique identifier as the key.
}

// NewConnectionStore creates a new connection manager instance.
func NewConnectionStore() *ConnectionStore {
	return &ConnectionStore{}
}

// AddConnection adds a new connection to the manager.
func (m *ConnectionStore) AddConnection(conn Connection) {
	m.connections.Store(conn.Identifier(), conn)
}

// RemoveConnection removes a connection from the manager by its identifier.
func (m *ConnectionStore) RemoveConnection(identifier string) {
	m.connections.Delete(identifier)
}

// GetConnection retrieves a connection by its identifier.
// Returns the connection and a boolean indicating if it was found.
func (m *ConnectionStore) GetConnection(identifier string) (Connection, bool) {
	conn, ok := m.connections.Load(identifier)
	if !ok {
		return nil, false
	}
	return conn.(Connection), true
}

// buildIDSet creates a map of identifiers for efficient lookup.
func buildIDSet(ids []string) map[string]struct{} {
	idSet := make(map[string]struct{}, len(ids))
	for _, id := range ids {
		idSet[id] = struct{}{}
	}
	return idSet
}

// broadcast sends a serialized packet to connections based on the provided filter function.
// The filter determines which connections should receive the packet.
func (m *ConnectionStore) broadcast(packet Packet, shouldSend func(conn Connection) bool) {
	raw := packet.Serialize()
	period, rate := packet.Rate()

	m.connections.Range(func(_, value interface{}) bool {
		conn := value.(Connection)
		if shouldSend(conn) {
			conn.SendRaw(raw, period, rate)
		}
		return true
	})
}

// BroadcastPacket sends a packet to all active connections.
func (m *ConnectionStore) BroadcastPacket(packet Packet) {
	m.broadcast(packet, func(conn Connection) bool { return true })
}

// BroadcastPacketToIDs sends a packet to a subset of connections identified by their IDs.
func (m *ConnectionStore) BroadcastPacketToIDs(packet Packet, ids []string) {
	idSet := buildIDSet(ids)
	m.broadcast(packet, func(conn Connection) bool {
		_, found := idSet[conn.Identifier()]
		return found
	})
}

// ConnectionCount returns the total number of active connections.
func (m *ConnectionStore) ConnectionCount() int {
	count := 0
	m.connections.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	return count
}
