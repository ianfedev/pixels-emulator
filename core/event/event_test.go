package event

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestBaseEvent_Owner tests the Owner method for the BaseEvent, ensuring it returns the correct owner.
func TestBaseEvent_Owner(t *testing.T) {
	event := New(42, map[string]string{"type": "test_event"})

	assert.Equal(t, uint16(42), event.Owner())
}

// TestBaseEvent_Key tests the Key method for the BaseEvent, ensuring it retrieves metadata correctly.
func TestBaseEvent_Key(t *testing.T) {
	event := New(42, map[string]string{"type": "test_event", "message": "Hello"})

	assert.Equal(t, "Hello", event.Key("message"))
	assert.Empty(t, event.Key("nonexistent_key"))
}

// TestBaseEvent_Metadata tests the Metadata method for the BaseEvent, ensuring it returns the correct metadata map.
func TestBaseEvent_Metadata(t *testing.T) {
	metadata := map[string]string{"type": "test_event", "message": "Hello"}
	event := New(42, metadata)

	assert.Equal(t, metadata, event.Metadata())
}

// TestCancellableEvent_Cancel tests the Cancel method of CancellableEvent, ensuring it sets the Cancelled flag.
func TestCancellableEvent_Cancel(t *testing.T) {
	event := NewCancellable(42, map[string]string{"type": "cancellable_event"})

	assert.False(t, event.(*CancellableEvent).IsCancelled())

	event.(*CancellableEvent).Cancel()

	assert.True(t, event.(*CancellableEvent).IsCancelled())
}

// TestCancellableEvent_IsCancelled tests the IsCancelled method of CancellableEvent, ensuring it returns the correct status.
func TestCancellableEvent_IsCancelled(t *testing.T) {
	event := NewCancellable(42, map[string]string{"type": "cancellable_event"})

	assert.False(t, event.(*CancellableEvent).IsCancelled())

	event.(*CancellableEvent).Cancel()

	assert.True(t, event.(*CancellableEvent).IsCancelled())
}

// TestNewCancellable tests the NewCancellable function, ensuring it creates a valid cancellable event.
func TestNewCancellable(t *testing.T) {
	event := NewCancellable(42, map[string]string{"type": "cancellable_event"})

	assert.NotNil(t, event)
	assert.Equal(t, uint16(42), event.Owner())
	assert.Equal(t, "cancellable_event", event.Key("type"))
}

// TestNew tests the New function for creating BaseEvent, ensuring it initializes correctly.
func TestNew(t *testing.T) {
	metadata := map[string]string{"type": "test_event", "message": "Test"}
	event := New(42, metadata)

	assert.NotNil(t, event)
	assert.Equal(t, uint16(42), event.Owner())
	assert.Equal(t, "Test", event.Key("message"))
}
