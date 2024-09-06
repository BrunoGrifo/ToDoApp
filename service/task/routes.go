package task

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /task", h.handleGetTasks)
}

func (h *Handler) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Word")
}
