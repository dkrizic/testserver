package main

import (
	"context"
	"fmt"
	"github.com/dkrizic/testserver/config"
	"github.com/dkrizic/testserver/meta"
	"github.com/dkrizic/testserver/service"
	"github.com/dkrizic/testserver/telemetry"
	"log/slog"
	"os"
)

func main() {

	ctx := context.Background()

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
		"database-name", c.Service().DatabaseName(),
		"opentelemetry-endpoint", c.OpenTelemetryEndpoint())

	var shutdown func(context.Context) error
	if c.OpenTelemetryEndpoint() != "" {
		slog.InfoContext(ctx, "OpenTelemetry enabled", "endpoint", c.OpenTelemetryEndpoint())
		shutdown, err = telemetry.OpenTelemetryConfig{
			OTLPEndpoint:   c.OpenTelemetryEndpoint(),
			ServiceName:    meta.ServiceName,
			ServiceVersion: meta.Version,
		}.InitOpenTelemetry(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "Failed to initialize OpenTelemetry", "error", err)
			os.Exit(1)
		}
	} else {
		slog.InfoContext(ctx, "OpenTelemetry disabled")
	}

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

	if shutdown != nil {
		slog.Info("Shutting down OpenTelemetry")
		err := shutdown(ctx)
		if err != nil {
			slog.Error("Failed to shutdown OpenTelemetry", "error", err)
		}
	}

	slog.Info("Stopping service", "service", meta.ServiceName)
}
