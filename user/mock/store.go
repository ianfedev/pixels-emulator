package mock

import (
	"github.com/stretchr/testify/mock"
	"pixels-emulator/core/store"
	"pixels-emulator/user"
)

// Store is a mock implementation of the user Store interface.
type Store struct {
	mock.Mock
}

// Records simulates the retrieving of an async store.
func (m *Store) Records() store.AsyncStore[*user.Player] {
	args := m.Called()
	return args.Get(0).(store.AsyncStore[*user.Player])
}
