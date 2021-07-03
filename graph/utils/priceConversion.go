package utils

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertPrice(old_price stripe.Price) *model.Price {
	currency := string(old_price.Currency)
	billing_scheme := string(old_price.BillingScheme)
	new_price := &model.Price{
		ID:            &old_price.ID,
		Nickname:      &old_price.Nickname,
		Livemode:      &old_price.Livemode,
		Currency:      &currency,
		Active:        &old_price.Active,
		BillingScheme: &billing_scheme,
	}

	return new_price
}
