package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nsltharaka/newsWaveAggregator/types"
	"github.com/nsltharaka/newsWaveAggregator/utils"
)

type contextKey string

const payloadKey = contextKey("payload")

var validate = validator.New()

func withValidation(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var payload types.RegisterUserPayload
		if err := utils.ParseJSON(r, &payload); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		if err := validate.Struct(payload); err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(
				"validation failed: either email or password contains invalid data",
			))
			return
		}

		ctx := context.WithValue(r.Context(), payloadKey, payload)
		next(w, r.WithContext(ctx))

	}

}
