package conversions

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

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
