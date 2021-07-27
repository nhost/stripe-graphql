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

func (r *mutationResolver) InsertCustomer(ctx context.Context, input model.CustomerInput) (*model.Customer, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	params := &stripe.CustomerParams{
		Name:        input.Name,
		Description: input.Description,
		Email:       input.Email,
	}
	c, err := client.Customers.New(params)
	if err != nil {
		return nil, err
	}
	return conversions.ConvertCustomer(c), nil
}

func (r *mutationResolver) UpdateCustomer(ctx context.Context, id string, input model.CustomerInput) (*model.Customer, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	params := &stripe.CustomerParams{
		Name:        input.Name,
		Description: input.Description,
		Email:       input.Email,
	}
	c, err := client.Customers.Update(id, params)
	if err != nil {
		return nil, err
	}
	return conversions.ConvertCustomer(c), nil
}

func (r *mutationResolver) DeleteCustomer(ctx context.Context, id string) (*model.Customer, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	c, err := client.Customers.Del(id, nil)

	if err != nil {
		return nil, err
	}

	return conversions.ConvertCustomer(c), nil
}

func (r *mutationResolver) InsertSubscription(ctx context.Context, input model.CreateSubscriptionInput) (*model.StripeSubscription, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	var items []*stripe.SubscriptionItemsParams

	for _, item := range input.Items {
		stripe_params := conversions.ConvertToSubscriptionItemsParams(item)
		items = append(items, stripe_params)
	}

	params := &stripe.SubscriptionParams{
		Customer: &input.Customer,
		Items:    items,
	}

	s, err := client.Subscriptions.New(params)

	if err != nil {
		return nil, err
	}

	return conversions.ConvertSubscription(s), nil
}

func (r *mutationResolver) UpdateSubscription(ctx context.Context, id string, input model.UpdateSubscriptionInput) (*model.StripeSubscription, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	var items []*stripe.SubscriptionItemsParams

	for _, item := range input.Items {
		stripe_params := conversions.ConvertToSubscriptionItemsParams(item)
		items = append(items, stripe_params)
	}

	params := &stripe.SubscriptionParams{
		Customer: &input.Customer,
		Items:    items,
	}

	s, err := client.Subscriptions.Update(id, params)

	if err != nil {
		return nil, err
	}

	return conversions.ConvertSubscription(s), nil
}

func (r *mutationResolver) CancelSubscription(ctx context.Context, id string) (*model.StripeSubscription, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	s, err := client.Subscriptions.Cancel(id, nil)

	if err != nil {
		return nil, err
	}

	return conversions.ConvertSubscription(s), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
