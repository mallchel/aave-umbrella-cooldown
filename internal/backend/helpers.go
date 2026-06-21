package backend

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(body)
}

func badParam(message string) error {
	return &ParamError{message: message}
}

type ParamError struct {
	message string
}

func (e *ParamError) Error() string {
	return e.message
}
