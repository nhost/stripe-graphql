package conversions

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertPrice(old_price *stripe.Price) *model.StripePrice {
	currency := string(old_price.Currency)
	billing_scheme := string(old_price.BillingScheme)

	var p *model.StripeProduct = nil
	if old_price.Product != nil {
		p = ConvertProduct(old_price.Product)
	}

	var r *model.StripePriceRecurring = nil
	if old_price.Recurring != nil {
		r = ConvertRecurring(old_price.Recurring)
	}

	new_price := &model.StripePrice{
		ID:            old_price.ID,
		Nickname:      old_price.Nickname,
		Livemode:      old_price.Livemode,
		Currency:      currency,
		Created:       old_price.Created,
		Active:        old_price.Active,
		BillingScheme: billing_scheme,
		Product:       p,
		Recurring:     r,
	}

	return new_price
}

func ConvertRecurring(old *stripe.PriceRecurring) *model.StripePriceRecurring {

	new := &model.StripePriceRecurring{
		IntervalCount:  old.IntervalCount,
		Interval:       (model.RecurringInterval)(old.Interval),
		AggregateUsage: (model.PriceRecurringAggregateUsage)(old.AggregateUsage),
		UsageType:      (model.PriceRecurringUsageType)(old.UsageType),
	}

	return new
}
