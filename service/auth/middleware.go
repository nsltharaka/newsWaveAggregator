package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/utils"
)

type ContextKey string

func WithAuthUser(db *database.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// extract the api key
			authorizationHeader := r.Header.Get("Authorization")
			if authorizationHeader == "" {
				utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf(
					"missing authorization headers ",
				))
				return
			}

			authHeaderSplit := strings.Split(authorizationHeader, " ")
			if authHeaderSplit[0] != "ApiKey" {
				utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf(
					"malformed authorization header",
				))
				return
			}
			apiKey := authHeaderSplit[1]

			// verify user
			user, err := db.GetUserByApiKey(r.Context(), apiKey)
			if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf(
					"user authentication failed",
				))
				return
			}

			// set user in the context so the handler can access user data
			ctxWithUser := context.WithValue(r.Context(), ContextKey("authUser"), user.ID)

			next.ServeHTTP(w, r.WithContext(ctxWithUser))

		})
	}
}
