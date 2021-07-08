package utils

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertInvoice(old_invoice *stripe.Invoice) *model.Invoice {
	currency := string(old_invoice.Currency)
	created := int(old_invoice.Created)
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
		ID:       &old_invoice.ID,
		Currency: &currency,
		Created:  &created,
		Livemode: &old_invoice.Livemode,
		Lines:    line_object,
	}

	return new_invoice
}

func ConvertInvoiceLine(old_line *stripe.InvoiceLine) *model.InvoiceLine {
	currency := string(old_line.Currency)
	amount := int(old_line.Amount)
	new := &model.InvoiceLine{
		ID:       &old_line.ID,
		Amount:   &amount,
		Currency: &currency,
	}

	return new
}
