package handler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	grant "pixels-emulator/auth/event"
	"pixels-emulator/auth/message"
	"pixels-emulator/core/config"
	"pixels-emulator/core/database"
	mockdb "pixels-emulator/core/database/mock"
	"pixels-emulator/core/event"
	mockevent "pixels-emulator/core/event/mock"
	"pixels-emulator/core/model"
	mockproto "pixels-emulator/core/protocol/mock"
	"pixels-emulator/core/util"
	"strconv"
	"testing"
)

func setupHandler(
	logger *zap.Logger,
	em event.Manager,
	ssoSvc database.DataService[model.SSOTicket],
	userSvc database.DataService[model.User],
	cfg *config.Config) *AuthTicketHandler {

	return &AuthTicketHandler{
		logger:  logger,
		em:      em,
		ssoSvc:  ssoSvc,
		userSvc: userSvc,
		cfg:     cfg,
	}
}

func TestAuthTicketHandler_Handle(t *testing.T) {

	log, _ := util.CreateTestLogger()
	demoId := "1"
	idNum, _ := strconv.Atoi(demoId)

	cfg := &config.Config{
		Server: config.ServerConfig{
			Environment: "PRODUCTION",
		},
	}

	con := &mockproto.MockConnection{}
	con.On("GrantIdentifier", mock.Anything).Return(nil)
	con.On("Identifier").Return(demoId)

	em := &mockevent.MockEventManager{}
	em.On("Fire", grant.AuthGrantEventName, mock.Anything).Return(nil)

	ssoSvc := &mockdb.ModelServiceMock[model.SSOTicket]{}
	ssoSvc.On("FindByQuery", mock.MatchedBy(func(q map[string]interface{}) bool {
		return q["ticket"] == demoId
	})).Return(util.MockAsyncResponse([]model.SSOTicket{{UserID: uint(idNum)}}, nil))

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

}

func TestAuthTicketHandler_HandleDebug(t *testing.T) {

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

	em := &mockevent.MockEventManager{}
	em.On("Fire", grant.AuthGrantEventName, mock.Anything).Return(nil)

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
	assert.Contains(t, buf.String(), "Attempting to log in debug mode, SSO ticket will be taken as user id, switch to production to prevent this")

}

func TestAuthTicketHandler_QueryError(t *testing.T) {

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
	})).Return(util.MockAsyncResponse([]model.SSOTicket{}, errors.New("query error")))

	userSvc := &mockdb.ModelServiceMock[model.User]{}

	pck := &message.AuthTicketPacket{
		Ticket: demoId,
		Time:   1,
	}

	handler := setupHandler(log, em, ssoSvc, userSvc, cfg)

	handler.Handle(pck, con)
	ssoSvc.AssertExpectations(t)
	assert.Contains(t, buf.String(), "query error")

}

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
	})).Return(util.MockAsyncResponse([]model.SSOTicket{}, nil))

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
