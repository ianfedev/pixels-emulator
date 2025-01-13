package event

import "testing"

func TestNewEvent(t *testing.T) {
	userID := 42
	owner := uint16(1)
	metadata := map[string]string{"key": "value"}

	// Crear el evento usando el constructor.
	event := NewEvent(userID, owner, metadata)

	// Verificar que los valores se inicializaron correctamente.
	if event.UserID() != userID {
		t.Errorf("expected userID %d, got %d", userID, event.UserID())
	}

	if event.Owner() != owner {
		t.Errorf("expected owner %d, got %d", owner, event.Owner())
	}

	if len(event.Metadata()) != 1 || event.Metadata()["key"] != "value" {
		t.Errorf("metadata not initialized correctly, got %v", event.Metadata())
	}

	// Verificar que el evento sea cancelable.
	if event.IsCancelled() {
		t.Error("expected event to not be cancelled initially")
	}
}

func TestCancellableEvent(t *testing.T) {
	userID := 42
	owner := uint16(1)
	metadata := map[string]string{"key": "value"}

	// Crear el evento.
	event := NewEvent(userID, owner, metadata)

	// Cancelar el evento.
	event.Cancel()

	// Verificar que el evento está cancelado.
	if !event.IsCancelled() {
		t.Error("expected event to be cancelled after calling Cancel")
	}
}
