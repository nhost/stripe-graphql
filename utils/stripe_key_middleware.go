package utils

import (
	"context"
	"net/http"

	"github.com/nhost/stripe-graphql/utils/constants"
)

// Gets stripe key from http header, initializes client, and puts it in request context
// So query and mutation resolvers can easily access client
func StripeKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// not all graphql requests (eg. getting the schema) use the Stripe Key, so still
		// allowing requests to go through. Error handling done in the resolvers.
		defer next.ServeHTTP(rw, r)

		header := r.Header[constants.STRIPE_KEY_HEADER]
		if len(header) == 0 {
			return
		}

		// Should be only one stripe key in the header, the first value
		key := header[0]

		client := GetStripeClient(key)

		// Exposing client to context so gqlgen resolvers(queries and mutations) can access it
		ctx := context.WithValue(r.Context(), constants.STRIPE_CLIENT_KEY, client)

		*r = *r.WithContext(ctx)
	})
}
