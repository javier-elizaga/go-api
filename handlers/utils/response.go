package utils

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, statusCode int, error string) {
	w.WriteHeader(statusCode)
	errorRes := struct {
		Error string `json:"error"`
	}{Error: error}
	json.NewEncoder(w).Encode(errorRes)
}
