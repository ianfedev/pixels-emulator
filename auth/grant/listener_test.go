package grant_test

import (
	"github.com/stretchr/testify/mock"
	"pixels-emulator/auth/event"
	"pixels-emulator/auth/grant"
	mockevent "pixels-emulator/core/event/mock"
	mockproto "pixels-emulator/core/protocol/mock"
	"pixels-emulator/core/server"
	mockserver "pixels-emulator/core/server/mock"
	"testing"
)

func Test_OnAuthGranted(t *testing.T) {

	sv := &mockserver.Server{}
	con := &mockproto.MockConnection{}

	connStore := &mockproto.MockConnectionManager{}
	connStore.On("GetConnection", "1").Return(con)

	em := &mockevent.MockEventManager{}

	sv.On("ConnStore").Return(connStore)
	sv.On("EventManager").Return(em)
	con.On("SendPacket", mock.Anything).Return(nil)

	server.UpdateInstance(sv)
	ev := event.NewEvent(1, 0, make(map[string]string))

	grant.OnAuthGrantEvent(ev)

	connStore.AssertExpectations(t)
	con.AssertExpectations(t)
	con.AssertCalled(t, "SendPacket", mock.Anything)
}
