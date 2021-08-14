package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nhost/stripe-graphql/graph"
	"github.com/nhost/stripe-graphql/graph/generated"
	"github.com/nhost/stripe-graphql/utils"
	"github.com/nhost/stripe-graphql/utils/constants"
)

// Run go generate in /graph before running the app to load all schemas, queries, and mutations
func CreateServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = constants.DEFAULT_PORT
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	http.Handle("/query", utils.StripeKeyMiddleware(srv))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	fmt.Printf("GraphQL server running on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
