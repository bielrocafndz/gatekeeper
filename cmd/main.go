package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bielrocafndz/gatekeeper/internal/handler"
)

func main() {
	http.HandleFunc("/health", handler.GetHealth)
	http.HandleFunc("/check", handler.CheckDomain)
	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server is closed\n")
	}
}
