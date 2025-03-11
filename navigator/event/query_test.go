package event_test

import (
	mockproto "pixels-emulator/core/protocol/mock"
	"pixels-emulator/navigator/event"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewNavigatorQueryEvent creates an event with a parsed query.
func TestNewNavigatorQueryEvent(t *testing.T) {
	realm := "test_realm"
	rawQuery := "owner:John"
	owner := uint16(123)
	metadata := map[string]string{"meta": "data"}
	con := &mockproto.MockConnection{}
	con.On("Identifier").Return(1)

	e := event.NewNavigatorQueryEvent(realm, rawQuery, con, owner, metadata)

	assert.NotNil(t, e)
	assert.Equal(t, realm, e.Realm())
	assert.Equal(t, rawQuery, e.RawQuery())
	assert.Equal(t, map[string]string{"owner": "John"}, e.Query())
	assert.Equal(t, owner, e.Owner())
	assert.Equal(t, metadata, e.Metadata())
}

// NewNavigatorQueryEvent assigns "query" as key when no key is provided.
func TestNewNavigatorQueryEvent_SimpleQuery(t *testing.T) {
	realm := "test_realm"
	rawQuery := "hotel_view"
	owner := uint16(123)
	metadata := map[string]string{"meta": "data"}
	con := &mockproto.MockConnection{}
	con.On("Identifier").Return(1)

	e := event.NewNavigatorQueryEvent(realm, rawQuery, con, owner, metadata)

	assert.NotNil(t, e)
	assert.Equal(t, rawQuery, e.RawQuery())
	assert.Equal(t, map[string]string{"query": "hotel_view"}, e.Query())
}

// Cancel marks the event as cancelled.
func TestNavigatorQueryEvent_Cancel(t *testing.T) {
	realm := "test_realm"
	rawQuery := "owner:John"
	owner := uint16(123)
	metadata := map[string]string{"meta": "data"}
	con := &mockproto.MockConnection{}
	con.On("Identifier").Return(1)

	e := event.NewNavigatorQueryEvent(realm, rawQuery, con, owner, metadata)
	e.Cancel()

	assert.True(t, e.IsCancelled())
}
