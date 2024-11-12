package main

import (
	"log"
	"net/http"
	"os"

	"github.com/OldBigBuddha/gqlgen/graphql/handler"
	"github.com/OldBigBuddha/gqlgen/graphql/handler/debug"
	"github.com/OldBigBuddha/gqlgen/graphql/playground"
	"github.com/OldBigBuddha/gqlgen/plugin/federation/testdata/computedrequires"
	"github.com/OldBigBuddha/gqlgen/plugin/federation/testdata/computedrequires/generated"
)

const defaultPort = "4003"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &computedrequires.Resolver{}}))
	srv.Use(&debug.Tracer{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
