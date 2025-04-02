package mock

import (
	"github.com/stretchr/testify/mock"
	"pixels-emulator/core/model"
	"pixels-emulator/core/store"
	"pixels-emulator/core/util"
)

// MemoryStore is a mock implementation of the room Store interface.
type MemoryStore struct {
	mock.Mock
}

// Records simulates the retrieving of an async store.
func (m *MemoryStore) Records() store.AsyncStore[*model.Room] {
	args := m.Called()
	return args.Get(0).(store.AsyncStore[*model.Room])
}

func (m *MemoryStore) Limits() *util.AttemptLimiter {
	args := m.Called()
	return args.Get(0).(*util.AttemptLimiter)
}
