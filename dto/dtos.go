package dto

import "github.com/google/uuid"

type Status int

const (
	Active Status = iota
	Completed
)

type TodoList struct {
	Tasks     []*TaskDto
	CsrfToken string
}

type TaskDto struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	Deleted     bool      `json:"deleted"`
}
