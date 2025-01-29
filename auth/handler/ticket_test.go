package handler

import (
	mockevent "pixels-emulator/core/event/mock"
	mockproto "pixels-emulator/core/protocol/mock"
	mockserver "pixels-emulator/core/server/mock"
	"pixels-emulator/core/util"
	"testing"
)

func TestAuthTicketHandler_Handle(t *testing.T) {

	sv := &mockserver.Server{}
	connStore := &mockproto.MockConnectionManager{}
	con := &mockproto.MockConnection{}
	em := &mockevent.MockEventManager{}

	log, buf := util.CreateTestLogger()

	sv.On("ConnStore").Return(connStore)
	sv.On("Logger").Return(log)

}
