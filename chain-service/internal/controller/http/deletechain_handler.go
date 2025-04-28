package http

import (
	"net/http"
	"path"
)

func (h *MessageChainHandler) DeleteMessageChain(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	uuid := path.Base(r.URL.Path)
	err := h.service.DeleteMessageChain(uuid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
