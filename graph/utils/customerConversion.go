package utils

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertCustomer(old_customer *stripe.Customer) *model.Customer {
	currency := string(old_customer.Currency)
	created := int(old_customer.Created)
	address := ConvertAddress(&old_customer.Address)
	var subscriptions []*model.StripeSubscription
	if old_customer.Subscriptions != nil {
		for _, sub := range old_customer.Subscriptions.Data {
			subscriptions = append(subscriptions, ConvertSubscription(*sub))
		}
	}
	new_customer := &model.Customer{
		ID:           old_customer.ID,
		Name:         &old_customer.Name,
		Email:        &old_customer.Email,
		Address:      address,
		Phone:        &old_customer.Phone,
		Currency:     &currency,
		Created:      &created,
		Livemode:     &old_customer.Livemode,
		Subscription: subscriptions,
	}

	return new_customer
}

func ConvertAddress(old_address *stripe.Address) *model.Address {
	new_address := &model.Address{
		City:       &old_address.City,
		Country:    &old_address.Country,
		PostalCode: &old_address.PostalCode,
		State:      &old_address.State,
		Line1:      &old_address.Line1,
		Line2:      &old_address.Line2,
	}

	return new_address
}
