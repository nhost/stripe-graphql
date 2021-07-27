package conversions

import (
	"github.com/nhost/stripe-graphql/graph/model"
	"github.com/stripe/stripe-go/v72"
)

func ConvertPaymentMethod(old *stripe.PaymentMethod) *model.PaymentMethod {

	var c *model.Customer = nil
	if old.Customer != nil {
		c = ConvertCustomer(old.Customer)
	}

	var card *model.Card = nil
	if old.Card != nil {
		card = ConvertCard(old.Card)
	}

	new := &model.PaymentMethod{
		ID:       old.ID,
		Customer: c,
		Created:  int(old.Created),
		Object:   old.Object,
		Livemode: old.Livemode,
		Type:     string(old.Type),
		Card:     card,
	}

	return new
}

func ConvertCard(old_card *stripe.PaymentMethodCard) *model.Card {
	new := &model.Card{
		Brand:    string(old_card.Brand),
		Country:  old_card.Country,
		ExpMonth: int(old_card.ExpMonth),
		ExpYear:  int(old_card.ExpYear),
		Funding:  model.FundingType(old_card.Funding),
		Last4:    old_card.Last4,
	}

	return new
}
