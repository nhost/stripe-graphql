package utils

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertInvoice(old_invoice *stripe.Invoice) *model.Invoice {
	currency := string(old_invoice.Currency)
	created := int(old_invoice.Created)
	new_invoice := &model.Invoice{
		ID:       &old_invoice.ID,
		Currency: &currency,
		Created:  &created,
		Livemode: &old_invoice.Livemode,
	}

	return new_invoice
}
