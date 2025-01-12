package main

import (
	"go.uber.org/zap"
	"net"
	"os"
	"os/signal"
	"pixels-emulator/core/server"
	"pixels-emulator/core/setup"
	"pixels-emulator/core/setup/ephemeral"
	"strconv"
	"syscall"
)

// main initializes the server, binds it, and handles system signals for graceful shutdown.
func main() {

	sv := server.GetServer()
	ephemeral.Processors()
	ephemeral.Handlers()
	ephemeral.Cron()
	ephemeral.Event()

	// Bind the server and handle any errors
	if err := bindServer(sv); err != nil {
		sv.Logger.Error("Error while binding HTTP server", zap.Error(err))
		os.Exit(1)
	}

	// Channel to listen for system termination signals
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

	// Wait for signal to stop the server
	sig := <-sigChannel
	sv.Logger.Info("Signal to stop server received", zap.String("signal", sig.String()))

	// Stop the server and handle any errors
	if err := sv.Stop(); err != nil {
		sv.Logger.Error("Error while stopping server", zap.Error(err))
	}
	sv.Logger.Info("Server stopped gracefully")
}

// bindServer configures the HTTP server and starts listening on the specified IP and port.
func bindServer(sv *server.Server) error {
	app, err := setup.Router(zap.L(), sv.PacketProcessors, sv.PacketHandlers, sv.ConnStore)
	if err != nil || app == nil {
		return err
	}

	bind := net.JoinHostPort(sv.Config.Server.IP, strconv.Itoa(int(sv.Config.Server.Port)))
	sv.Logger.Info("Starting HTTP server", zap.String("bind", bind))
	return app.Listen(bind)
}
