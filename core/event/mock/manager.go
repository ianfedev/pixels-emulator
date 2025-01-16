package event_test

import (
	"github.com/stretchr/testify/mock"
	"pixels-emulator/core/event"
)

// MockEventManager is a mock implementation of the Manager interface.
type MockEventManager struct {
	mock.Mock
}

// Fire simulates the Fire method of the Manager interface.
func (m *MockEventManager) Fire(eventName string, event event.Event) error {
	args := m.Called(eventName, event)
	return args.Error(0)
}

// AddListener simulates the AddListener method of the Manager interface.
func (m *MockEventManager) AddListener(eventName string, listener func(event event.Event), priority int) {
	m.Called(eventName, listener, priority)
}

// Close simulates the Close method of the Manager interface.
func (m *MockEventManager) Close() error {
	args := m.Called()
	return args.Error(0)
}
