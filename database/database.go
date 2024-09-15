package database

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"github.com/pressly/goose/v3"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

//go:embed schema/*.sql
var embedMigrations embed.FS

type config struct {
	host     string
	port     string
	username string
	password string
	database string
	clean    bool
}

type Opts func(*config)

func Host(host string) Opts {
	return func(c *config) {
		c.host = host
	}
}

func Port(port string) Opts {
	return func(c *config) {
		c.port = port
	}
}

func Username(username string) Opts {
	return func(c *config) {
		c.username = username
	}
}

func Password(password string) Opts {
	return func(c *config) {
		c.password = password
	}
}

func Database(database string) Opts {
	return func(c *config) {
		c.database = database
	}
}

func NewConnection(ctx context.Context, opts ...Opts) (db *sql.DB, err error) {
	c := &config{
		host:     "localhost",
		port:     "3306",
		username: "root",
		password: "",
		database: "testserver",
		clean:    false,
	}

	for _, opt := range opts {
		opt(c)
	}

	slog.Debug("Connecting to database", "host", c.host, "port", c.port, "username", c.username, "database", c.database)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.username, c.password, c.host, c.port, c.database)
	db, err = otelsql.Open("mysql", connectionString,
		otelsql.WithDBName(c.database),
		otelsql.WithAttributes(semconv.DBSystemMySQL))
	if err != nil {
		slog.Warn("Failed to connect to database", "error", err)
		return nil, err
	}

	otelsql.ReportDBStatsMetrics(db, otelsql.WithAttributes(
		semconv.DBSystemMySQL))
	if err != nil {
		slog.Warn("Failed to report database stats", "error", err)
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		slog.Warn("Failed to ping database", "error", err)
		return
	}

	slog.Info("Connected to database", "host", c.host, "port", c.port, "username", c.username, "database", c.database)
	return db, nil
}

func CleanSchema(ctx context.Context, db *sql.DB) error {
	slog.Info("Cleaning schema")
	goose.SetBaseFS(embedMigrations)

	err := goose.SetDialect("mysql")
	if err != nil {
		return err
	}

	err = goose.DownContext(ctx, db, "schema")
	if err != nil {
		return err
	}

	return nil
}

func MigrateSchema(ctx context.Context, db *sql.DB) error {
	slog.Info("Migrating schema")
	goose.SetBaseFS(embedMigrations)

	err := goose.SetDialect("mysql")
	if err != nil {
		return err
	}

	err = goose.UpContext(ctx, db, "schema")
	if err != nil {
		return err
	}

	return nil
}
