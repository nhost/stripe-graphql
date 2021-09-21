package utils

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/nhost/stripe-graphql/utils/conversions"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

func AddPriceToProduct(p *model.StripeProduct, client *client.API) {
	if p == nil {
		return
	}

	params := stripe.PriceListParams{
		Product: &p.ID,
	}

	prices := client.Prices.List(&params)
	var converted_prices []*model.StripePrice
	for prices.Next() {
		converted_prices = append(converted_prices, conversions.ConvertPrice(prices.Price()))
	}

	p.Prices = converted_prices
}
