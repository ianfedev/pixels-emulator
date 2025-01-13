package event

import (
	gEvent "github.com/gookit/event"
)

// Manager defines the interface for managing event dispatching and listener registration.
type Manager interface {
	// Fire triggers an event with the given name and event data.
	Fire(eventName string, event Event) error

	// AddListener registers a listener for the specified event name.
	AddListener(eventName string, listener func(event Event), priority int)

	// Close waits for all registered events to complete and closes the manager.
	Close() error
}

// GookitManager is an implementation of Manager using gookit/event for event handling.
type GookitManager struct {
	// NativeRegistry is the gookit event manager used for dispatching and managing events.
	NativeRegistry *gEvent.Manager
}

// Fire triggers an event with the given name and event data.
func (gm *GookitManager) Fire(eventName string, event Event) error {
	ev := gEvent.NewBasic(eventName, map[string]interface{}{
		"owner":    event.Owner(),
		"metadata": event.Metadata(),
		"origin":   event,
	})
	return gm.NativeRegistry.FireEvent(ev)
}

// AddListener registers a listener for the specified event name.
func (gm *GookitManager) AddListener(eventName string, listener func(event Event), priority int) {
	gm.NativeRegistry.On(eventName, gEvent.ListenerFunc(func(e gEvent.Event) error {
		if customEvent, ok := e.Get("origin").(Event); ok {
			listener(customEvent)
		}
		return nil
	}), priority)
}

// Close waits for all registered events to complete and closes the manager.
func (gm *GookitManager) Close() error {
	return gm.NativeRegistry.CloseWait()
}

// NewManager creates a new instance of the GookitManager event manager.
func NewManager() Manager {
	return &GookitManager{
		NativeRegistry: gEvent.NewManager("pixels"),
	}
}
