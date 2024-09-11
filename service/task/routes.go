package task

import (
	"html/template"
	"log"
	"net/http"
	"todo/dto"
	"todo/mappers"
	"todo/types"

	"github.com/google/uuid"
	"github.com/gorilla/csrf"
)

type Handler struct {
	repository types.TaskRepository
}

func NewHandler(repository types.TaskRepository) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /task", h.handleGetAllTasks)
	router.HandleFunc("GET /task/{id}", h.handleGetTasks)
	router.HandleFunc("POST /task", h.handleCreateTasks)
	router.HandleFunc("PUT /task", h.handleUpdateTasks)
	router.HandleFunc("DELETE /task/{id}", h.handleDeleteTasks)
	router.HandleFunc("GET /show_csrf_form", h.handleCsrfForm)

}

func (h *Handler) handleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	// get all
	log.Println("Getting all tasks...")
	tasks, err := h.repository.GetAllTasks()
	if err != nil {
		http.Error(w, "Internel server error", http.StatusInternalServerError)
		return
	}

	var allTasks dto.TodoList = dto.TodoList{
		Tasks:     mappers.FromTasksToDto(tasks),
		CsrfToken: csrf.Token(r),
	}
	// Create a template using the html
	tmpl, err := template.ParseFiles("templates/view.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, allTasks)

	// // Successfully return the tasks
	// w.WriteHeader(http.StatusOK)
	// return
}

func (h *Handler) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}
	task, err := h.repository.GetTaskById(uuid)
	if err != nil {
		log.Fatal(err)
	}

	taskDto := mappers.FromTaskToDto(task)

	// Create a template using the html
	tmpl, err := template.ParseFiles("templates/update_task_form.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, map[string]interface{}{
		"Task":      taskDto,
		"CsrfToken": csrf.Token(r),
	})

}

func (h *Handler) handleCreateTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating task...")
	// CSRF validation has already been handled by the middleware at this point
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	var name string = r.FormValue("title")
	var description string = r.FormValue("description")

	// // validator

	var task types.Task = types.Task{
		ID:          uuid.New(),
		Title:       name,
		Description: description,
		Status:      types.Active,
		Deleted:     false,
	}
	err = h.repository.CreateTask(task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/task", http.StatusSeeOther)
}

func (h *Handler) handleUpdateTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("yeeeahhhhhh boyyyy")

}

func (h *Handler) handleDeleteTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting task...")
	id := r.PathValue("id")
	log.Println(id)
	// id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Missing task ID", http.StatusBadRequest)
		return
	}

	taskID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = h.repository.DeleteTask(taskID)
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) handleCsrfForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/task_form.html")
	if err != nil {
		log.Fatal(err)
	}
	csrfField := csrf.TemplateField(r)
	tmpl.Execute(w, map[string]interface{}{
		csrf.TemplateTag: csrfField,
	})

}

// func (h *Handler) handleCreateTasks(w http.ResponseWriter, r *http.Request) {
// 	var taskDto types.TaskDto
// 	var err error = utils.ParseJson(r, &taskDto)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}
// 	// validator

// 	var task types.Task = types.Task{
// 		ID:          uuid.New(),
// 		Title:       taskDto.Title,
// 		Description: taskDto.Description,
// 		Status:      taskDto.Status,
// 		Deleted:     false,
// 	}
// 	err = h.repository.CreateTask(task)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// }
