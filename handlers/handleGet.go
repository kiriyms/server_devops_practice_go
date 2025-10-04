package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/kiriyms/server_devops_practice_go/common"
)

type JSONResponse struct {
	Msg          string `json:"msg"`
	VisitorCount int    `json:"visitorCount"`
}

type VisitorHandler struct {
	count int
	mu    sync.Mutex
}

func NewVisitorHandler() *VisitorHandler {
	return &VisitorHandler{}
}

func (h *VisitorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	h.count++
	visitorCount := h.count

	w.Header().Set("Content-Type", "application/json")

	userId := common.GetUserId()
	response := JSONResponse{
		Msg:          fmt.Sprintf("Hello, user %s!", userId),
		VisitorCount: visitorCount,
	}

	json.NewEncoder(w).Encode(response)
}
