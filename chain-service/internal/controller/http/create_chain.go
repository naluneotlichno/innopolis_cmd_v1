package http

import (
	"encoding/json"
	"net/http"
)

func (h *MessageChainHandler) CreateMessageChain(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	chain, err := h.service.CreateMessageChain(req.UserID, req.Title)
	if err != nil {
		http.Error(w, "Failed to create message chain", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(chain); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}

var req struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
}
