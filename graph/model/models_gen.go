// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Card struct {
	Brand             *string            `json:"brand"`
	Country           *string            `json:"country"`
	ExpMonth          *int               `json:"exp_month"`
	ExpYear           *int               `json:"exp_year"`
	Fingerprint       *string            `json:"fingerprint"`
	Funding           *string            `json:"funding"`
	GeneratedFrom     *string            `json:"generated_from"`
	Last4             *string            `json:"last4"`
	Wallet            *string            `json:"wallet"`
	ThreeDSecureUsage *ThreeDSecureUsage `json:"three_d_secure_usage"`
	Networks          *Networks          `json:"networks"`
	Checks            *Checks            `json:"checks"`
}

type Checks struct {
	AddressLine1Check      *string `json:"address_line1_check"`
	AddressPostalCodeCheck *string `json:"address_postal_code_check"`
	CvcCheck               *string `json:"cvc_check"`
}

type Customer struct {
	ID       string  `json:"id"`
	Address  *string `json:"address"`
	Created  *int    `json:"created"`
	Currency *string `json:"currency"`
	Email    *string `json:"email"`
	Livemode *bool   `json:"livemode"`
	Name     *string `json:"name"`
	Phone    *string `json:"phone"`
}

type Invoice struct {
	ID               *string `json:"id"`
	Created          *int    `json:"created"`
	Currency         *string `json:"currency"`
	Customer         *string `json:"customer"`
	HostedInvoiceURL *string `json:"hosted_invoice_url"`
	Livemode         *bool   `json:"livemode"`
	Paid             *bool   `json:"paid"`
	PeriodEnd        *int    `json:"period_end"`
	PeriodStart      *int    `json:"period_start"`
	Status           *string `json:"status"`
	Subtotal         *int    `json:"subtotal"`
	Tax              *string `json:"tax"`
	Total            *int    `json:"total"`
	Lines            *Lines  `json:"lines"`
}

type InvoiceLine struct {
	ID          *string `json:"id"`
	Object      *string `json:"object"`
	Amount      *int    `json:"amount"`
	Currency    *string `json:"currency"`
	Description *string `json:"description"`
	Livemode    *bool   `json:"livemode"`
	Quantity    *int    `json:"quantity"`
	Type        *string `json:"type"`
}

type Lines struct {
	Object *string        `json:"object"`
	Data   []*InvoiceLine `json:"data"`
}

type Metadata struct {
	OrderID *string `json:"order_id"`
}

type Networks struct {
	Preferred *string   `json:"preferred"`
	Available []*string `json:"available"`
}

type PaymentMethod struct {
	ID       *string   `json:"id"`
	Created  *int      `json:"created"`
	Customer *string   `json:"customer"`
	Livemode *bool     `json:"livemode"`
	Type     *string   `json:"type"`
	Metadata *Metadata `json:"metadata"`
	Card     *Card     `json:"card"`
}

type Price struct {
	ID            *string    `json:"id"`
	Active        *bool      `json:"active"`
	BillingScheme *string    `json:"billing_scheme"`
	Created       *int       `json:"created"`
	Currency      *string    `json:"currency"`
	Livemode      *bool      `json:"livemode"`
	Nickname      *string    `json:"nickname"`
	Product       *string    `json:"product"`
	Recurring     *Recurring `json:"recurring"`
}

type Product struct {
	ID          *string `json:"id"`
	Active      *bool   `json:"active"`
	Created     *int    `json:"created"`
	Description *string `json:"description"`
	Livemode    *bool   `json:"livemode"`
	Name        *string `json:"name"`
}

type Recurring struct {
	AggregateUsage *string `json:"aggregate_usage"`
	Interval       *string `json:"interval"`
	IntervalCount  *int    `json:"interval_count"`
	UsageType      *string `json:"usage_type"`
}

type StripeSubscription struct {
	ID                 *string `json:"id"`
	CancelAtPeriodEnd  *bool   `json:"cancel_at_period_end"`
	CurrentPeriodEnd   *int    `json:"current_period_end"`
	CurrentPeriodStart *int    `json:"current_period_start"`
	Status             *string `json:"status"`
}

type ThreeDSecureUsage struct {
	Supported *bool `json:"supported"`
}

type PaymentMethodTypes string

const (
	PaymentMethodTypesAcssDebit        PaymentMethodTypes = "acss_debit"
	PaymentMethodTypesAfterpayClearpay PaymentMethodTypes = "afterpay_clearpay"
	PaymentMethodTypesAlipay           PaymentMethodTypes = "alipay"
	PaymentMethodTypesAuBecsDebit      PaymentMethodTypes = "au_becs_debit"
	PaymentMethodTypesBacsDebit        PaymentMethodTypes = "bacs_debit"
	PaymentMethodTypesBancontact       PaymentMethodTypes = "bancontact"
	PaymentMethodTypesBoleto           PaymentMethodTypes = "boleto"
	PaymentMethodTypesCard             PaymentMethodTypes = "card"
	PaymentMethodTypesEps              PaymentMethodTypes = "eps"
	PaymentMethodTypesFpx              PaymentMethodTypes = "fpx"
	PaymentMethodTypesGiropay          PaymentMethodTypes = "giropay"
	PaymentMethodTypesGrabpay          PaymentMethodTypes = "grabpay"
	PaymentMethodTypesIdeal            PaymentMethodTypes = "ideal"
	PaymentMethodTypesOxxo             PaymentMethodTypes = "oxxo"
	PaymentMethodTypesP24              PaymentMethodTypes = "p24"
	PaymentMethodTypesSepaDebit        PaymentMethodTypes = "sepa_debit"
	PaymentMethodTypesSofort           PaymentMethodTypes = "sofort"
	PaymentMethodTypesWechatPay        PaymentMethodTypes = "wechat_pay"
)

var AllPaymentMethodTypes = []PaymentMethodTypes{
	PaymentMethodTypesAcssDebit,
	PaymentMethodTypesAfterpayClearpay,
	PaymentMethodTypesAlipay,
	PaymentMethodTypesAuBecsDebit,
	PaymentMethodTypesBacsDebit,
	PaymentMethodTypesBancontact,
	PaymentMethodTypesBoleto,
	PaymentMethodTypesCard,
	PaymentMethodTypesEps,
	PaymentMethodTypesFpx,
	PaymentMethodTypesGiropay,
	PaymentMethodTypesGrabpay,
	PaymentMethodTypesIdeal,
	PaymentMethodTypesOxxo,
	PaymentMethodTypesP24,
	PaymentMethodTypesSepaDebit,
	PaymentMethodTypesSofort,
	PaymentMethodTypesWechatPay,
}

func (e PaymentMethodTypes) IsValid() bool {
	switch e {
	case PaymentMethodTypesAcssDebit, PaymentMethodTypesAfterpayClearpay, PaymentMethodTypesAlipay, PaymentMethodTypesAuBecsDebit, PaymentMethodTypesBacsDebit, PaymentMethodTypesBancontact, PaymentMethodTypesBoleto, PaymentMethodTypesCard, PaymentMethodTypesEps, PaymentMethodTypesFpx, PaymentMethodTypesGiropay, PaymentMethodTypesGrabpay, PaymentMethodTypesIdeal, PaymentMethodTypesOxxo, PaymentMethodTypesP24, PaymentMethodTypesSepaDebit, PaymentMethodTypesSofort, PaymentMethodTypesWechatPay:
		return true
	}
	return false
}

func (e PaymentMethodTypes) String() string {
	return string(e)
}

func (e *PaymentMethodTypes) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PaymentMethodTypes(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid payment_method_types", str)
	}
	return nil
}

func (e PaymentMethodTypes) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
