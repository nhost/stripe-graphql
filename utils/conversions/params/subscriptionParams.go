package params

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertToSubscriptionParams(input *model.CreateSubscriptionInput) *stripe.SubscriptionParams {

	var items []*stripe.SubscriptionItemsParams

	for _, item := range input.Items {
		stripe_params := ConvertToSubscriptionItemsParams(item)
		items = append(items, stripe_params)
	}

	params := &stripe.SubscriptionParams{
		Customer:             &input.Customer,
		Items:                items,
		CancelAtPeriodEnd:    input.CancelAtPeriodEnd,
		DefaultPaymentMethod: input.DefaultPaymentMethodID,
		CancelAt:             input.CancelAt,
		DaysUntilDue:         input.DaysUntilDue,
		TrialEnd:             input.TrialEnd,
	}

	return params
}

func ConvertToSubscriptionUpdateParams(input *model.UpdateSubscriptionInput) *stripe.SubscriptionParams {

	var items []*stripe.SubscriptionItemsParams

	for _, item := range input.Items {
		stripe_params := ConvertToSubscriptionItemsParams(item)
		items = append(items, stripe_params)
	}

	params := &stripe.SubscriptionParams{
		Items:                items,
		CancelAtPeriodEnd:    input.CancelAtPeriodEnd,
		DefaultPaymentMethod: input.DefaultPaymentMethodID,
		CancelAt:             input.CancelAt,
		DaysUntilDue:         input.DaysUntilDue,
		TrialEnd:             input.TrialEnd,
	}

	return params
}

func ConvertToSubscriptionItemsParams(item *model.SubscriptionItemInput) *stripe.SubscriptionItemsParams {
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

	return stripe_params
}
