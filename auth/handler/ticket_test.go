package handler

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	grant "pixels-emulator/auth/event"
	"pixels-emulator/auth/message"
	"pixels-emulator/core/config"
	"pixels-emulator/core/database"
	mockdb "pixels-emulator/core/database/mock"
	mockevent "pixels-emulator/core/event/mock"
	"pixels-emulator/core/model"
	mockproto "pixels-emulator/core/protocol/mock"
	"pixels-emulator/core/util"
	"testing"
)

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

	con := &mockproto.MockConnection{}
	con.On("GrantIdentifier", mock.Anything).Return(nil)
	con.On("Identifier").Return(id)
	con.On("Dispose").Return(nil)

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

	ath.Handle(pck, con)

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

func TestAuthTicketHandler_Handle(t *testing.T) {
	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{{UserID: uint(idNum)}}, nil)
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, nil)
	setupTestEnvironment(t, "PRODUCTION", "1", qRes, gRes, true, true, true, nil)
}

func TestAuthTicketHandler_HandleDebug(t *testing.T) {
	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{{UserID: uint(idNum)}}, nil)
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, nil)
	_, buf := setupTestEnvironment(t, "DEVELOPMENT", "1", qRes, gRes, true, false, true, nil)
	assert.Contains(t, buf.String(), "Attempting to log in debug mode, SSO ticket will be taken as user id, switch to production to prevent this")
}

func TestAuthTicketHandler_QueryError(t *testing.T) {
	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{}, errors.New("query error"))
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, nil)
	_, buf := setupTestEnvironment(t, "PRODUCTION", "1", qRes, gRes, false, true, false, nil)
	assert.Contains(t, buf.String(), "query error")
}

func TestAuthTicketHandler_GetError(t *testing.T) {
	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{}, nil)
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, errors.New("get error"))
	_, buf := setupTestEnvironment(t, "DEVELOPMENT", "1", qRes, gRes, true, false, false, nil)
	assert.Contains(t, buf.String(), "get error")
}

func TestAuthTicketHandler_DuplicatedSessionError(t *testing.T) {
	idNum := 1
	qRes := util.MockAsyncResponse([]model.SSOTicket{{}, {}}, nil)
	gRes := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, nil)
	_, buf := setupTestEnvironment(t, "DEVELOPMENT", "1", qRes, gRes, false, true, false, nil)
	assert.Contains(t, buf.String(), "session is being duplicated")
}

/*

func TestAuthTicketHandler_DuplicatedSession(t *testing.T) {

	log, buf := util.CreateTestLogger()
	demoId := "1"

	cfg := &config.Config{
		Server: config.ServerConfig{
			Environment: "PRODUCTION",
		},
	}

	con := &mockproto.MockConnection{}
	con.On("Dispose").Return(nil)

	em := &mockevent.MockEventManager{}

	ssoSvc := &mockdb.ModelServiceMock[model.SSOTicket]{}
	ssoSvc.On("FindByQuery", mock.MatchedBy(func(q map[string]interface{}) bool {
		return q["ticket"] == demoId
	})).Return(util.MockAsyncResponse([]model.SSOTicket{{}, {}}, nil))

	userSvc := &mockdb.ModelServiceMock[model.User]{}

	pck := &message.AuthTicketPacket{
		Ticket: demoId,
		Time:   1,
	}

	handler := setupHandler(log, em, ssoSvc, userSvc, cfg)

	handler.Handle(pck, con)
	ssoSvc.AssertExpectations(t)
	assert.Contains(t, buf.String(), "session is being duplicated")

}

func TestAuthTicketHandler_InvalidSession(t *testing.T) {

	log, buf := util.CreateTestLogger()
	demoId := "1"

	cfg := &config.Config{
		Server: config.ServerConfig{
			Environment: "PRODUCTION",
		},
	}

	con := &mockproto.MockConnection{}
	con.On("Dispose").Return(nil)

	em := &mockevent.MockEventManager{}

	ssoSvc := &mockdb.ModelServiceMock[model.SSOTicket]{}
	ssoSvc.On("FindByQuery", mock.MatchedBy(func(q map[string]interface{}) bool {
		return q["ticket"] == demoId
	})).Return(util.MockAsyncResponse([]model.SSOTicket{}, nil))

	userSvc := &mockdb.ModelServiceMock[model.User]{}

	pck := &message.AuthTicketPacket{
		Ticket: demoId,
		Time:   1,
	}

	handler := setupHandler(log, em, ssoSvc, userSvc, cfg)

	handler.Handle(pck, con)
	ssoSvc.AssertExpectations(t)
	assert.Contains(t, buf.String(), "session is being created with not valid ticket")

}

func TestAuthTicketHandler_FireError(t *testing.T) {

	log, buf := util.CreateTestLogger()
	demoId := "1"
	idNum, _ := strconv.Atoi(demoId)

	cfg := &config.Config{
		Server: config.ServerConfig{
			Environment: "DEVELOPMENT",
		},
	}

	con := &mockproto.MockConnection{}
	con.On("GrantIdentifier", mock.Anything).Return(nil)
	con.On("Identifier").Return(demoId)
	con.On("Dispose").Return(nil)

	em := &mockevent.MockEventManager{}
	em.On("Fire", grant.AuthGrantEventName, mock.Anything).Return(errors.New("test fire error"))

	ssoSvc := &mockdb.ModelServiceMock[model.SSOTicket]{}

	userSvc := &mockdb.ModelServiceMock[model.User]{}
	userSvc.On("Get", mock.MatchedBy(func(id uint) bool {
		return id == uint(idNum)
	})).Return(util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(idNum)}}, nil))

	pck := &message.AuthTicketPacket{
		Ticket: demoId,
		Time:   1,
	}

	handler := setupHandler(log, em, ssoSvc, userSvc, cfg)

	handler.Handle(pck, con)
	ssoSvc.AssertExpectations(t)
	userSvc.AssertExpectations(t)
	em.AssertExpectations(t)
	assert.Contains(t, buf.String(), "test fire error")

}

*/
