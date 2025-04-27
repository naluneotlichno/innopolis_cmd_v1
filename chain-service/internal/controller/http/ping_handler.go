package http

import (
	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if _, err := fmt.Fprintln(w, "pong"); err != nil {
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
		return
	}

}
