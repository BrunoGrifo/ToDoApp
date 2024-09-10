package task

import (
	"html/template"
	"log"
	"net/http"
	"todo/types"
	"todo/utils"

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
	router.HandleFunc("PUT /task", h.handleUpdateTasks)
	router.HandleFunc("DELETE /task", h.handleDeleteTasks)
}

func (h *Handler) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		// get all
		log.Println("Getting all tasks...")
		tasks, err := h.repository.GetAllTasks()
		if err != nil {
			http.Error(w, "Internel server error", http.StatusInternalServerError)
			return
		}

		var allTasks types.TodoList = types.TodoList{
			Tasks: tasks,
		}
		// Create a template using the html
		tmpl, err := template.ParseFiles("view.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, allTasks)

		// // Convert the tasks list to JSON
		// w.Header().Set("Content-Type", "application/json")
		// err = json.NewEncoder(w).Encode(tasks)
		// if err != nil {
		// 	http.Error(w, "Failed to encode tasks to JSON", http.StatusInternalServerError)
		// 	return
		// }

		// // Successfully return the tasks
		// w.WriteHeader(http.StatusOK)
		// return

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
	var taskDto types.TaskDto
	var err error = utils.ParseJson(r, &taskDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// validator

	var task types.Task = types.Task{
		ID:          uuid.New(),
		Title:       taskDto.Title,
		Description: taskDto.Description,
		Status:      taskDto.Status,
		Deleted:     false,
	}
	err = h.repository.CreateTask(task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *Handler) handleUpdateTasks(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleDeleteTasks(w http.ResponseWriter, r *http.Request) {

}
