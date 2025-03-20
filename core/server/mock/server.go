package mock

import (
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"pixels-emulator/core/config"
	"pixels-emulator/core/event"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
	"pixels-emulator/core/scheduler"
	"pixels-emulator/room"
)

// Server is a mock implementation of the Server interface.
type Server struct {
	mock.Mock
}

// Reload simulates the Reload method of the Server interface.
func (m *Server) Reload(log bool) {
	m.Called(log)
}

// Stop simulates the Stop method of the Server interface.
func (m *Server) Stop() error {
	args := m.Called()
	return args.Error(0)
}

// Config simulates the Config method of the Server interface.
func (m *Server) Config() *config.Config {
	args := m.Called()
	return args.Get(0).(*config.Config)
}

// Logger simulates the Logger method of the Server interface.
func (m *Server) Logger() *zap.Logger {
	args := m.Called()
	return args.Get(0).(*zap.Logger)
}

// ConnStore simulates the ConnStore method of the Server interface.
func (m *Server) ConnStore() protocol.ConnectionManager {
	args := m.Called()
	return args.Get(0).(protocol.ConnectionManager)
}

// Scheduler simulates the Scheduler method of the Server interface.
func (m *Server) Scheduler() scheduler.Scheduler {
	args := m.Called()
	return args.Get(0).(scheduler.Scheduler)
}

// PacketProcessors simulates the PacketProcessors method of the Server interface.
func (m *Server) PacketProcessors() registry.ProcessorRegistry {
	args := m.Called()
	return args.Get(0).(registry.ProcessorRegistry)
}

// PacketHandlers simulates the PacketHandlers method of the Server interface.
func (m *Server) PacketHandlers() registry.HandlerRegistry {
	args := m.Called()
	return args.Get(0).(registry.HandlerRegistry)
}

// EventManager simulates the EventManager method of the Server interface.
func (m *Server) EventManager() event.Manager {
	args := m.Called()
	return args.Get(0).(event.Manager)
}

// Database simulates the Database method of the Server interface.
func (m *Server) Database() *gorm.DB {
	args := m.Called()
	return args.Get(0).(*gorm.DB)
}

// RoomStore simulates the RoomStore method of the Server instance.
func (m *Server) RoomStore() *room.Store {
	args := m.Called()
	return args.Get(0).(*room.Store)
}
