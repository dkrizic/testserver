package database

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/pressly/goose/v3"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
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

func NewConnection(opts ...Opts) (db *sql.DB, err error) {
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
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		slog.Warn("Failed to connect to database", "error", err)
		return nil, err
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		slog.Warn("Failed to ping database", "error", err)
		return
	}

	slog.Info("Connected to database", "host", c.host, "port", c.port, "username", c.username, "database", c.database)
	return db, nil
}

func CleanSchema(db *sql.DB) error {
	slog.Info("Cleaning schema")
	goose.SetBaseFS(embedMigrations)

	err := goose.SetDialect("mysql")
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = goose.DownToContext(ctx, db, "schema", 0)
	if err != nil {
		return err
	}

	return nil
}

func MigrateSchema(db *sql.DB) error {
	slog.Info("Migrating schema")
	goose.SetBaseFS(embedMigrations)

	err := goose.SetDialect("mysql")
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = goose.UpContext(ctx, db, "schema")
	if err != nil {
		return err
	}

	return nil
}
