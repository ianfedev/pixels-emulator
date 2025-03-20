package room

import "pixels-emulator/core/store"

type Store struct {
	store.AsyncStore[*Room]
}

func NewStore() *Store {
	return &Store{
		AsyncStore: store.NewMemoryStore[*Room](),
	}
}
