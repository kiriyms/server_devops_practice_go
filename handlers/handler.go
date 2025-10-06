package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/kiriyms/server_devops_practice_go/services"
)

type JSONResponse struct {
	Msg          string `json:"msg"`
	VisitorCount int    `json:"visitorCount"`
}

type Handler struct {
	svc   services.Service
	count int
	mu    sync.Mutex
}

func NewHandler(service services.Service) *Handler {
	return &Handler{
		svc: service,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	h.count++
	visitorCount := h.count

	w.Header().Set("Content-Type", "application/json")

	msg, err := h.svc.Greet(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err), http.StatusInternalServerError)
		return
	}

	response := JSONResponse{
		Msg:          msg,
		VisitorCount: visitorCount,
	}

	json.NewEncoder(w).Encode(response)
}
