package subgraphs

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/OldBigBuddha/gqlgen/graphql"
	"github.com/OldBigBuddha/gqlgen/graphql/handler"
	"github.com/OldBigBuddha/gqlgen/graphql/handler/debug"
	"github.com/OldBigBuddha/gqlgen/graphql/playground"
	"golang.org/x/sync/errgroup"
)

type Config struct {
	EnableDebug bool
}

type Subgraphs struct {
	servers []*http.Server
}

type SubgraphConfig struct {
	Name   string
	Schema graphql.ExecutableSchema
	Port   string
}

func (s *Subgraphs) Shutdown(ctx context.Context) error {
	for _, srv := range s.servers {
		if err := srv.Shutdown(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *Subgraphs) ListenAndServe(ctx context.Context) error {
	group, _ := errgroup.WithContext(ctx)
	for _, srv := range s.servers {
		srv := srv
		group.Go(func() error {
			err := srv.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Printf("error listening and serving: %v", err)
				return err
			}
			return nil
		})
	}
	return group.Wait()
}

func newServer(name, port string, schema graphql.ExecutableSchema) *http.Server {
	if port == "" {
		panic(fmt.Errorf("port for %s is empty", name))
	}
	srv := handler.NewDefaultServer(schema)
	srv.Use(&debug.Tracer{})
	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)
	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}

func New(ctx context.Context, subgraphs ...SubgraphConfig) (*Subgraphs, error) {
	servers := make([]*http.Server, len(subgraphs))
	for i, config := range subgraphs {
		servers[i] = newServer(config.Name, config.Port, config.Schema)
	}

	return &Subgraphs{
		servers: servers,
	}, nil
}
