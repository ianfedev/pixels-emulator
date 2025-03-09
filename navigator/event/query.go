package event

import (
	"pixels-emulator/core/event"
)

const NavigatorQueryEventName = "navigator.query"

// NavigatorQueryEvent represents a navigation search event.
type NavigatorQueryEvent struct {
	*event.CancellableEvent                   // Extends functionality for cancellation.
	realm                   string            // View or category being queried.
	query                   map[string]string // Query parameters for filtering results.
}

// NewNavigatorQueryEvent creates a new NavigatorQueryEvent instance.
func NewNavigatorQueryEvent(realm string, query map[string]string, owner uint16, metadata map[string]string) *NavigatorQueryEvent {
	ce := event.NewCancellable(owner, metadata)
	return &NavigatorQueryEvent{
		CancellableEvent: ce.(*event.CancellableEvent),
		realm:            realm,
		query:            query,
	}
}

// Realm returns the navigation view or category.
func (e *NavigatorQueryEvent) Realm() string {
	return e.realm
}

// Query returns the query parameters for filtering.
func (e *NavigatorQueryEvent) Query() map[string]string {
	return e.query
}
