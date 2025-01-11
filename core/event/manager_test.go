package event

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ExampleEvent struct {
	BaseEvent
	Message string
}

// NewExampleEvent creates a new ExampleEvent with provided owner, metadata, and message.
func NewExampleEvent(owner uint16, metadata map[string]string, message string) *ExampleEvent {
	return &ExampleEvent{
		BaseEvent: BaseEvent{
			owner:    owner,
			metadata: metadata,
		},
		Message: message,
	}
}

// TestManager_Fire tests the Fire method of the Manager, ensuring event firing works as expected.
func TestManager_Fire(t *testing.T) {
	manager := NewManager()

	event := NewExampleEvent(42, map[string]string{
		"type": "test_event",
	}, "Test Event Message!")

	listenerCalled := false
	manager.AddListener("test_event", func(e Event) {
		listenerCalled = true
		assert.Equal(t, uint16(42), e.Owner())
		assert.Equal(t, "Test Event Message!", e.(*ExampleEvent).Message)
	}, 0)

	err := manager.Fire("test_event", event)
	if err != nil {
		t.Errorf("Error firing event: %v", err)
	}

	assert.True(t, listenerCalled)
}

// TestManager_AddListener tests the AddListener method, ensuring listeners are added correctly.
func TestManager_AddListener(t *testing.T) {
	manager := NewManager()

	listenerCalled := false
	manager.AddListener("sample_event", func(e Event) {
		listenerCalled = true
	}, 0)

	err := manager.Fire("sample_event", &ExampleEvent{})

	assert.True(t, listenerCalled, "Listener should be called")
	assert.Nil(t, err, "Error should not be present on firing")
}

// TestManager_Close tests the Close method, ensuring the Manager closes without errors.
func TestManager_Close(t *testing.T) {
	manager := NewManager()

	err := manager.Close()
	assert.NoError(t, err)
}
