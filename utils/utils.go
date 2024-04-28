package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nsltharaka/newsWaveAggregator/types"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {

	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)

}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, errorStatus int, err error) error {
	return WriteJSON(w, errorStatus, map[string]string{
		"error": err.Error(),
	})
}

func ValidateInput[T types.CanValidated](
	w http.ResponseWriter,
	r *http.Request,
	payload *T,

) (*T, error) {

	if err := ParseJSON(r, payload); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return nil, err
	}

	// validation
	if err := Validate.Struct(payload); err != nil {
		WriteError(w, http.StatusBadRequest, fmt.Errorf(
			"validation - payload validation failed, please refer to the api documentation",
		))
		return nil, err
	}

	return payload, nil

}
