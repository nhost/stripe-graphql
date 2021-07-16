package constants

import "fmt"

type ContextKey string

const (
	STRIPE_KEY_HEADER            = "X-Stripe-Secret-Key"
	STRIPE_CLIENT_KEY ContextKey = "STRIPE_KEY"
)

const (
	STRIPE_KEY_MISSING_ERROR = "no stripe api key provided"
	INTERNAL_ERROR           = "internal error, please try again later"
)

var STRIPE_CLIENT_CONTEXT_INVALID string = fmt.Sprintf("key %v in context had invalid value, needs stripe *client.API", STRIPE_CLIENT_KEY)

const DEFAULT_PORT = "8080"
