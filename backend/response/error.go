package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

func RespondError(w http.ResponseWriter, status int, err error, message string) {
	log.Printf("[ERROR] %d: %s (%s)", status, message, err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   err.Error(),
		Code:    status,
		Message: message,
	})
}
