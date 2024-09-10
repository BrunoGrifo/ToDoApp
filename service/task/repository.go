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
	log.Println("checkpoint")
	rows, err := r.db.Query("SELECT * from tasks WHERE id = ?", id)
	log.Println("checkpoint 1")
	if err != nil {
		log.Println("upsi")
		return nil, err
	}
	log.Println("checkpoint 2")
	var task *types.Task = new(types.Task)
	log.Println("hmmm")
	for rows.Next() {
		log.Println("loop")
		task, err = scanRowsIntoTask(rows)
		if err != nil {
			log.Println("buuuuu")
			return nil, err
		}
	}

	if task.ID == uuid.Nil {
		log.Println("hello there")
		return nil, fmt.Errorf("user not found")
	}

	return task, nil
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

func (r *Repository) GetAllTasks() (*types.Task, error) {
	return nil, nil
}

func (r *Repository) CreateTask(task types.Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (id, title, description, status, deleted) VALUES (?,?,?,?,?)",
		task.ID, task.Title, task.Description, types.Active, false)
	if err != nil {
		return err
	}
	return nil
}
