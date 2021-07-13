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
	"github.com/stripe/stripe-go/v72/sub"
)

func (r *mutationResolver) InsertCustomer(ctx context.Context, input *model.CustomerInput) (*model.Customer, error) {
	params := &stripe.CustomerParams{
		Name:        input.Name,
		Description: input.Description,
		Email:       input.Email,
	}
	c, err := customer.New(params)
	if err != nil {
		return nil, err
	}
	return utils.ConvertCustomer(c), nil
}

func (r *mutationResolver) UpdateCustomer(ctx context.Context, id string, input *model.CustomerInput) (*model.Customer, error) {
	params := &stripe.CustomerParams{
		Name:        input.Name,
		Description: input.Description,
		Email:       input.Email,
	}
	c, err := customer.Update(id, params)
	if err != nil {
		return nil, err
	}
	return utils.ConvertCustomer(c), nil
}

func (r *mutationResolver) DeleteCustomer(ctx context.Context, id string) (*model.Customer, error) {
	c, err := customer.Del(id, nil)

	if err != nil {
		return nil, err
	}

	return utils.ConvertCustomer(c), nil
}

func (r *mutationResolver) InsertSubscription(ctx context.Context, input *model.SubscriptionInput) (*model.StripeSubscription, error) {
	var items []*stripe.SubscriptionItemsParams

	for _, item := range input.Items {
		var price_data_params *stripe.SubscriptionItemPriceDataParams

		if item.PriceData != nil {
			interval_count := int64(*item.PriceData.Recurring.IntervalCount)
			var recurring_params *stripe.SubscriptionItemPriceDataRecurringParams
			if item.PriceData.Recurring != nil {
				recurring_params = &stripe.SubscriptionItemPriceDataRecurringParams{
					Interval:      (*string)(&item.PriceData.Recurring.Interval),
					IntervalCount: &interval_count,
				}
			}
			price_data_params = &stripe.SubscriptionItemPriceDataParams{
				Currency:  (*string)(&item.PriceData.Currency),
				Product:   &item.PriceData.Product,
				Recurring: recurring_params,
			}
		}

		stripe_params := &stripe.SubscriptionItemsParams{
			Price:     item.Price,
			PriceData: price_data_params,
		}
		items = append(items, stripe_params)
	}

	params := &stripe.SubscriptionParams{
		Customer: &input.Customer,
		Items:    items,
	}

	s, err := sub.New(params)

	if err != nil {
		return nil, err
	}

	return utils.ConvertSubscription(*s), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
