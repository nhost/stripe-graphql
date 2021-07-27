package conversions

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertInvoice(old_invoice *stripe.Invoice) *model.Invoice {

	currencyStr := (string(old_invoice.Currency))
	currency := (model.CurrencyTypes)(currencyStr)
	created := int(old_invoice.Created)

	period_end := (int)(old_invoice.PeriodEnd)
	period_start := (int)(old_invoice.PeriodStart)
	tax := int(old_invoice.Tax)
	total := int(old_invoice.Total)

	var c *model.Customer = nil
	if old_invoice.Customer != nil {
		c = ConvertCustomer(old_invoice.Customer)
	}

	var line_object *model.Lines = nil

	if len(old_invoice.Lines.Data) != 0 {
		var lines []*model.InvoiceLine
		for _, line := range old_invoice.Lines.Data {
			lines = append(lines, ConvertInvoiceLine(line))
		}

		object := "list"
		line_object = &model.Lines{
			Object: &object,
			Data:   lines,
		}
	}

	new_invoice := &model.Invoice{
		ID:               &old_invoice.ID,
		Currency:         &currency,
		Created:          &created,
		Customer:         c,
		Livemode:         &old_invoice.Livemode,
		Lines:            line_object,
		HostedInvoiceURL: &old_invoice.HostedInvoiceURL,
		Paid:             &old_invoice.Paid,
		PeriodEnd:        &period_end,
		PeriodStart:      &period_start,
		Status:           (*string)(&old_invoice.Status),
		Tax:              &tax,
		Total:            &total,
	}

	return new_invoice
}

func ConvertInvoiceLine(old_line *stripe.InvoiceLine) *model.InvoiceLine {
	currency := string(old_line.Currency)
	amount := int(old_line.Amount)
	quantity := int(old_line.Quantity)
	new := &model.InvoiceLine{
		ID:          &old_line.ID,
		Amount:      &amount,
		Currency:    &currency,
		Object:      &old_line.Object,
		Description: &old_line.Description,
		Livemode:    &old_line.Livemode,
		Quantity:    &quantity,
		Type:        (*string)(&old_line.Type),
	}

	return new
}
