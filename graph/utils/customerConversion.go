package utils

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertCustomer(old_customer *stripe.Customer) *model.Customer {
	currency := string(old_customer.Currency)
	created := int(old_customer.Created)
	var subscriptions []*model.StripeSubscription
	for _, sub := range old_customer.Subscriptions.Data {
		subscriptions = append(subscriptions, ConvertSubscription(*sub))
	}
	new_customer := &model.Customer{
		ID:           old_customer.ID,
		Name:         &old_customer.Name,
		Email:        &old_customer.Email,
		Address:      &old_customer.Address.Line1,
		Phone:        &old_customer.Phone,
		Currency:     &currency,
		Created:      &created,
		Livemode:     &old_customer.Livemode,
		Subscription: subscriptions,
	}

	return new_customer
}
