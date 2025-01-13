package server

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"pixels-emulator/core/config"
	"pixels-emulator/core/database"
	"pixels-emulator/core/event"
	"pixels-emulator/core/log"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
	"pixels-emulator/core/scheduler"
	"pixels-emulator/core/setup"
	"sync"
	"time"
)

// MainServer is the implementation of Server providing global state management.
type MainServer struct {
	config           *config.Config             // config provides the loaded configuration from file.
	logger           *zap.Logger                // logger provides the server logger.
	connStore        protocol.ConnectionManager // connStore provides all the active connections on the server.
	scheduler        scheduler.Scheduler        // scheduler provides the scheduler instance.
	packetProcessors registry.ProcessorRegistry // packetProcessors provides all the packets available to be processed.
	packetHandlers   registry.HandlerRegistry   // packetHandlers provides the handlers of the packets processed.
	eventManager     event.Manager              // eventManager provides the event management system global instance.
	database         *gorm.DB                   // database provides a connection for ORM.
}

var (
	once     sync.Once
	instance Server
)

// GetServer provides the only server instance.
func GetServer() Server {
	once.Do(func() {
		if instance == nil {
			UpdateInstance(setupServer())
		}
	})
	return instance
}

func UpdateInstance(server Server) {
	if instance == nil {
		instance = server
	}
}

// Reload performs the reloading of processors, handlers, cron jobs, and events.
// It logs the duration of the reload process and alerts if any issues occur.
func (s *MainServer) Reload(log bool) {
	startTime := time.Now()

	// Perform the reload steps

	duration := time.Since(startTime)

	// Check if there was any alert or error during the reload process
	if log {
		s.logger.Info("Reload completed", zap.Duration("duration", duration))

		if duration > time.Second*5 {
			// Log a warning if reload takes longer than expected
			s.logger.Warn("Reload took longer than expected", zap.Duration("duration", duration))
		} else {
			// Otherwise, log the successful reload duration
			s.logger.Info("Reload finished within expected time", zap.Duration("duration", duration))
		}
	}
}

// Stop stops the server and closes active connections.
func (s *MainServer) Stop() error {
	s.connStore.CloseActive()
	return nil
}

// Config returns the configuration of the server.
func (s *MainServer) Config() *config.Config {
	return s.config
}

// Logger returns the logger of the server.
func (s *MainServer) Logger() *zap.Logger {
	return s.logger
}

// ConnStore returns the connection manager of the server.
func (s *MainServer) ConnStore() protocol.ConnectionManager {
	return s.connStore
}

// Scheduler returns the scheduler of the server.
func (s *MainServer) Scheduler() scheduler.Scheduler {
	return s.scheduler
}

// PacketProcessors returns the packet processors of the server.
func (s *MainServer) PacketProcessors() registry.ProcessorRegistry {
	return s.packetProcessors
}

// PacketHandlers returns the packet handlers of the server.
func (s *MainServer) PacketHandlers() registry.HandlerRegistry {
	return s.packetHandlers
}

// EventManager returns the event manager of the server.
func (s *MainServer) EventManager() event.Manager {
	return s.eventManager
}

// Database returns the database connection of the server.
func (s *MainServer) Database() *gorm.DB {
	return s.database
}

func setupServer() *MainServer {

	var setupErr error
	var logger *zap.Logger
	defer func() {
		if setupErr != nil {
			logger.Error("Error while starting server", zap.Error(setupErr))
			os.Exit(1)
		}
	}()

	logger = log.CreateTempLogger()
	logger.Info("Starting Pixels emulator")

	setupErr = config.CreateDefaultConfig("config.ini", logger)
	if setupErr != nil {
		return nil
	}

	cfg, setupErr := setup.Config("config.ini", logger)
	if setupErr != nil {
		return nil
	}

	log.SetupLogger(cfg)
	logger = zap.L()
	logger.Debug("Logger instantiated")

	db, setupErr := database.SetupDatabase(cfg, zap.L())
	if setupErr != nil {
		return nil
	}

	setupErr = setup.ModelMigration(zap.L(), db)
	if setupErr != nil {
		return nil
	}

	connStore := protocol.NewConnectionStore()

	hReg := registry.New()
	pReg := registry.NewProcessor()

	em := event.NewManager()
	logger.Info("Started event broadcasting")
	sc := scheduler.NewCronScheduler()
	sc.Start()
	logger.Info("Started scheduler")

	return &MainServer{
		config:           cfg,
		logger:           logger,
		connStore:        connStore,
		scheduler:        sc,
		packetProcessors: pReg,
		packetHandlers:   hReg,
		eventManager:     em,
		database:         db,
	}
}
