package task

import (
	"database/sql"
	"fmt"
	"log"
	"todo/types"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetTaskById(id uuid.UUID) (*types.Task, error) {
	rows, err := r.db.Query("SELECT * from tasks WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var task *types.Task = new(types.Task)
	for rows.Next() {
		task, err = scanRowsIntoTask(rows)
		if err != nil {
			return nil, err
		}
	}

	if task.ID == uuid.Nil {
		return nil, fmt.Errorf("user not found")
	}

	return task, nil
}

func (r *Repository) GetAllTasks() ([]*types.Task, error) {
	rows, err := r.db.Query("SELECT * FROM tasks WHERE deleted = ? ORDER BY status", false)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var tasks []*types.Task
	for rows.Next() {
		task, err := scanRowsIntoTask(rows)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		log.Println("No tasks found")
	}

	return tasks, nil
}

func scanRowsIntoTask(rows *sql.Rows) (*types.Task, error) {
	var task *types.Task = new(types.Task)

	err := rows.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.Deleted,
	)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *Repository) CreateTask(task types.Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (id, title, description, status, deleted) VALUES (?,?,?,?,?)",
		task.ID, task.Title, task.Description, types.Active, false)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteTask(id uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
