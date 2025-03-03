package handler

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pixels-emulator/core/protocol"
	mockproto "pixels-emulator/core/protocol/mock"
	"pixels-emulator/core/server"
	mockserver "pixels-emulator/core/server/mock"
	"pixels-emulator/core/util"
	healthMsg "pixels-emulator/healthcheck/message"
	"pixels-emulator/navigator/message"
	"testing"
)

func setupTestServer() (*mockserver.Server, *bytes.Buffer) {
	sv := &mockserver.Server{}
	log, buf := util.CreateTestLogger()
	sv.On("Logger").Return(log)
	server.UpdateInstance(sv)
	return sv, buf
}

func setupTestConnection() *mockproto.MockConnection {
	con := &mockproto.MockConnection{}
	con.On("Identifier").Return("1")
	con.On("SendPacket", mock.Anything).Return(nil)
	return con
}

func Test_NewNavigatorInit(t *testing.T) {
	setupTestServer()
	NewNavigatorInit()
	t.Cleanup(func() {
		server.ResetInstance()
	})
}

func Test_Handle_ValidPacket(t *testing.T) {
	_, buf := setupTestServer()
	con := setupTestConnection()

	pck := message.ComposeNavigatorInit(protocol.RawPacket{})
	handler := NewNavigatorInit()

	handler.Handle(pck, con)

	assert.Contains(t, buf.String(), "Navigator fired by user")
	con.AssertExpectations(t)

	t.Cleanup(func() {
		server.ResetInstance()
	})
}

func Test_Handle_InvalidPacket(t *testing.T) {
	_, buf := setupTestServer()
	con := setupTestConnection()

	pck := &healthMsg.HelloPacket{}
	handler := NewNavigatorInit()

	handler.Handle(pck, con)

	fmt.Println(buf.String())
	assert.Contains(t, buf.String(), "cannot cast ping packet, skipping processing")
	t.Cleanup(func() {
		server.ResetInstance()
	})
}
