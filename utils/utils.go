package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
