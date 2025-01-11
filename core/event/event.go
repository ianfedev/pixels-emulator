package event

// Event defines the interface for an event with metadata and an owner.
type Event interface {
	// Owner returns the owner identifier of the event.
	Owner() uint16

	// Key returns the value associated with the given key in metadata.
	Key(key string) string

	// Metadata returns the metadata map.
	Metadata() map[string]string
}

// BaseEvent is a basic implementation of the Event interface, storing metadata and an owner.
type BaseEvent struct {
	metadata map[string]string
	owner    uint16
}

// Owner returns the owner identifier of the BaseEvent.
func (e *BaseEvent) Owner() uint16 {
	return e.owner
}

// Key returns the value associated with the given key in metadata for BaseEvent.
func (e *BaseEvent) Key(key string) string {
	return e.metadata[key]
}

// Metadata returns the metadata associated with the event.
func (e *BaseEvent) Metadata() map[string]string {
	return e.metadata
}

// New creates a new Event with provided metadata and owner.
func New(owner uint16, metadata map[string]string) Event {
	return &BaseEvent{
		owner:    owner,
		metadata: metadata,
	}
}

// CancellableEvent defines an event that can be cancelled, extending BaseEvent.
type CancellableEvent struct {
	// Embedding BaseEvent to reuse its functionality.
	*BaseEvent
	// Cancelled indicates whether the event has been cancelled.
	Cancelled bool
}

// NewCancellable creates a new CancellableEvent with the provided metadata and owner.
func NewCancellable(owner uint16, metadata map[string]string) Event {
	return &CancellableEvent{
		BaseEvent: &BaseEvent{
			owner:    owner,
			metadata: metadata,
		},
		Cancelled: false,
	}
}

// Cancel marks the event as cancelled.
func (e *CancellableEvent) Cancel() {
	e.Cancelled = true
}

// IsCancelled returns a boolean indicating if the event has been cancelled.
func (e *CancellableEvent) IsCancelled() bool {
	return e.Cancelled
}
