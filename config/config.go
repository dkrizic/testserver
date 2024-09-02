package config

import (
	"flag"
	"github.com/dkrizic/testserver/telemetry/otelslog"
	"log/slog"
	"os"
	"strconv"
)

const (
	keyVerbose                          = "verbose"
	keyVerboseEnvironment               = "VERBOSE"
	keyPort                             = "port"
	keyPortEnvironment                  = "PORT"
	keyDatabaseHost                     = "database-host"
	keyDatabaseHostEnvironment          = "DATABASE_HOST"
	keyDatabasePort                     = "database-port"
	keyDatabasePortEnvironment          = "DATABASE_PORT"
	keyDatabaseUser                     = "database-user"
	keyDatabaseUserEnvironment          = "DATABASE_USER"
	keyDatabasePassword                 = "database-password"
	keyDatabasePasswordEnvironment      = "DATABASE_PASSWORD"
	keyDatabaseName                     = "database-name"
	keyDatabaseNameEnvironment          = "DATABASE_NAME"
	keyOpenTelemetryEndpoint            = "opentelemetry-endpoint"
	keyOpenTelemetryEndpointEnvironment = "OPENTELEMETRY_ENDPOINT"
	keyCheckToken                       = "check-token"
	keyCheckTokenEnvironment            = "CHECK_TOKEN"
)

type Database struct {
	host     string
	port     int
	user     string
	password string
	name     string
}

type Token struct {
	checkToken bool
}

type Service struct {
	port     int
	database Database
	token    Token
}

type Config struct {
	service               Service
	openTelemetryEndpoint string
}

func New() (*Config, error) {
	c := &Config{
		service: Service{},
	}
	verbose := flag.Int(keyVerbose, lookupEnvOrInt(keyVerboseEnvironment, 0), "Verbosity level, 0=info, 1=debug. Overrides the environment variable VERBOSE.")
	flag.IntVar(&c.service.port, keyPort, lookupEnvOrInt(keyPortEnvironment, 8000), "The port to listen on.")
	flag.StringVar(&c.service.database.host, keyDatabaseHost, lookupEnvOrString(keyDatabaseHostEnvironment, "localhost"), "The host of the database")
	flag.IntVar(&c.service.database.port, keyDatabasePort, lookupEnvOrInt(keyDatabasePortEnvironment, 3306), "The port of the database")
	flag.StringVar(&c.service.database.user, keyDatabaseUser, lookupEnvOrString(keyDatabaseUserEnvironment, "testserver"), "The user of the database")
	flag.StringVar(&c.service.database.password, keyDatabasePassword, lookupEnvOrString(keyDatabasePasswordEnvironment, "postgres"), "The password of the database")
	flag.StringVar(&c.service.database.name, keyDatabaseName, lookupEnvOrString(keyDatabaseNameEnvironment, "testserver"), "The name of the database")
	flag.StringVar(&c.openTelemetryEndpoint, keyOpenTelemetryEndpoint, lookupEnvOrString(keyOpenTelemetryEndpointEnvironment, ""), "The OpenTelemetry endpoint")
	flag.BoolVar(&c.service.token.checkToken, keyCheckToken, lookupEnvOrBool(keyCheckTokenEnvironment, true), "Check the token")

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
	return c.database.host
}

func (c Service) DatabasePort() int {
	return c.database.port
}

func (c Service) DatabaseUser() string {
	return c.database.user
}

func (c Service) DatabasePassword() string {
	return c.database.password
}

func (c Service) DatabaseName() string {
	return c.database.name
}

func (c Config) OpenTelemetryEndpoint() string {
	return c.openTelemetryEndpoint
}

func (c Config) CheckToken() bool {
	return c.service.token.checkToken
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
