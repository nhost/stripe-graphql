package utils

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertSubscription(old_subscription stripe.Subscription) *model.StripeSubscription {
	current_period_start := (int)(old_subscription.CurrentPeriodStart)
	current_period_end := (int)(old_subscription.CurrentPeriodEnd)
	return &model.StripeSubscription{
		ID:                 &old_subscription.ID,
		CancelAtPeriodEnd:  &old_subscription.CancelAtPeriodEnd,
		CurrentPeriodStart: &current_period_start,
		CurrentPeriodEnd:   &current_period_end,
	}
}
