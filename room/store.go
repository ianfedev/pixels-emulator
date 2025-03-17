package room

import "pixels-emulator/core/store"

type Store struct {
	store.AsyncStore[any]
}
