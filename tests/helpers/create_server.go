package helpers

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/nhost/stripe-graphql/graph"
	"github.com/nhost/stripe-graphql/graph/generated"
	"github.com/stripe/stripe-go/v72"
)

const defaultPort = "8080"

func CreateServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	stripe.Key = os.Getenv("STRIPE_KEY")
	if stripe.Key == "" {
		log.Fatal("No STRIPE_KEY env variable provided for testing.")
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/query", srv)
	http.ListenAndServe(":"+port, nil)
}
