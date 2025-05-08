package http_errors

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

func SendJSONError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Message: message,
	})
}
