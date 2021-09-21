package conversions

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertInvoice(old_invoice *stripe.Invoice) *model.StripeInvoice {

	currencyStr := (string(old_invoice.Currency))
	currency := (model.CurrencyTypes)(currencyStr)

	var c *model.StripeCustomer = nil
	if old_invoice.Customer != nil {
		c = ConvertCustomer(old_invoice.Customer)
	}

	var line_object *model.StripeLines = nil

	if len(old_invoice.Lines.Data) != 0 {
		var lines []*model.StripeInvoiceLine
		for _, line := range old_invoice.Lines.Data {
			lines = append(lines, ConvertInvoiceLine(line))
		}

		object := "list"
		line_object = &model.StripeLines{
			Object: &object,
			Data:   lines,
		}
	}

	new_invoice := &model.StripeInvoice{
		ID:               old_invoice.ID,
		Currency:         currency,
		Created:          old_invoice.Created,
		Customer:         c,
		Livemode:         old_invoice.Livemode,
		Lines:            line_object,
		HostedInvoiceURL: old_invoice.HostedInvoiceURL,
		Paid:             old_invoice.Paid,
		PeriodEnd:        old_invoice.PeriodEnd,
		PeriodStart:      old_invoice.PeriodStart,
		Status:           (string)(old_invoice.Status),
		Tax:              old_invoice.Tax,
		Total:            old_invoice.Total,
	}

	return new_invoice
}

func ConvertInvoiceLine(old_line *stripe.InvoiceLine) *model.StripeInvoiceLine {
	currency := string(old_line.Currency)
	new := &model.StripeInvoiceLine{
		ID:          old_line.ID,
		Amount:      old_line.Amount,
		Currency:    currency,
		Object:      old_line.Object,
		Description: old_line.Description,
		Livemode:    old_line.Livemode,
		Quantity:    old_line.Quantity,
		Type:        (string)(old_line.Type),
	}

	return new
}
