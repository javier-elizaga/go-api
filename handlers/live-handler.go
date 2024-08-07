package handlers

import (
	"encoding/json"
	"net/http"
)

func GetLive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Ok bool `json:"ok"`
	}{Ok: true}
	json.NewEncoder(w).Encode(response)
}
