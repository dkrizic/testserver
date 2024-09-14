package service

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/dkrizic/testserver/config"
	"github.com/dkrizic/testserver/database"
	"github.com/dkrizic/testserver/graph"
	"log/slog"
	"os"
	"testing"
)

func TestGraphQl(t *testing.T) {
	os.Setenv("VERBOSE", "4")
	config, err := config.New()
	if err != nil {
		t.Fatalf("failed to load configuration: %v", err)
	}

	db, err := database.NewConnection(
		database.Host(config.Service().DatabaseHost()),
		database.Port("3306"),
		database.Username("testserver"),
		database.Password("testserver"),
		database.Database("testserver"),
	)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	resolver := graph.NewResolver(graph.DB(db))
	c := client.New(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver})))

	t.Run("Query", func(t *testing.T) {
		var resp struct {
			Group []struct {
				ID   string
				Name string
			}
		}
		c.MustPost(`{group{id,name}}`, &resp)
		slog.Info("Response", "resp", resp)
	})
}
