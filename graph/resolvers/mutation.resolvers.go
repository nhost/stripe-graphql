package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nhost/stripe-graphql/graph/generated"
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/nhost/stripe-graphql/utils"
	"github.com/nhost/stripe-graphql/utils/conversions"
	"github.com/nhost/stripe-graphql/utils/conversions/params"
	stripe "github.com/stripe/stripe-go/v72"
)

func (r *mutationResolver) StripeCreateCustomer(ctx context.Context, input model.StripeCustomerInput) (*model.StripeCustomer, error) {
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

func (r *mutationResolver) StripeUpdateCustomer(ctx context.Context, id string, input model.StripeCustomerInput) (*model.StripeCustomer, error) {
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

func (r *mutationResolver) StripeDeleteCustomer(ctx context.Context, id string) (*model.StripeCustomer, error) {
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

func (r *mutationResolver) StripeCreateSubscription(ctx context.Context, input model.StripeCreateSubscriptionInput) (*model.StripeSubscription, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	s, err := client.Subscriptions.New(params.ConvertToSubscriptionParams(&input))

	if err != nil {
		return nil, err
	}

	return conversions.ConvertSubscription(s), nil
}

func (r *mutationResolver) StripeUpdateSubscription(ctx context.Context, id string, input model.StripeUpdateSubscriptionInput) (*model.StripeSubscription, error) {
	client, err := utils.GetClientFromContext(ctx)

	if err != nil {
		return nil, err
	}

	s, err := client.Subscriptions.Update(id, params.ConvertToSubscriptionUpdateParams(&input))

	if err != nil {
		return nil, err
	}

	return conversions.ConvertSubscription(s), nil
}

func (r *mutationResolver) StripeCancelSubscription(ctx context.Context, id string) (*model.StripeSubscription, error) {
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
