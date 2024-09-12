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
	mockTaskRepository := &mockTaskRepository{}
	handler := NewHandler(mockTaskRepository)

	t.Run("Should create a task", func(t *testing.T) {
		var taskdto dto.TaskDto = dto.TaskDto{
			Title:       "This is a task",
			Description: "This is a description",
		}

		payload, _ := json.Marshal(taskdto)
		req, err := http.NewRequest(http.MethodPost, "/task", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
		var router *http.ServeMux = http.NewServeMux()
		router.HandleFunc("POST /task", handler.handleCreateTasks)
		router.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusSeeOther {
			t.Errorf("expected status code %d, got %d", http.StatusSeeOther, responseRecorder.Code)
		}

	})

	t.Run("Should delete a task successfully", func(t *testing.T) {
		taskID := uuid.New().String()

		req, err := http.NewRequest(http.MethodDelete, "/task/"+taskID, nil)
		if err != nil {
			t.Fatal(err)
		}

		var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
		var router *http.ServeMux = http.NewServeMux()
		router.HandleFunc("/task/{id}", handler.handleDeleteTasks)
		router.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, responseRecorder.Code)
		}
	})

	t.Run("Should fail with invalid task ID", func(t *testing.T) {
		invalidTaskID := "invalid-uuid"

		req, err := http.NewRequest(http.MethodDelete, "/task/"+invalidTaskID, nil)
		if err != nil {
			t.Fatal(err)
		}

		var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
		var router *http.ServeMux = http.NewServeMux()
		router.HandleFunc("/task/{id}", handler.handleDeleteTasks)
		router.ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, responseRecorder.Code)
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
