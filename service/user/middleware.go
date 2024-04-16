package user

import (
	"context"
	"net/http"

	"github.com/nsltharaka/newsWaveAggregator/types"
	"github.com/nsltharaka/newsWaveAggregator/utils"
)

type contextKey string

const payloadKey = contextKey("payload")

func withValidation(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var payload types.RegisterUserPayload
		if err := utils.ParseJSON(r, &payload); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.WithValue(r.Context(), payloadKey, payload)
		next(w, r.WithContext(ctx))

	}

}
