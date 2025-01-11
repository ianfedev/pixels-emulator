package event

import (
	"fmt"
	gEvent "github.com/gookit/event"
)

// Manager manages event dispatching and listener registration using gookit/event.
type Manager struct {
	NativeRegistry *gEvent.Manager
}

// Fire triggers an event with the given name and event data.
func (em *Manager) Fire(eventName string, event Event) error {
	ev := gEvent.NewBasic(eventName, map[string]interface{}{
		"owner":    event.Owner(),
		"metadata": event.Metadata(),
		"origin":   event,
	})
	return em.NativeRegistry.FireEvent(ev)
}

// AddListener registers a listener for the specified event name.
func (em *Manager) AddListener(eventName string, listener func(event Event), priority int) {
	fmt.Println("added listener")
	em.NativeRegistry.On(eventName, gEvent.ListenerFunc(func(e gEvent.Event) error {
		if customEvent, ok := e.Get("origin").(Event); ok {
			listener(customEvent)
		}
		return nil
	}), priority)
}

// Close waits for all registered events to complete and closes the manager.
func (em *Manager) Close() error {
	return em.NativeRegistry.CloseWait()
}

// NewManager creates a new instance of the event manager.
func NewManager() *Manager {
	return &Manager{
		NativeRegistry: gEvent.NewManager("pixels"),
	}
}
