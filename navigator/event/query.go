package event

import (
	"pixels-emulator/core/event"
	"strings"
)

const NavigatorQueryEventName = "navigator.query"

// NavigatorQueryEvent represents a navigation search event.
type NavigatorQueryEvent struct {
	*event.CancellableEvent                   // Extends functionality for cancellation.
	realm                   string            // View or category being queried.
	query                   map[string]string // Query parameters for filtering results.
	rawQuery                string            // rawQuery provides the query as it came from client.
}

// NewNavigatorQueryEvent creates a new NavigatorQueryEvent instance.
func NewNavigatorQueryEvent(realm string, rawQuery string, owner uint16, metadata map[string]string) *NavigatorQueryEvent {
	parsedQuery := parseQuery(rawQuery)
	ce := event.NewCancellable(owner, metadata)
	return &NavigatorQueryEvent{
		CancellableEvent: ce.(*event.CancellableEvent),
		realm:            realm,
		query:            parsedQuery,
		rawQuery:         rawQuery,
	}
}

// Realm returns the navigation view or category.
func (e *NavigatorQueryEvent) Realm() string {
	return e.realm
}

// RawQuery returns the original raw query string.
func (e *NavigatorQueryEvent) RawQuery() string {
	return e.rawQuery
}

// Query returns the query parameters for filtering.
func (e *NavigatorQueryEvent) Query() map[string]string {
	return e.query
}

// parseQuery processes the raw query into a key-value map.
func parseQuery(raw string) map[string]string {
	queryParams := make(map[string]string)

	if raw == "" {
		return queryParams
	}

	parts := strings.Split(raw, ":")
	if len(parts) > 1 {
		queryParams[parts[0]] = parts[1]
	} else {
		queryParams["query"] = raw
	}

	return queryParams
}
