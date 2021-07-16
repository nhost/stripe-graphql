package constants

type ContextKey string

const (
	STRIPE_KEY_HEADER            = "X-Stripe-Secret-Key"
	STRIPE_MAP_KEY    ContextKey = "STRIPE_KEY"
)
