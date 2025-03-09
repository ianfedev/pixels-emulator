package event_test

import (
	event2 "pixels-emulator/navigator/event"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNavigatorQueryEvent(t *testing.T) {
	realm := "test_realm"
	query := map[string]string{"key": "value"}
	owner := uint16(123)
	metadata := map[string]string{"meta": "data"}
	e := event2.NewNavigatorQueryEvent(realm, query, owner, metadata)

	assert.NotNil(t, e)
	assert.Equal(t, realm, e.Realm())
	assert.Equal(t, query, e.Query())
	assert.Equal(t, owner, e.Owner())
	assert.Equal(t, metadata, e.Metadata())
}

func TestNavigatorQueryEvent_Cancel(t *testing.T) {
	realm := "test_realm"
	query := map[string]string{"key": "value"}
	owner := uint16(123)
	metadata := map[string]string{"meta": "data"}

	e := event2.NewNavigatorQueryEvent(realm, query, owner, metadata)
	e.Cancel()

	assert.True(t, e.IsCancelled())
}
