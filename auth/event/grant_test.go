package event

import "testing"

// TestNewEvent tests the initialization of a new event.
func TestNewEvent(t *testing.T) {
	userID := 42
	owner := uint16(1)
	metadata := map[string]string{"key": "value"}

	event := NewEvent(userID, owner, metadata)

	if event.UserID() != userID {
		t.Errorf("expected userID %d, got %d", userID, event.UserID())
	}

	if event.Owner() != owner {
		t.Errorf("expected owner %d, got %d", owner, event.Owner())
	}

	if len(event.Metadata()) != 1 || event.Metadata()["key"] != "value" {
		t.Errorf("metadata not initialized correctly, got %v", event.Metadata())
	}

	if event.IsCancelled() {
		t.Error("expected event to not be cancelled initially")
	}
}

// TestCancellableEvent tests cancelling an event.
func TestCancellableEvent(t *testing.T) {
	userID := 42
	owner := uint16(1)
	metadata := map[string]string{"key": "value"}

	event := NewEvent(userID, owner, metadata)
	event.Cancel()

	if !event.IsCancelled() {
		t.Error("expected event to be cancelled after calling Cancel")
	}
}
