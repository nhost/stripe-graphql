package tests

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	graphql "github.com/hasura/go-graphql-client"
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/nhost/stripe-graphql/tests/helpers"
)

// Run go generate in /graph to load the methods that are being tested
func TestAddAndDeleteCustomer(t *testing.T) {
	go helpers.CreateServer()

	// Setup
	rand.Seed(time.Now().UnixNano())
	client := graphql.NewClient("http://localhost:8080/query", nil)
	name := fmt.Sprintf("Random_Customer %v", rand.Intn(1000))
	email := strings.ReplaceAll(name, " ", "") + "@gmail.com"
	description := "This is a random test description."

	t.Log("Creating new customer.")

	var m struct {
		InsertCustomer struct {
			ID    string
			Email string
			Name  string
		} `graphql:"insert_customer(input: $customer)"`
	}

	variables := map[string]interface{}{
		"customer": model.CustomerInput{
			Name:        &name,
			Email:       &email,
			Description: &description,
		},
	}

	err := client.Mutate(context.Background(), &m, variables)

	if err != nil {
		t.Error(err)
	}

	new_id := &m.InsertCustomer.ID

	t.Logf("New customer created. ID: %v, Email: %v Name: %v", m.InsertCustomer.ID, m.InsertCustomer.Email, m.InsertCustomer.Name)

	// Deleting new Customer - Checks if customer was successfully added and provides cleanup
	// Also tests Delete Customer mutation
	t.Cleanup(func() {
		var d struct {
			DeletedCustomer struct {
				ID    string
				Email string
			} `graphql:"delete_customer(id: $id)"`
		}

		variables = map[string]interface{}{
			"id": *new_id,
		}

		t.Logf("Deleting customer with ID: %v", m.InsertCustomer.ID)

		err = client.Mutate(context.Background(), &d, variables)

		if err != nil {
			t.Error(err.Error())
		}
	})
}
