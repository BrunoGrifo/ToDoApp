package task

import (
	"log"
	"net/http"
	"todo/types"

	"github.com/google/uuid"
)

type Handler struct {
	repository types.TaskRepository
}

func NewHandler(repository types.TaskRepository) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /task", h.handleGetTasks)
	router.HandleFunc("POST /task", h.handleCreateTasks)
}

func (h *Handler) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		// get all
		log.Println("hello")
	} else {
		// get by id
		// Validate UUID
		uuid, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, "Invalid UUID", http.StatusBadRequest)
			return
		}
		h.repository.GetTaskById(uuid)
		log.Println(uuid)
	}
}

func (h *Handler) handleCreateTasks(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
}
