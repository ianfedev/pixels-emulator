package grant_test

import (
	"github.com/stretchr/testify/assert"
	oEvent "pixels-emulator/core/event"
	mockproto "pixels-emulator/core/protocol/mock"
	mockserver "pixels-emulator/core/server/mock"

	"bytes"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	log, buf := createLogger()

	connStore.On("GetConnection", mock.Anything).Return(con, exist)
	sv.On("ConnStore").Return(connStore)
	sv.On("Logger").Return(log)

	return sv, con, buf
}

// createLogger creates a zap logger with a buffer for capturing log output.
func createLogger() (*zap.Logger, *bytes.Buffer) {
	var buf bytes.Buffer
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, zapcore.AddSync(&buf), zapcore.DebugLevel)
	logger := zap.New(core)
	return logger, &buf
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
	})
}

// TestProvideAuth tests if event is provided correctly.
func TestProvideAuth(t *testing.T) {
	authFunc := grant.ProvideAuth()
	assert.IsType(t, func(event oEvent.Event) {}, authFunc)
}
