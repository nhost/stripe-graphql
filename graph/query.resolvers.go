package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nhost/stripe-graphql/graph/generated"
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/nhost/stripe-graphql/utils"
	"github.com/nhost/stripe-graphql/utils/conversions"
	stripe "github.com/stripe/stripe-go/v72"
)

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	params := &stripe.CustomerListParams{}
	params.AddExpand("data.subscriptions")
	i := client.Customers.List(params)

	if i.Err() != nil {
		return nil, i.Err()
	}

	var new_customers []*model.Customer
	for i.Next() {
		new_customers = append(new_customers, conversions.ConvertCustomer(i.Customer()))
	}

	return new_customers, nil
}

func (r *queryResolver) Customer(ctx context.Context, id *string) (*model.Customer, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	params := &stripe.CustomerParams{}
	params.AddExpand("subscriptions")
	c, err := client.Customers.Get(*id, params)

	if err != nil {
		return nil, err
	}

	if c != nil {
		converted_customer := conversions.ConvertCustomer(c)
		return converted_customer, nil
	}
	return nil, nil
}

func (r *queryResolver) Invoices(ctx context.Context) ([]*model.Invoice, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	params := &stripe.InvoiceListParams{}
	params.AddExpand("data.customer")
	i := client.Invoices.List(params)

	var invoices []*model.Invoice

	for i.Next() {
		converted_invoice := conversions.ConvertInvoice(i.Invoice())
		invoices = append(invoices, converted_invoice)
	}

	return invoices, nil
}

func (r *queryResolver) Invoice(ctx context.Context, id string) (*model.Invoice, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	params := &stripe.InvoiceParams{}
	params.AddExpand("customer")
	i, err := client.Invoices.Get(id, nil)

	if err != nil {
		return nil, err
	}

	if i != nil {
		converted_invoice := conversions.ConvertInvoice(i)
		return converted_invoice, nil
	}
	return nil, nil
}

func (r *queryResolver) Prices(ctx context.Context) ([]*model.Price, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	params := &stripe.PriceListParams{}
	var prices []*model.Price
	i := client.Prices.List(params)

	for i.Next() {
		converted_price := conversions.ConvertPrice(i.Price())
		prices = append(prices, converted_price)
	}
	return prices, nil
}

func (r *queryResolver) Price(ctx context.Context, id string) (*model.Price, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	p, err := client.Prices.Get(id, nil)

	if err != nil {
		return nil, err
	}

	if p != nil {
		converted_price := conversions.ConvertPrice(p)
		return converted_price, nil
	}
	return nil, nil
}

func (r *queryResolver) PaymentMethods(ctx context.Context, customer string, typeArg model.PaymentMethodTypes) ([]*model.PaymentMethod, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	params := &stripe.PaymentMethodListParams{
		Customer: &customer,
		Type:     (*string)(&typeArg),
	}
	i := client.PaymentMethods.List(params)
	var payment_methods []*model.PaymentMethod
	for i.Next() {
		converted_object := conversions.ConvertPaymentMethod(i.PaymentMethod())
		payment_methods = append(payment_methods, converted_object)
	}
	return payment_methods, nil
}

func (r *queryResolver) PaymentMethod(ctx context.Context, id string) (*model.PaymentMethod, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	p, err := client.PaymentMethods.Get(id, nil)

	if err != nil {
		return nil, err
	}

	if p != nil {
		converted_pmethod := conversions.ConvertPaymentMethod(p)
		return converted_pmethod, nil
	}

	return nil, nil
}

func (r *queryResolver) Subscriptions(ctx context.Context) ([]*model.StripeSubscription, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	params := &stripe.SubscriptionListParams{}
	i := client.Subscriptions.List(params)

	var subscriptions []*model.StripeSubscription
	for i.Next() {
		converted_object := conversions.ConvertSubscription(i.Subscription())
		subscriptions = append(subscriptions, converted_object)
	}
	return subscriptions, nil
}

func (r *queryResolver) Subscription(ctx context.Context, id string) (*model.StripeSubscription, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	p, err := client.Subscriptions.Get(id, nil)

	if err != nil {
		return nil, err
	}

	if p != nil {
		converted_sub := conversions.ConvertSubscription(p)
		return converted_sub, nil
	}

	return nil, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	params := stripe.ProductListParams{}
	i := client.Products.List(&params)

	var products []*model.Product
	for i.Next() {
		converted_product := conversions.ConvertProduct(i.Product())
		products = append(products, converted_product)
	}

	return products, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	p, err := client.Products.Get(id, nil)

	if err != nil {
		return nil, err
	}

	if p != nil {
		converted_product := conversions.ConvertProduct(p)
		return converted_product, nil
	}

	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
