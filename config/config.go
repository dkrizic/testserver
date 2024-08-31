package config

import (
	"flag"
	"github.com/dkrizic/testserver/telemetry/otelslog"
	"log/slog"
	"os"
	"strconv"
)

const (
	keyVerbose                     = "verbose"
	keyVerboseEnvironment          = "VERBOSE"
	keyPort                        = "port"
	keyPortEnvironment             = "PORT"
	keyDatabaseHost                = "database-host"
	keyDatabaseHostEnvironment     = "DATABASE_HOST"
	keyDatabasePort                = "database-port"
	keyDatabasePortEnvironment     = "DATABASE_PORT"
	keyDatabaseUser                = "database-user"
	keyDatabaseUserEnvironment     = "DATABASE_USER"
	keyDatabasePassword            = "database-password"
	keyDatabasePasswordEnvironment = "DATABASE_PASSWORD"
	keyDatabaseName                = "database-name"
	keyDatabaseNameEnvironment     = "DATABASE_NAME"
)

type Service struct {
	port             int
	databaseHost     string
	databasePort     int
	databaseUser     string
	databasePassword string
	databaseName     string
}

type Config struct {
	service Service
}

func New() (*Config, error) {
	c := &Config{
		service: Service{},
	}
	verbose := flag.Int(keyVerbose, lookupEnvOrInt(keyVerboseEnvironment, 0), "Verbosity level, 0=info, 1=debug. Overrides the environment variable VERBOSE.")
	flag.IntVar(&c.service.port, keyPort, lookupEnvOrInt(keyPortEnvironment, 8000), "The port to listen on.")
	flag.StringVar(&c.service.databaseHost, keyDatabaseHost, lookupEnvOrString(keyDatabaseHostEnvironment, "localhost"), "The host of the database")
	flag.IntVar(&c.service.databasePort, keyDatabasePort, lookupEnvOrInt(keyDatabasePortEnvironment, 5432), "The port of the database")
	flag.StringVar(&c.service.databaseUser, keyDatabaseUser, lookupEnvOrString(keyDatabaseUserEnvironment, "postgres"), "The user of the database")
	flag.StringVar(&c.service.databasePassword, keyDatabasePassword, lookupEnvOrString(keyDatabasePasswordEnvironment, "postgres"), "The password of the database")
	flag.StringVar(&c.service.databaseName, keyDatabaseName, lookupEnvOrString(keyDatabaseNameEnvironment, "postgres"), "The name of the database")

	flag.Parse()

	level := slog.LevelInfo
	switch *verbose {
	case 0:
		level = slog.LevelError
	case 1:
		level = slog.LevelWarn
	case 2:
		level = slog.LevelInfo
	case 3:
		level = slog.LevelDebug
	}
	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level})
	otelhandler := otelslog.NewHandler(handler)
	slog.SetDefault(slog.New(otelhandler))

	return c, nil
}

func (c Config) Service() Service {
	return c.service
}

func (c Service) Port() int {
	return c.port
}

func (c Service) DatabaseHost() string {
	return c.databaseHost
}

func (c Service) DatabasePort() int {
	return c.databasePort
}

func (c Service) DatabaseUser() string {
	return c.databaseUser
}

func (c Service) DatabasePassword() string {
	return c.databasePassword
}

func (c Service) DatabaseName() string {
	return c.databaseName
}

func lookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func lookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			slog.Error("LookupEnvOrInt", "key", key, "error", err)
		}
		return v
	}
	return defaultVal
}

func lookupEnvOrBool(key string, defaultVal bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseBool(val)
		if err != nil {
			slog.Error("LookupEnvOrBool", "key", key, "error", err)
		}
		return v
	}
	return defaultVal
}
