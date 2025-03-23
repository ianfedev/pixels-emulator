package handler

import (
	"bytes"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	grant "pixels-emulator/auth/event"
	"pixels-emulator/auth/message"
	"pixels-emulator/core/config"
	"pixels-emulator/core/database"
	mockdb "pixels-emulator/core/database/mock"
	mockevent "pixels-emulator/core/event/mock"
	"pixels-emulator/core/model"
	mockproto "pixels-emulator/core/protocol/mock"
	"pixels-emulator/core/server"
	mockserver "pixels-emulator/core/server/mock"
	"pixels-emulator/core/util"
	"testing"
)

// setupMockConn creates a basic connection to be modified
func setupMockConn(id string, dErr error) *mockproto.MockConnection {
	con := &mockproto.MockConnection{}
	con.On("GrantIdentifier", mock.Anything).Return(nil)
	con.On("Identifier").Return(id)
	con.On("Dispose").Return(dErr)
	return con
}

// setupTestEnvironment creates a parametrized test environment with multiple uses to prevent code duplication on testings.
// For example: Checking all the model services, modifying query and user get mock response, environment and some errors.
// Please refer to specific tests to find the usage.
func setupTestEnvironment(
	t *testing.T,
	env string,
	id string,
	qRes interface{},
	gRes interface{},
	aUser bool,
	aSso bool,
	aEvent bool,
	fErr error) (*AuthTicketHandler, *bytes.Buffer) {

	log, buf := util.CreateTestLogger()

	cfg := &config.Config{
		Server: config.ServerConfig{
			Environment: env,
		},
	}

	con := setupMockConn(id, nil)

	em := &mockevent.MockEventManager{}
	em.On("Fire", grant.AuthGrantEventName, mock.Anything).Return(fErr)

	ssoSvc := &mockdb.ModelServiceMock[model.SSOTicket]{}
	ssoSvc.On("FindByQuery", mock.MatchedBy(func(q map[string]interface{}) bool {
		return q["ticket"] == id
	})).Return(qRes)

	userSvc := &mockdb.ModelServiceMock[model.User]{}
	userSvc.On("Get", mock.Anything).Return(gRes)

	ath := &AuthTicketHandler{
		logger:  log,
		em:      em,
		ssoSvc:  ssoSvc,
		userSvc: userSvc,
		cfg:     cfg,
	}

	pck := &message.AuthTicketPacket{
		Ticket: id,
		Time:   1,
	}

	ath.Handle(context.Background(), pck, con)

	if aUser {
		userSvc.AssertExpectations(t)
	}

	if aSso {
		ssoSvc.AssertExpectations(t)
	}

	if aEvent {
		em.AssertExpectations(t)
	}

	return ath, buf
}

// TestAuthTicketHandler_Create tests if the ticket handler is created correctly.
func TestAuthTicketHandler_Create(t *testing.T) {
	sv := &mockserver.Server{}
	sv.On("Database").Return(&gorm.DB{})
	sv.On("Logger").Return(util.CreateTestLogger())
	sv.On("EventManager").Return(&mockevent.MockEventManager{})
	sv.On("Config").Return(&config.Config{})
	server.UpdateInstance(sv)
	NewAuthTicket()
	t.Cleanup(func() {
		server.ResetInstance()
	})
}

// TestAuthTicketHandler_Handle tests if production behaviour is successful and event is called if everything is correct.
func TestAuthTicketHandler_Handle(t *testing.T) {
	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{{UserID: uint(idNum)}}, nil)
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, nil)
	setupTestEnvironment(t, "PRODUCTION", "1", qRes, gRes, true, true, true, nil)
}

// TestAuthTicketHandler_DisposeError tests if error is logged when connection is not disposed correctly.
func TestAuthTicketHandler_DisposeError(t *testing.T) {

	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{{UserID: uint(idNum)}}, nil)
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, nil)
	ath, buf := setupTestEnvironment(t, "PRODUCTION", "1", qRes, gRes, true, true, true, nil)

	incPck := &message.AuthOkPacket{}
	con := setupMockConn("1", errors.New("dispose error"))

	ath.Handle(context.Background(), incPck, con)
	assert.Contains(t, buf.String(), "cannot cast packet, skipping processing")

	pck := &message.AuthTicketPacket{
		Ticket: "1",
		Time:   1,
	}
	ath.Handle(context.Background(), pck, con)

}

// TestAuthTicketHandler_HandleDebug check if correct development procedure like security logging and normal
// packet handling is present.
func TestAuthTicketHandler_HandleDebug(t *testing.T) {
	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{{UserID: uint(idNum)}}, nil)
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, nil)
	_, buf := setupTestEnvironment(t, "DEVELOPMENT", "1", qRes, gRes, true, false, true, nil)
	assert.Contains(t, buf.String(), "Attempting to log in debug mode, SSO ticket will be taken as user id, switch to production to prevent this")
}

// TestAuthTicketHandler_QueryError checks if query error is logged and handled.
func TestAuthTicketHandler_QueryError(t *testing.T) {
	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{}, errors.New("query error"))
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, nil)
	_, buf := setupTestEnvironment(t, "PRODUCTION", "1", qRes, gRes, false, true, false, nil)
	assert.Contains(t, buf.String(), "query error")
}

// TestAuthTicketHandler_GetError checks if get error is logged and handled.
func TestAuthTicketHandler_GetError(t *testing.T) {
	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{}, nil)
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, errors.New("get error"))
	_, buf := setupTestEnvironment(t, "DEVELOPMENT", "1", qRes, gRes, true, false, false, nil)
	assert.Contains(t, buf.String(), "get error")
}

// TestAuthTicketHandler_DuplicatedSessionError checks if session duplication error is logged and handled.
func TestAuthTicketHandler_DuplicatedSessionError(t *testing.T) {
	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{{}, {}}, nil)
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, nil)
	_, buf := setupTestEnvironment(t, "PRODUCTION", "1", qRes, gRes, false, true, false, nil)
	assert.Contains(t, buf.String(), "session is being duplicated")
}

// TestAuthTicketHandler_EmptySession checks if empty session (No SSO) error is logged and handled.
func TestAuthTicketHandler_EmptySession(t *testing.T) {
	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{}, nil)
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, nil)
	_, buf := setupTestEnvironment(t, "PRODUCTION", "1", qRes, gRes, false, true, false, nil)
	assert.Contains(t, buf.String(), "session is being created with not valid ticket")
}

// TestAuthTicketHandler_FireError checks if event firing error is logged and handled.
func TestAuthTicketHandler_FireError(t *testing.T) {
	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{{}}, nil)
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, nil)
	_, buf := setupTestEnvironment(t, "DEVELOPMENT", "1", qRes, gRes, false, false, false, errors.New("fire error"))
	assert.Contains(t, buf.String(), "fire error")
}

// TestAuthTicketHandler_ParseError checks if id error parsing is logged and handled.
func TestAuthTicketHandler_ParseError(t *testing.T) {
	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{{}}, nil)
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, nil)
	_, buf := setupTestEnvironment(t, "DEVELOPMENT", "tricky", qRes, gRes, false, false, false, nil)
	assert.Contains(t, buf.String(), "invalid syntax")
}
