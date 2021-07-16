package tests

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	graphql "github.com/machinebox/graphql"
	"github.com/nhost/stripe-graphql/tests/helpers"
	"github.com/nhost/stripe-graphql/utils/constants"
)

// Run go generate in /graph to load the methods that are being tested
func TestAddAndDeleteCustomer(t *testing.T) {
	err := godotenv.Load("testing.env")

	if err != nil {
		t.Fatal("Error loading .env file")
	}

	stripe_key := os.Getenv("STRIPE_KEY")

	go helpers.CreateServer()

	// Setup
	rand.Seed(time.Now().UnixNano())
	client := graphql.NewClient("http://localhost:8080/query")

	ctx := context.Background()

	name := fmt.Sprintf("Random_Customer %v", rand.Intn(1000))
	email := strings.ReplaceAll(name, " ", "") + "@gmail.com"
	description := "This is a random test description."

	t.Log("Creating new customer.")

	var m struct {
		InsertCustomer struct {
			ID    string
			Email string
			Name  string
		}
	}

	req := graphql.NewRequest(`
		mutation ($customer: CustomerInput) {
			insert_customer(input: $customer) {
			  id
			  name
			  balance
			  created
			}
	  }
	`)

	var customer = map[string]interface{}{
		"name":        name,
		"description": description,
		"email":       email,
	}

	req.Var("customer", customer)
	req.Header.Set(constants.STRIPE_KEY_HEADER, stripe_key)

	t.Log(req.Header)

	err = client.Run(ctx, req, &m)

	if err != nil {
		t.Fatal(err)
	}

	// new_id := &m.InsertCustomer.ID

	t.Logf("New customer created. ID: %v, Email: %v Name: %v", m.InsertCustomer.ID, m.InsertCustomer.Email, m.InsertCustomer.Name)

	// Deleting new Customer - Checks if customer was successfully added and provides cleanup
	// Also tests Delete Customer mutation
	// t.Cleanup(func() {
	// 	var d struct {
	// 		DeletedCustomer struct {
	// 			ID    string
	// 			Email string
	// 		} `graphql:"delete_customer(id: $id)"`
	// 	}

	// 	variables = map[string]interface{}{
	// 		"id": *new_id,
	// 	}

	// 	t.Logf("Deleting customer with ID: %v", m.InsertCustomer.ID)

	// 	err = client.Mutate(context.Background(), &d, variables)

	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// })
}
