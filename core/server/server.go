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

var (
	once     sync.Once
	instance *Server
)

// GetServer provides the only server instance.
func GetServer() *Server {
	once.Do(func() {
		instance = setupServer()
	})
	return instance
}

// Server defines a single instance of server which can be accessed to get the
// global state of the Pixels Emulator.
type Server struct {
	Config           *config.Config              // Config provides the loaded configuration from file.
	Logger           *zap.Logger                 // Logger provides the server logger.
	ConnStore        *protocol.ConnectionStore   // ConnStore provides all the active connections on the server.
	Scheduler        *scheduler.Scheduler        // Schedule provides the scheduler instance.
	PacketProcessors *registry.ProcessorRegistry // PacketProcessors provides all the packets available to be processed.
	PacketHandlers   *registry.Registry          // PacketHandlers provides the handlers of the packets processed.
	EventManager     *event.Manager              // EventManager provides the event management system global instance.
	Database         *gorm.DB                    // Database provides a connection for ORM.
}

// Reload performs the reloading of processors, handlers, cron jobs, and events.
// It logs the duration of the reload process and alerts if any issues occur.
func (s *Server) Reload(log bool) {
	startTime := time.Now()

	// Perform the reload steps

	duration := time.Since(startTime)

	// Check if there was any alert or error during the reload process
	if log {
		s.Logger.Info("Reload completed", zap.Duration("duration", duration))

		if duration > time.Second*5 {
			// Log a warning if reload takes longer than expected
			s.Logger.Warn("Reload took longer than expected", zap.Duration("duration", duration))
		} else {
			// Otherwise, log the successful reload duration
			s.Logger.Info("Reload finished within expected time", zap.Duration("duration", duration))
		}
	}
}

func (s *Server) Stop() error {
	s.ConnStore.CloseActive()
	return nil
}

func setupServer() *Server {

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
	pReg := registry.NewProcessorRegistry()

	em := event.NewManager()
	logger.Info("Started event broadcasting")
	sc := scheduler.NewCronScheduler()
	sc.Start()
	logger.Info("Started scheduler")

	return &Server{
		Config:           cfg,
		Logger:           logger,
		ConnStore:        connStore,
		Scheduler:        &sc,
		PacketProcessors: pReg,
		PacketHandlers:   hReg,
		EventManager:     em,
		Database:         db,
	}

}
