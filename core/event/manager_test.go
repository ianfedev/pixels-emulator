package event

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExampleEvent struct {
	BaseEvent
	Message string
}

func NewExampleEvent(owner uint16, metadata map[string]string, message string) *ExampleEvent {
	return &ExampleEvent{
		BaseEvent: BaseEvent{
			owner:    owner,
			metadata: metadata,
		},
		Message: message,
	}
}

func TestManager_Fire(t *testing.T) {
	manager := NewManager()
	defer manager.Close()

	event := NewExampleEvent(42, map[string]string{
		"type": "test_event",
	}, "Test Event Message!")

	var wg sync.WaitGroup
	wg.Add(1)

	manager.AddListener("test_event", func(e Event) {
		defer wg.Done()
		assert.Equal(t, uint16(42), e.Owner())
		assert.Equal(t, "Test Event Message!", e.(*ExampleEvent).Message)
	}, 0)

	manager.Fire("test_event", event)

	wg.Wait() // esperar a que se dispare el listener
}

func TestManager_AddListener(t *testing.T) {
	manager := NewManager()
	defer manager.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	manager.AddListener("sample_event", func(e Event) {
		defer wg.Done()
	}, 0)

	manager.Fire("sample_event", &ExampleEvent{})

	wg.Wait()
}

func TestManager_Close(t *testing.T) {
	manager := NewManager()

	err := manager.Close()
	assert.NoError(t, err)
}
