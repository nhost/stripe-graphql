package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nhost/stripe-graphql/graph/generated"
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/nhost/stripe-graphql/utils/constants"
	"github.com/nhost/stripe-graphql/utils/conversions"
	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/invoice"
	"github.com/stripe/stripe-go/v72/paymentmethod"
	"github.com/stripe/stripe-go/v72/price"
	"github.com/stripe/stripe-go/v72/product"
	"github.com/stripe/stripe-go/v72/sub"
)

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	fmt.Printf("Stripe Key from context: %v\n", ctx.Value(constants.STRIPE_MAP_KEY))
	params := &stripe.CustomerListParams{}
	params.AddExpand("data.subscriptions")
	i := customer.List(params)

	var new_customers []*model.Customer
	for i.Next() {
		new_customers = append(new_customers, conversions.ConvertCustomer(i.Customer()))
	}

	return new_customers, nil
}

func (r *queryResolver) Customer(ctx context.Context, id *string) (*model.Customer, error) {
	c, err := customer.Get(*id, nil)

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
	params := &stripe.InvoiceListParams{}
	i := invoice.List(params)

	var invoices []*model.Invoice

	for i.Next() {
		converted_invoice := conversions.ConvertInvoice(i.Invoice())
		invoices = append(invoices, converted_invoice)
	}

	return invoices, nil
}

func (r *queryResolver) Invoice(ctx context.Context, id string) (*model.Invoice, error) {
	i, err := invoice.Get(id, nil)

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
	params := &stripe.PriceListParams{}
	var prices []*model.Price
	i := price.List(params)
	for i.Next() {
		converted_price := conversions.ConvertPrice(*i.Price())
		prices = append(prices, converted_price)
	}
	return prices, nil
}

func (r *queryResolver) Price(ctx context.Context, id string) (*model.Price, error) {
	p, err := price.Get(id, nil)

	if err != nil {
		return nil, err
	}

	if p != nil {
		converted_price := conversions.ConvertPrice(*p)
		return converted_price, nil
	}
	return nil, nil
}

func (r *queryResolver) PaymentMethods(ctx context.Context, customer string, typeArg *model.PaymentMethodTypes) ([]*model.PaymentMethod, error) {
	params := &stripe.PaymentMethodListParams{
		Customer: &customer,
		Type:     (*string)(typeArg),
	}
	i := paymentmethod.List(params)
	var payment_methods []*model.PaymentMethod
	for i.Next() {
		converted_object := conversions.ConvertPaymentMethod(i.PaymentMethod())
		payment_methods = append(payment_methods, converted_object)
	}
	return payment_methods, nil
}

func (r *queryResolver) PaymentMethod(ctx context.Context, id string) (*model.PaymentMethod, error) {
	p, err := paymentmethod.Get(id, nil)

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
	params := &stripe.SubscriptionListParams{}
	i := sub.List(params)

	var subscriptions []*model.StripeSubscription
	for i.Next() {
		converted_object := conversions.ConvertSubscription(*i.Subscription())
		subscriptions = append(subscriptions, converted_object)
	}
	return subscriptions, nil
}

func (r *queryResolver) Subscription(ctx context.Context, id string) (*model.StripeSubscription, error) {
	p, err := sub.Get(id, nil)

	if err != nil {
		return nil, err
	}

	if p != nil {
		converted_sub := conversions.ConvertSubscription(*p)
		return converted_sub, nil
	}

	return nil, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	params := stripe.ProductListParams{}
	i := product.List(&params)

	var products []*model.Product
	for i.Next() {
		converted_product := conversions.ConvertProduct(*i.Product())
		products = append(products, converted_product)
	}

	return products, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	p, err := product.Get(id, nil)

	if err != nil {
		return nil, err
	}

	if p != nil {
		converted_product := conversions.ConvertProduct(*p)
		return converted_product, nil
	}

	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
