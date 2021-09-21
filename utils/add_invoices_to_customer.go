package utils

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/nhost/stripe-graphql/utils/conversions"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

func AddInvoicesToCustomer(c *model.StripeCustomer, client *client.API) {
	if c == nil {
		return
	}

	params := stripe.InvoiceListParams{
		Customer: &c.ID,
	}

	invoices := client.Invoices.List(&params)
	var converted []*model.StripeInvoice

	for invoices.Next() {
		converted = append(converted, conversions.ConvertInvoice(invoices.Invoice()))
	}

	c.Invoices = converted
}
