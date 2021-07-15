package utils

import (
	"context"
	"net/http"

	"github.com/nhost/stripe-graphql/graph/utils/constants"
)

func StripeKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		header := r.Header[constants.STRIPE_KEY_HEADER]

		if len(header) == 0 {
			// not all graphql requests (eg. getting the schema) use the Stripe Key, so still
			// allowing requests to go through. Error handling done in the resolvers.
			next.ServeHTTP(rw, r)
			return
		}

		// Should be only one stripe key in the header, the first value
		key := header[0]

		// Exposing key to context so gqlgen resolvers(queries and mutations) can access it
		ctx := context.WithValue(r.Context(), constants.STRIPE_MAP_KEY, key)

		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
