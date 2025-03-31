package room

import (
	"pixels-emulator/core/store"
	"pixels-emulator/core/util"
)

type Store struct {
	PassLimit *util.AttemptLimiter
	store.AsyncStore[*Room]
}

func NewRoomStore() *Store {
	return &Store{
		PassLimit:  util.NewAttemptLimiter(),
		AsyncStore: store.NewMemoryStore[*Room](),
	}
}
