package conversions

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertProduct(old_product *stripe.Product) *model.Product {
	new := &model.Product{
		ID:          old_product.ID,
		Description: old_product.Description,
		Created:     old_product.Created,
		Livemode:    old_product.Livemode,
		Active:      old_product.Active,
		Name:        old_product.Name,
	}
	return new
}
