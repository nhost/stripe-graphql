package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nhost/stripe-graphql/graph"
	"github.com/nhost/stripe-graphql/graph/generated"
	"github.com/nhost/stripe-graphql/utils"
	"github.com/stripe/stripe-go/v72"
)

const defaultPort = "8080"

// Run go generate in /graph before running the app to load all schemas, queries, and mutations
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	stripe.Key = os.Getenv("STRIPE_KEY")
	if stripe.Key == "" {
		log.Fatal("No STRIPE_KEY env variable provided.")
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", utils.StripeKeyMiddleware(srv))
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
