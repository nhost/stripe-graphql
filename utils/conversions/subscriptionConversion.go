package conversions

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertSubscription(old_subscription *stripe.Subscription) *model.StripeSubscription {
	var items []*model.SubscriptionItem
	if old_subscription.Items != nil {
		for _, item := range old_subscription.Items.Data {
			items = append(items, ConvertSubscriptionItem(item))
		}
	}

	var customer *model.Customer = nil
	if old_subscription.Customer != nil {
		customer = ConvertCustomer(old_subscription.Customer)
	}

	var default_pay_method *model.PaymentMethod = nil
	if old_subscription.DefaultPaymentMethod != nil {
		default_pay_method = ConvertPaymentMethod(old_subscription.DefaultPaymentMethod)
	}

	return &model.StripeSubscription{
		ID:                   old_subscription.ID,
		CancelAtPeriodEnd:    old_subscription.CancelAtPeriodEnd,
		CancelAt:             old_subscription.CancelAt,
		CurrentPeriodStart:   old_subscription.CurrentPeriodStart,
		CurrentPeriodEnd:     old_subscription.CurrentPeriodEnd,
		Customer:             customer,
		Items:                items,
		DefaultPaymentMethod: default_pay_method,
		Status:               string(old_subscription.Status),
		Created:              old_subscription.Created,
		EndedAt:              old_subscription.EndedAt,
		CanceledAt:           old_subscription.CanceledAt,
		DaysUntilDue:         old_subscription.DaysUntilDue,
		Livemode:             old_subscription.Livemode,
		TrialStart:           old_subscription.TrialStart,
		TrialEnd:             old_subscription.TrialEnd,
	}
}

func ConvertSubscriptionItem(old *stripe.SubscriptionItem) *model.SubscriptionItem {
	var p *model.Price = nil
	if old.Price != nil {
		p = ConvertPrice(old.Price)
	}

	return &model.SubscriptionItem{
		ID:       old.ID,
		Quantity: old.Quantity,
		Price:    p,
	}
}
