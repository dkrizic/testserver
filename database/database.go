package database

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	host     string
	port     string
	username string
	password string
	database string
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

	err = db.Ping()
	if err != nil {
		slog.Warn("Failed to ping database", "error", err)
		return
	}

	slog.Info("Connected to database", "host", c.host, "port", c.port, "username", c.username, "database", c.database)
	return db, nil
}
