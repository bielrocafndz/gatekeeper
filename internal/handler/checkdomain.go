package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

func CheckDomain(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")

	client := http.Client{Timeout: (5 * time.Second)}

	resp, err := client.Get("https://" + domain)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	status := "online"

	if err != nil {
		status = "offline"
	}else {
		defer resp.Body.Close()
	}

	encoder.Encode(map[string]string{"status": status, "service": domain})

}
