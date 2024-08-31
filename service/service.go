package service

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dkrizic/testserver/database"
	"github.com/dkrizic/testserver/graph"
	"github.com/dkrizic/testserver/handler/health"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const defaultPort = "8000"

type Service struct {
	Port             string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
}

type Opts func(*Service)

func Port(port string) Opts {
	return func(s *Service) {
		s.Port = port
	}
}

func DatabaseHost(host string) Opts {
	return func(s *Service) {
		s.DatabaseHost = host
	}
}

func DatabasePort(port string) Opts {
	return func(s *Service) {
		s.DatabasePort = port
	}
}

func DatabaseUsername(username string) Opts {
	return func(s *Service) {
		s.DatabaseUsername = username
	}
}

func DatabasePassword(password string) Opts {
	return func(s *Service) {
		s.DatabasePassword = password
	}
}

func DatabaseName(name string) Opts {
	return func(s *Service) {
		s.DatabaseName = name
	}
}

func NewService(opts ...Opts) (*Service, error) {
	s := &Service{
		Port:         defaultPort,
		DatabaseHost: "localhost",
		DatabasePort: "3306",
	}
	for _, o := range opts {
		o(s)
	}
	return s, nil
}

func (s *Service) Run() error {
	slog.Info("Starting server")

	db, err := database.NewConnection(
		// database.Host("localhost"),
		// database.Port("3306"),
		database.Username("testserver"),
		database.Password("testserver"),
		database.Database("testserver"))
	if err != nil {
		slog.Error("Failed to connecto to database", "error", err)
		return err
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := graph.NewResolver(graph.DB(db))
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	slog.Info("Binding GraphQL playground to /")
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	slog.Info("Binding GraphQL query handler to /query")
	http.Handle("/query", srv)

	slog.Info("Binding health handler to /health")
	http.HandleFunc("/health", health.HealthHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	// start server and wait for interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)

	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			slog.Error("Failed to start server", "error", err)
			return
		}
	}()

	<-c
	slog.Info("Shutting down server")
	return nil
}
