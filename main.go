package main

import (
	"fmt"
	"github.com/dkrizic/testserver/config"
	"github.com/dkrizic/testserver/meta"
	"github.com/dkrizic/testserver/service"
	"log/slog"
)

func main() {

	c, err := config.New()
	if err != nil {
		slog.Error("Failed to create config", "error", err)
		return
	}

	slog.Info("Starting service", "service", meta.ServiceName, "version", meta.Version)
	slog.Debug("Configuration",
		"port", c.Service().Port(),
		"database-host", c.Service().DatabaseHost(),
		"database-port", c.Service().DatabasePort(),
		"database-user", c.Service().DatabaseUser(),
		"database-name", c.Service().DatabaseName())

	s, err := service.NewService(
		service.Port(fmt.Sprintf(":%d", c.Service().Port())),
		service.DatabaseHost(c.Service().DatabaseHost()),
		service.DatabasePort(fmt.Sprintf("%d", c.Service().DatabasePort())),
		service.DatabaseUsername(c.Service().DatabaseUser()),
		service.DatabasePassword(c.Service().DatabasePassword()),
		service.DatabaseName(c.Service().DatabaseHost()))
	if err != nil {
		slog.Error("Failed to create service", "error", err)
		return
	}

	err = s.Run()
	if err != nil {
		slog.Error("Failed to run service", "error", err)
		return
	}

	slog.Info("Stopping service", "service", meta.ServiceName)
}
