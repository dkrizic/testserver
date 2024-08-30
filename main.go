package main

import (
	"github.com/dkrizic/testserver/database"
	"github.com/dkrizic/testserver/handler/health"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dkrizic/testserver/graph"
)

const defaultPort = "8080"

func main() {
	slog.Info("Starting server")
	
	db, err := database.NewConnection(
		// database.Host("localhost"),
		// database.Port("3306"),
		database.Username("testserver"),
		database.Password("testserver"),
		database.Database("testserver"))
	if err != nil {
		slog.Error("Failed to connecto to database", "error", err)
		return
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
}
