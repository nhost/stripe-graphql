package utils

import (
	"context"
	"errors"
	"log"

	"github.com/nhost/stripe-graphql/utils/constants"
	"github.com/stripe/stripe-go/v72/client"
)

func GetStripeClient(key string) *client.API {
	stripe_client := &client.API{}
	stripe_client.Init(key, nil)

	return stripe_client
}

func GetClientFromContext(ctx context.Context) (*client.API, error) {
	stripe_client_interface := ctx.Value(constants.STRIPE_CLIENT_KEY)

	if stripe_client_interface == nil {
		return nil, errors.New(constants.STRIPE_KEY_MISSING_ERROR)
	}

	client, ok := stripe_client_interface.(*client.API)

	if !ok {
		log.Default().Println(constants.STRIPE_CLIENT_CONTEXT_INVALID)

		/*
			Returns internal error because it is passed on by the request, so don't
			want to reveal whats actually going on internally
		*/
		return nil, errors.New(constants.INTERNAL_ERROR)
	}

	return client, nil
}
