package handlers

import (
	"encoding/json"
	"net/http"
)

// Represents an error that can be returned in an API response.
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Writes an error response in JSON format.
func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	apiError := APIError{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(apiError); err != nil {
		http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
	}
}
