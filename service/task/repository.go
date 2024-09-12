package task

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
	"todo/types"

	"github.com/google/uuid"
)

type Repository struct {
	db    *sql.DB
	mutex sync.Mutex
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetTaskById(id uuid.UUID) (*types.Task, error) {
	startTime := time.Now()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	errorCh := make(chan error)
	resultCh := make(chan *types.Task)

	go func() {
		defer wg.Done()
		rows, err := r.db.Query("SELECT * from tasks WHERE id = ?", id)
		if err != nil {
			errorCh <- err
		}
		defer rows.Close()

		var task *types.Task = new(types.Task)
		for rows.Next() {
			task, err = scanRowsIntoTask(rows)
			if err != nil {
				errorCh <- err
			}
		}

		if task.ID == uuid.Nil {
			errorCh <- fmt.Errorf("user not found")
		}
		resultCh <- task
	}()

	go func() {
		wg.Wait()
		close(errorCh)
		close(resultCh)
	}()

	elapsed := time.Since(startTime)
	log.Printf("GetAllTasks took %s\n", elapsed)

	select {
	case tasks := <-resultCh:
		return tasks, nil
	case err := <-errorCh:
		return nil, err
	}
}

func (r *Repository) GetAllTasks() ([]*types.Task, error) {
	startTime := time.Now()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	errorCh := make(chan error)
	resultCh := make(chan []*types.Task)

	go func() {
		defer wg.Done()
		rows, err := r.db.Query("SELECT * FROM tasks WHERE deleted = ? ORDER BY status", false)
		if err != nil {
			log.Println(err)
			errorCh <- err
		}

		defer rows.Close()
		var tasks []*types.Task
		for rows.Next() {
			task, err := scanRowsIntoTask(rows)
			if err != nil {
				errorCh <- err
			}

			tasks = append(tasks, task)
		}

		if err = rows.Err(); err != nil {
			errorCh <- err
		}

		if len(tasks) == 0 {
			log.Println("No tasks found")
		}

		resultCh <- tasks
	}()

	go func() {
		wg.Wait()
		close(errorCh)
		close(resultCh)
	}()

	elapsed := time.Since(startTime)
	log.Printf("GetAllTasks took %s\n", elapsed)

	select {
	case tasks := <-resultCh:
		return tasks, nil
	case err := <-errorCh:
		return nil, err
	}

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

func (r *Repository) CreateTask(task *types.Task) error {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	errorCh := make(chan error)

	go func() {
		defer wg.Done()
		_, err := r.db.Exec("INSERT INTO tasks (id, title, description, status, deleted) VALUES (?,?,?,?,?)",
			task.ID, task.Title, task.Description, types.Active, false)
		if err != nil {
			errorCh <- err
		}
		errorCh <- nil
	}()
	go func() {
		wg.Wait()
		close(errorCh)
	}()

	if err := <-errorCh; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteTask(id uuid.UUID) error {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	errorCh := make(chan error)

	go func() {
		defer wg.Done()
		_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
		if err != nil {
			errorCh <- err
			return
		}
		errorCh <- nil
	}()
	go func() {
		wg.Wait()
		close(errorCh)
	}()

	if err := <-errorCh; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateTask(task *types.Task) error {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	errorCh := make(chan error)

	go func() {
		defer wg.Done()
		_, err := r.db.Exec("UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ?",
			task.Title, task.Description, task.Status, task.ID)
		if err != nil {
			errorCh <- fmt.Errorf("failed to update task with ID %d: %w", task.ID, err)
		}
		errorCh <- nil
	}()
	go func() {
		wg.Wait()
		close(errorCh)
	}()

	if err := <-errorCh; err != nil {
		return err
	}

	return nil
}
