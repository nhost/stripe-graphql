package main

import (
	"github.com/nhost/stripe-graphql/server"
)

// Run go generate in /graph before running the app to load all schemas, queries, and mutations
func main() {
	server.CreateServer()
}
