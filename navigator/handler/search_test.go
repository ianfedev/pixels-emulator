package handler

import (
	"bytes"
	"context"
	"errors"
	message2 "pixels-emulator/auth/message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mockevent "pixels-emulator/core/event/mock"
	mockproto "pixels-emulator/core/protocol/mock"
	"pixels-emulator/core/util"
	navEvent "pixels-emulator/navigator/event"
	"pixels-emulator/navigator/message"
)

// setupMockConn creates a basic connection to be modified
func setupMockConn() *mockproto.MockConnection {
	con := &mockproto.MockConnection{}
	return con
}

// setupTestEnvironment sets up the test environment for NavigatorSearchHandler
func setupTestEnvironment(t *testing.T, query string, view string, val interface{}) (*NavigatorSearchHandler, *bytes.Buffer) {
	log, buf := util.CreateTestLogger()

	em := &mockevent.MockEventManager{}
	em.On("Fire", navEvent.NavigatorQueryEventName, mock.Anything).Return(val)

	h := &NavigatorSearchHandler{
		logger: log,
		em:     em,
	}

	pck := &message.NavigatorSearchPacket{
		Query: query,
		View:  view,
	}

	con := setupMockConn()

	h.Handle(context.Background(), pck, con)

	em.AssertExpectations(t)

	return h, buf
}

// TestNavigatorSearchHandler_Fire checks if event firing error is logged and handled.
func TestNavigatorSearchHandler_Fire(t *testing.T) {
	setupTestEnvironment(t, "error query", "private", nil)
}

// TestNavigatorSearchHandler_FireError checks if event firing error is logged and handled.
func TestNavigatorSearchHandler_FireError(t *testing.T) {
	_, buf := setupTestEnvironment(t, "error query", "private", errors.New("fire error"))
	assert.Contains(t, buf.String(), "fire error")
}

// TestNavigatorSearchHandler_InvalidPacket checks if invalid packet type is handled correctly.
func TestNavigatorSearchHandler_InvalidPacket(t *testing.T) {
	log, buf := util.CreateTestLogger()

	h := &NavigatorSearchHandler{
		logger: log,
	}

	invalidPck := &message2.AuthTicketPacket{}
	con := setupMockConn()

	h.Handle(context.Background(), invalidPck, con)

	assert.Contains(t, buf.String(), "cannot cast navigator search packet, skipping processing")
}
