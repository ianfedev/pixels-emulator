package server

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"pixels-emulator/core/config"
	"pixels-emulator/core/event"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
	"pixels-emulator/core/scheduler"
	"pixels-emulator/room"
	"pixels-emulator/user"
)

// Server defines the behavior of a server instance.
type Server interface {
	// Reload reloads the processors, handlers, cron jobs, and events.
	Reload(log bool)

	// Stop stops the server and closes active connections.
	Stop() error

	// Config returns the configuration of the server.
	Config() *config.Config

	// Logger returns the logger of the server.
	Logger() *zap.Logger

	// ConnStore returns the connection manager of the server.
	ConnStore() protocol.ConnectionManager

	// Scheduler returns the scheduler of the server.
	Scheduler() scheduler.Scheduler

	// PacketProcessors returns the packet processors of the server.
	PacketProcessors() registry.ProcessorRegistry

	// PacketHandlers returns the packet handlers of the server.
	PacketHandlers() registry.HandlerRegistry

	// EventManager returns the event manager of the server.
	EventManager() event.Manager

	// Database returns the database connection of the server.
	Database() *gorm.DB

	// RoomStore provides the in-memory loaded rooms.
	RoomStore() *room.Store

	// UserStore provides the in-memory loaded users.
	UserStore() *user.Store
}
