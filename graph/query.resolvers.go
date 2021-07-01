package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nhost/stripe-graphql/graph/generated"
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/nhost/stripe-graphql/graph/utils"
	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/invoice"
)

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	params := &stripe.CustomerListParams{}
	i := customer.List(params)

	var new_customers []*model.Customer
	for i.Next() {
		new_customers = append(new_customers, utils.ConvertCustomer(i.Customer()))
	}

	return new_customers, nil
}

func (r *queryResolver) Customer(ctx context.Context, id *string) (*model.Customer, error) {
	c, _ := customer.Get(*id, nil)
	if c != nil {
		converted_customer := utils.ConvertCustomer(c)
		return converted_customer, nil
	}
	return nil, nil
}

func (r *queryResolver) Invoices(ctx context.Context) ([]*model.Invoice, error) {
	params := &stripe.InvoiceListParams{}
	i := invoice.List(params)

	var invoices []*model.Invoice

	for i.Next() {
		converted_invoice := utils.ConvertInvoice(i.Invoice())
		invoices = append(invoices, converted_invoice)
	}

	return invoices, nil
}

func (r *queryResolver) Invoice(ctx context.Context, id string) (*model.Invoice, error) {
	i, _ := invoice.Get(id, nil)
	if i != nil {
		converted_invoice := utils.ConvertInvoice(i)
		return converted_invoice, nil
	}
	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
