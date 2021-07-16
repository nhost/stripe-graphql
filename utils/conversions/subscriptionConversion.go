package conversions

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertSubscription(old_subscription stripe.Subscription) *model.StripeSubscription {
	current_period_start := (int)(old_subscription.CurrentPeriodStart)
	current_period_end := (int)(old_subscription.CurrentPeriodEnd)
	var items []*model.SubscriptionItem
	if old_subscription.Items != nil {
		for _, item := range old_subscription.Items.Data {
			items = append(items, ConvertSubscriptionItem(*item))
		}
	}
	return &model.StripeSubscription{
		ID:                 &old_subscription.ID,
		CancelAtPeriodEnd:  &old_subscription.CancelAtPeriodEnd,
		CurrentPeriodStart: &current_period_start,
		CurrentPeriodEnd:   &current_period_end,
		Items:              items,
	}
}

func ConvertSubscriptionItem(old stripe.SubscriptionItem) *model.SubscriptionItem {
	created := (int)(old.Created)
	quantity := (int)(old.Quantity)
	return &model.SubscriptionItem{
		ID:       &old.ID,
		Object:   &old.Object,
		Created:  &created,
		Quantity: &quantity,
	}
}
