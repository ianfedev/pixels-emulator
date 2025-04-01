package user

import (
	"pixels-emulator/core/store"
)

// Store defines an ephemeral storage of online users
type Store interface {
	// Records provide the safe storage of players
	Records() store.AsyncStore[*Player]
}

type MemoryStore struct {
	store.AsyncStore[*Player]
}

// Records provide the safe storage of players
func (m *MemoryStore) Records() store.AsyncStore[*Player] {
	return m.AsyncStore
}

func NewUserStore() Store {
	return &MemoryStore{
		AsyncStore: store.NewMemoryStore[*Player](),
	}
}
