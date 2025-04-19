package grant_test

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/database"
	mockdb "pixels-emulator/core/database/mock"
	oEvent "pixels-emulator/core/event"
	"pixels-emulator/core/model"
	mockproto "pixels-emulator/core/protocol/mock"
	scmock "pixels-emulator/core/scheduler/mock"
	mockserver "pixels-emulator/core/server/mock"
	"pixels-emulator/core/store"
	"pixels-emulator/core/util"
	"pixels-emulator/user"
	mockuser "pixels-emulator/user/mock"

	"bytes"
	"github.com/stretchr/testify/mock"
	"pixels-emulator/auth/event"
	"pixels-emulator/auth/grant"
	"pixels-emulator/core/server"
	"testing"
)

// setupMocks initializes mock server, connection, and logger for testing.
func setupMocks(exist bool) (*mockserver.Server, *mockproto.MockConnection, *bytes.Buffer) {
	sv := &mockserver.Server{}
	connStore := &mockproto.MockConnectionManager{}
	con := &mockproto.MockConnection{}
	sc := &scmock.MockScheduler{}
	log, buf := util.CreateTestLogger()

	us := &mockuser.Store{}
	us.On("Records", mock.Anything, mock.Anything).Return(store.NewMemoryStore[*user.Player]())

	grant.UserDatabaseFunc = func() database.DataService[model.User] {
		ssoSvc := &mockdb.ModelServiceMock[model.User]{}
		res := util.MockAsyncResponse(&model.User{BaseModel: database.BaseModel{ID: uint(1)}}, nil)
		ssoSvc.On("Get", mock.Anything, mock.Anything).Return(res)
		return ssoSvc
	}

	connStore.On("GetConnection", mock.Anything).Return(con, exist)
	sv.On("UserStore").Return(us)
	sv.On("ConnStore").Return(connStore)
	sv.On("Scheduler").Return(sc)
	sv.On("Logger").Return(log)

	return sv, con, buf
}

// Test_OnAuthGrantEvent tests OnAuthGrantEvent with a valid connection.
func Test_OnAuthGrantEvent(t *testing.T) {
	sv, con, _ := setupMocks(true)
	authFunc := grant.ProvideAuth()

	assert.IsType(t, func(event oEvent.Event) {}, authFunc)

	con.On("SendPacket", mock.Anything).Return(nil)

	server.UpdateInstance(sv)
	authFunc(event.NewEvent(1, 0, nil))

	con.AssertExpectations(t)

	t.Cleanup(func() {
		server.ResetInstance()
		grant.UserDatabaseFunc = grant.GetUserDatabase
	})
}

// Test_OnAuthGrantEvent_Cancelled tests OnAuthGrantEvent when the event is cancelled.
func Test_OnAuthGrantEvent_Cancelled(t *testing.T) {
	sv, con, _ := setupMocks(true)

	con.On("Dispose").Return(nil)

	authEv := event.NewEvent(1, 0, nil)
	authEv.CancellableEvent.Cancelled = true

	server.UpdateInstance(sv)
	grant.OnAuthGrantEvent(authEv)

	con.AssertExpectations(t)

	t.Cleanup(func() {
		server.ResetInstance()
		grant.UserDatabaseFunc = grant.GetUserDatabase
	})
}

// Test_OnAuthGrantEvent_InvalidConnection tests OnAuthGrantEvent when the connection is invalid.
func Test_OnAuthGrantEvent_InvalidConnection(t *testing.T) {

	sv, _, buf := setupMocks(false)
	authEv := event.NewEvent(1, 0, nil)

	server.UpdateInstance(sv)
	grant.OnAuthGrantEvent(authEv)

	err := "connection not found"
	if !bytes.Contains(buf.Bytes(), []byte("connection not found")) {
		t.Errorf("Expected '%s' in logs, but it was not found.", err)
	}

	t.Cleanup(func() {
		server.ResetInstance()
		grant.UserDatabaseFunc = grant.GetUserDatabase
	})

}

// Test_OnAuthGrant_Event_InvalidEvent tests OnAuthGrantEvent when an invalid event is provided.
func Test_OnAuthGrant_Event_InvalidEvent(t *testing.T) {

	sv, _, buf := setupMocks(false)
	authEv := oEvent.New(0, make(map[string]string))

	server.UpdateInstance(sv)
	grant.OnAuthGrantEvent(authEv)

	err := "event proportioned was not authentication"
	if !bytes.Contains(buf.Bytes(), []byte(err)) {
		t.Errorf("Expected '%s' in logs, but it was not found.", err)
	}

	t.Cleanup(func() {
		server.ResetInstance()
		grant.UserDatabaseFunc = grant.GetUserDatabase
	})
}

// TestProvideAuth tests if event is provided correctly.
func TestProvideAuth(t *testing.T) {
	authFunc := grant.ProvideAuth()
	assert.IsType(t, func(event oEvent.Event) {}, authFunc)
}
