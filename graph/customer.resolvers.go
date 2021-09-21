package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/nhost/stripe-graphql/utils"
	"github.com/nhost/stripe-graphql/utils/conversions"
	stripe "github.com/stripe/stripe-go/v72"
)

func (r *customerResolver) PaymentMethods(ctx context.Context, obj *model.StripeCustomer, typeArg *model.PaymentMethodTypes) ([]*model.StripePaymentMethod, error) {
	client, err := utils.GetClientFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if typeArg == nil {
		*typeArg = "card"
	}

	params := stripe.PaymentMethodListParams{
		Customer: &obj.ID,
		Type:     (*string)(typeArg),
	}

	pms := client.PaymentMethods.List(&params)
	err = pms.Err()
	if err != nil {
		return nil, err
	}

	var converted []*model.StripePaymentMethod

	for pms.Next() {
		converted = append(converted, conversions.ConvertPaymentMethod(pms.PaymentMethod()))
	}

	return converted, nil
}

type customerResolver struct{ *Resolver }
