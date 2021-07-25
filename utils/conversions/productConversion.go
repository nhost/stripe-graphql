package conversions

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertProduct(old_product *stripe.Product) *model.Product {
	created := (int)(old_product.Created)
	new := &model.Product{
		ID:          &old_product.ID,
		Description: &old_product.Description,
		Created:     &created,
	}
	return new
}
