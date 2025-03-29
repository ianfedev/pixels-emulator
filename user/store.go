package user

import "pixels-emulator/core/store"

type Store struct {
	store.AsyncStore[*Player]
}

func NewStore() *Store {
	return &Store{
		AsyncStore: store.NewMemoryStore[*Player](),
	}
}
