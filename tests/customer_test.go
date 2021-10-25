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
	"github.com/nhost/stripe-graphql/server"
	"github.com/nhost/stripe-graphql/utils/constants"
)

const addErrorPrefix = "Error adding stripe customer: "
const deleteErrorPrefix = "Error deleting stripe customer: "

// Run go generate in /graph to load the methods that are being tested
func TestAddAndDeleteCustomer(t *testing.T) {
	go server.CreateServer()

	err := godotenv.Load("testing.env")

	if err != nil {
		t.Fatal("Error loading testing.env file")
	}

	stripe_key := os.Getenv("STRIPE_KEY")

	// Setup
	rand.Seed(time.Now().UnixNano())
	client := graphql.NewClient("http://localhost:8080/query")

	ctx := context.Background()

	name := fmt.Sprintf("Random_Customer %v", rand.Intn(1000))
	email := strings.ReplaceAll(name, " ", "") + "@gmail.com"
	description := "This is a random test description."

	t.Log("Creating new customer.")

	// Variable for New customer
	var m map[string]interface{}

	req := graphql.NewRequest(`
		mutation ($customer: StripeCustomerInput!) {
			createCustomer(input: $customer) {
			  id
			  name
			  email
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

	err = client.Run(ctx, req, &m)

	if err != nil {
		t.Fatalf("%v%v", addErrorPrefix, err)
	}

	new_customer, ok := m["createCustomer"].(map[string]interface{})

	new_id := fmt.Sprint(new_customer["id"])

	if ok == false {
		t.Fatalf("%vincorrect graphql response", addErrorPrefix)
	}

	t.Logf("New customer created. ID: %v, Email: %v Name: %v", new_customer["id"], new_customer["email"], new_customer["name"])

	// Deleting new Customer - Checks if customer was successfully added and provides cleanup
	// Also tests Delete Customer mutation
	t.Cleanup(func() {

		req := graphql.NewRequest(`
		mutation ($id: String!) {
			deleteCustomer(id: $id) {
			  id
			  email
			}
	  	}
		`)

		req.Var("id", new_id)
		req.Header.Set(constants.STRIPE_KEY_HEADER, stripe_key)

		t.Logf("Deleting customer with ID: %v", new_id)

		err = client.Run(context.Background(), req, nil)

		if err != nil {
			t.Fatalf("%v%v", deleteErrorPrefix, err)
		}
	})
}
