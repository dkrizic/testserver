package service

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dkrizic/testserver/database"
	"github.com/dkrizic/testserver/graph"
	"github.com/dkrizic/testserver/service/handler/health"
	log2 "github.com/dkrizic/testserver/service/handler/log"
	"github.com/dkrizic/testserver/service/handler/token"
	"github.com/dkrizic/testserver/service/version"
	"github.com/ravilushqa/otelgqlgen"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultPort = "8000"
	defaultPath = "/"
	versionPath = "/version"
	healthPath  = "/health"
	queryPath   = "/graph"
)

type Service struct {
	Port             string
	CheckToken       bool
	DatabaseHost     string
	DatabasePort     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
	CleanDatabase    bool
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

func CleanDatabase(clean bool) Opts {
	return func(s *Service) {
		s.CleanDatabase = clean
	}
}

func CheckToken(checkToken bool) Opts {
	return func(s *Service) {
		s.CheckToken = checkToken
	}
}

func NewService(opts ...Opts) (*Service, error) {
	s := &Service{
		Port:         defaultPort,
		DatabaseHost: "localhost",
		DatabasePort: "3306",
		CheckToken:   true,
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

	if s.CleanDatabase {
		err = database.CleanSchema(db)
		if err != nil {
			slog.Error("Failed to clean schema", "error", err)
			return err
		}
	}

	err = database.MigrateSchema(db)
	if err != nil {
		slog.Error("Failed to migrate schema", "error", err)
		return err
	}

	mux := http.NewServeMux()

	slog.Info("Binding health handler", "path", healthPath)
	mux.HandleFunc(healthPath, health.HealthHandler)

	slog.Info("Binding version handler", "path", versionPath)
	mux.HandleFunc(versionPath, version.VersionHandler)

	resolver := graph.NewResolver(graph.DB(db))
	queryHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	slog.Info("Binding query handler", "path", queryPath)
	queryHandler.Use(otelgqlgen.Middleware(otelgqlgen.WithSpanKindSelector(func(operationName string) trace.SpanKind {
		switch operationName {
		case "Query/group":
			return trace.SpanKindClient
		case "Query/user":
			return trace.SpanKindClient
		case "Group/users":
			return trace.SpanKindClient
		case "User/groups":
			return trace.SpanKindClient
		case "Query/tag":
			return trace.SpanKindClient
		case "Query/tagValue":
			return trace.SpanKindClient
		case "Tag/assets":
			return trace.SpanKindClient
		case "TagValue/tag":
			return trace.SpanKindClient
		case "TagValue/asset":
			return trace.SpanKindClient
		case "Query/identity":
			return trace.SpanKindClient
		}
		return trace.SpanKindInternal
	})))
	mux.Handle(queryPath, queryHandler)

	slog.Info("Binding playground", "path", defaultPath)
	mux.Handle("/", playground.Handler("GraphQL playground", queryPath))

	logHandler := log2.LogHandler(mux)

	tokenHandler := token.NewTokenHandler(
		token.WithCheckToken(s.CheckToken),
		token.WithAlwaysPass(func(r *http.Request) bool {
			return r.URL.Path == healthPath || r.URL.Path == versionPath || r.URL.Path == defaultPath
		}),
	).TokenHandler(logHandler)

	server := &http.Server{
		Addr:    s.Port,
		Handler: tokenHandler,
	}

	// start server and wait for interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			slog.Error("Failed to start server", "error", err)
			return
		}
	}()

	<-c
	slog.Info("Shutting down server")
	return nil
}
