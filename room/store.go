package room

import (
	"pixels-emulator/core/store"
	"pixels-emulator/core/util"
)

// Store defines an ephemeral storage of online users
type Store interface {
	// Records provide the safe storage of players
	Records() store.AsyncStore[*Room]

	// Limits provide the attempt limiter for the store.
	Limits() *util.AttemptLimiter
}

type MemoryStore struct {
	PassLimit *util.AttemptLimiter
	store.AsyncStore[*Room]
}

func (m *MemoryStore) Records() store.AsyncStore[*Room] {
	return m.AsyncStore
}

func (m *MemoryStore) Limits() *util.AttemptLimiter {
	return m.PassLimit
}

func NewRoomStore() Store {
	return &MemoryStore{
		PassLimit:  util.NewAttemptLimiter(),
		AsyncStore: store.NewMemoryStore[*Room](),
	}
}
