package protocol

import (
	"sync"
)

// ConnectionManager defines a manager which can thread-safely manage active connections.
type ConnectionManager interface {
	AddConnection(conn *Connection)                      // AddConnection adds a new connection to the manager.
	RemoveConnection(identifier string)                  // RemoveConnection removes a connection from the manager by its identifier.
	GetConnection(identifier string) (*Connection, bool) // GetConnection retrieves a connection by its identifier.
	BroadcastPacket(packet Packet)                       // BroadcastPacket sends a packet to all active connections.
	BroadcastPacketToIDs(packet Packet, ids []string)    // BroadcastPacketToIDs sends a packet to a subset of connections identified by their IDs.
	ConnectionCount() int                                // ConnectionCount returns the total number of active connections.
	CloseActive()                                        // CloseActive closes all active connections and clears the connection store.
}

// ConnectionStore centralizes all active connections and provides thread-safe access.
type ConnectionStore struct {
	connections []*Connection // Stores all active connections.
	mutex       sync.Mutex    // Protects access to the connections slice.
}

// NewConnectionStore creates a new connection manager instance.
func NewConnectionStore() ConnectionManager {
	return &ConnectionStore{}
}

// AddConnection adds a new connection to the manager.
func (m *ConnectionStore) AddConnection(conn *Connection) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.connections = append(m.connections, conn)
}

// RemoveConnection removes a connection from the manager by its identifier.
func (m *ConnectionStore) RemoveConnection(identifier string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for i, conn := range m.connections {
		if (*conn).Identifier() == identifier {
			// Remove the connection by slicing it out.
			m.connections = append(m.connections[:i], m.connections[i+1:]...)
			break
		}
	}
}

// GetConnection retrieves a connection by its identifier.
// Returns the connection and a boolean indicating if it was found.
func (m *ConnectionStore) GetConnection(identifier string) (*Connection, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for _, conn := range m.connections {
		if (*conn).Identifier() == identifier {
			return conn, true
		}
	}
	return nil, false
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
func (m *ConnectionStore) broadcast(packet Packet, shouldSend func(conn *Connection) bool) {
	raw := packet.Serialize()
	period, rate := packet.Rate()

	m.mutex.Lock()
	defer m.mutex.Unlock()
	for _, conn := range m.connections {
		if shouldSend(conn) {
			(*conn).SendRaw(raw, period, rate)
		}
	}
}

// BroadcastPacket sends a packet to all active connections.
func (m *ConnectionStore) BroadcastPacket(packet Packet) {
	m.broadcast(packet, func(conn *Connection) bool { return true })
}

// BroadcastPacketToIDs sends a packet to a subset of connections identified by their IDs.
func (m *ConnectionStore) BroadcastPacketToIDs(packet Packet, ids []string) {
	idSet := buildIDSet(ids)
	m.broadcast(packet, func(conn *Connection) bool {
		_, found := idSet[(*conn).Identifier()]
		return found
	})
}

// ConnectionCount returns the total number of active connections.
func (m *ConnectionStore) ConnectionCount() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return len(m.connections)
}

// CloseActive closes all active connections and clears the connection store.
func (m *ConnectionStore) CloseActive() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for _, conn := range m.connections {
		_ = (*conn).Dispose() // Assuming Close is a method on Connection that handles cleanup.
	}
	m.connections = nil // Clear all connections after closing them.
}
