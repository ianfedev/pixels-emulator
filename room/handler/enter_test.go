package handler

import (
	"bytes"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mockevent "pixels-emulator/core/event/mock"
	mockproto "pixels-emulator/core/protocol/mock"
	"pixels-emulator/core/server"
	mockserver "pixels-emulator/core/server/mock"
	"pixels-emulator/core/util"
	healthMsg "pixels-emulator/healthcheck/message"
	"pixels-emulator/room/event"
	"pixels-emulator/room/message"
	"testing"
)

var pck = &message.RoomEnterPacket{
	RoomId:   1,
	Password: "yes",
}

// setupMockConn creates a basic connection to be modified
func setupMockConn(id string, dErr error) *mockproto.MockConnection {
	con := &mockproto.MockConnection{}
	con.On("GrantIdentifier", mock.Anything).Return(nil)
	con.On("Identifier").Return(id)
	con.On("Dispose").Return(dErr)
	return con
}

// setupTestEnvironment sets a new testing environment for handler.
func setupTestEnvironment(em *mockevent.MockEventManager, err string) (*RoomEnterHandler, *bytes.Buffer) {

	log, buf := util.CreateTestLogger()
	var e error
	if err != "" {
		e = errors.New(err)
	}

	em.On("Fire", event.RoomJoinEventName, mock.Anything).Return(e)
	handler := &RoomEnterHandler{
		logger: log,
		em:     em,
	}

	return handler, buf

}

// TestAuthTicketHandler_Create tests if the ticket handler is created correctly.
func TestAuthTicketHandler_Create(t *testing.T) {
	sv := &mockserver.Server{}
	sv.On("Logger").Return(util.CreateTestLogger())
	sv.On("EventManager").Return(&mockevent.MockEventManager{})
	server.UpdateInstance(sv)
	NewRoomEnter()
	t.Cleanup(func() {
		server.ResetInstance()
	})
}

func TestRoomEnterHandler_Handle(t *testing.T) {
	em := &mockevent.MockEventManager{}
	handler, _ := setupTestEnvironment(em, "")
	conn := setupMockConn("1", nil)
	handler.Handle(context.Background(), pck, conn)
	em.AssertExpectations(t)
}

func TestRoomEnterHandler_HandleErr(t *testing.T) {
	em := &mockevent.MockEventManager{}
	handler, buff := setupTestEnvironment(em, "mock event error")
	conn := setupMockConn("1", nil)
	handler.Handle(context.Background(), pck, conn)
	em.AssertExpectations(t)
	assert.Contains(t, buff.String(), "mock event error")
}

func TestRoomEnterHandler_HandleInvalidPacket(t *testing.T) {
	em := &mockevent.MockEventManager{}
	handler, buff := setupTestEnvironment(em, "")
	conn := setupMockConn("1", nil)
	handler.Handle(context.Background(), &healthMsg.HelloPacket{}, conn)
	assert.Contains(t, buff.String(), "cannot cast navigator search packet, skipping processing")
}

func TestRoomEnterHandler_HandleInvalidDispose(t *testing.T) {
	em := &mockevent.MockEventManager{}
	handler, buff := setupTestEnvironment(em, "")
	conn := setupMockConn("1", errors.New("dispose error"))
	handler.Handle(context.Background(), &healthMsg.HelloPacket{}, conn)
	assert.Contains(t, buff.String(), "dispose error")
}
