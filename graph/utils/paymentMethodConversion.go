package utils

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertPaymentMethod(old *stripe.PaymentMethod) *model.PaymentMethod {

	new := &model.PaymentMethod{
		ID: &old.ID,
	}

	return new
}
