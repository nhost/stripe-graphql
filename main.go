package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nhost/stripe-graphql/graph"
	"github.com/nhost/stripe-graphql/graph/generated"
	"github.com/nhost/stripe-graphql/utils"
)

const defaultPort = "8080"

// Run go generate in /graph before running the app to load all schemas, queries, and mutations
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/query", utils.StripeKeyMiddleware(srv))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	fmt.Printf("GraphQL server running on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
