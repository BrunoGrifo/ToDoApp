package task

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/dto"
	"todo/types"

	"github.com/google/uuid"
)

func TestTaskServiceHandlers(t *testing.T) {
	tastRepository := &mockTaskRepository{}
	handler := NewHandler(tastRepository)

	t.Run("Should create a task", func(t *testing.T) {
		var taskdto dto.TaskDto = dto.TaskDto{
			Title:       "This is a task",
			Description: "This is a description",
		}

		payload, _ := json.Marshal(taskdto)
		req, err := http.NewRequest(http.MethodPost, "/task/{id}", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
		var router *http.ServeMux = http.NewServeMux()
		router.HandleFunc("POST /task", handler.handleCreateTasks)
		router.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusSeeOther {
			t.Errorf("expected status code %d, got %d", http.StatusOK, responseRecorder.Code)
		}

	})

}

type mockTaskRepository struct{}

func (r *mockTaskRepository) GetTaskById(id uuid.UUID) (*types.Task, error) {
	return nil, nil
}

func (r *mockTaskRepository) GetAllTasks() ([]*types.Task, error) {
	return nil, nil
}

func (r *mockTaskRepository) CreateTask(task *types.Task) error {
	return nil
}

func (r *mockTaskRepository) DeleteTask(id uuid.UUID) error {
	return nil
}

func (r *mockTaskRepository) UpdateTask(task *types.Task) error {
	return nil
}
