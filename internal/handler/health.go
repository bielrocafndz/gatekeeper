package handler

import (
	"encoding/json"
	"net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	encoder.Encode(map[string]string{"status": "ok", "service": "gatekeeper"})
}
