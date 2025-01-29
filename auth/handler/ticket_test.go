package handler

import (
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"pixels-emulator/auth/message"
	"pixels-emulator/core/config"
	"pixels-emulator/core/database"
	mockdb "pixels-emulator/core/database/mock"
	"pixels-emulator/core/event"
	mockevent "pixels-emulator/core/event/mock"
	"pixels-emulator/core/model"
	"pixels-emulator/core/protocol"
	mockproto "pixels-emulator/core/protocol/mock"
	"pixels-emulator/core/util"
	"testing"
)

func setupHandler(
	logger *zap.Logger,
	connStore protocol.ConnectionManager,
	em event.Manager,
	ssoSvc database.DataService[model.SSOTicket],
	userSvc database.DataService[model.User],
	cfg *config.Config) *AuthTicketHandler {

	return &AuthTicketHandler{
		logger:    logger,
		connStore: connStore,
		em:        em,
		ssoSvc:    ssoSvc,
		userSvc:   userSvc,
		cfg:       cfg,
	}
}

func TestAuthTicketHandler_Handle(t *testing.T) {

	log, _ := util.CreateTestLogger()

	cfg := &config.Config{
		Server: config.ServerConfig{
			Environment: "PRODUCTION",
		},
	}

	con := &mockproto.MockConnection{}
	connStore := &mockproto.MockConnectionManager{}
	em := &mockevent.MockEventManager{}

	connStore.On("GetConnection", mock.Anything).Return(con, true)

	ssoSvc := &mockdb.ModelServiceMock[model.SSOTicket]{}

	queryChan := make(chan struct {
		Entities []model.SSOTicket
		Error    error
	}, 1)

	go func() {
		queryChan <- struct {
			Entities []model.SSOTicket
			Error    error
		}{
			Entities: []model.SSOTicket{{UserID: 1}},
			Error:    nil,
		}
		close(queryChan)
	}()

	ssoSvc.On("FindByQuery", mock.MatchedBy(func(q map[string]interface{}) bool {
		return q["ticket"] == "1"
	})).Return((<-chan struct {
		Entities []model.SSOTicket
		Error    error
	})(queryChan))

	userSvc := &mockdb.ModelServiceMock[model.User]{}

	userChan := make(chan struct {
		Entities []model.SSOTicket
		Error    error
	}, 1)

	go func() {
		queryChan <- struct {
			Entities []model.SSOTicket
			Error    error
		}{
			Entities: []model.SSOTicket{{UserID: 1}},
			Error:    nil,
		}
		close(queryChan)
	}()

	userSvc.On("Get", mock.Anything).Return()

	pck := &message.AuthTicketPacket{
		Ticket: "1",
		Time:   1,
	}

	handler := setupHandler(
		log, connStore, em, ssoSvc, userSvc, cfg)

	handler.Handle(pck, con)

	ssoSvc.AssertCalled(t, "FindByQuery", mock.MatchedBy(func(q map[string]interface{}) bool {
		return q["ticket"] == "1"
	}))
}
