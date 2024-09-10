package concurrency

import (
	"fmt"
	"sync"
	"time"
	"todo/model"
)

// Worker pool definition
type WorkerPool struct {
	Tasks             []model.Task
	ConcurrentWorkers int
	tasksChan         chan model.Task
	wg                sync.WaitGroup
}

func printTitle(task model.Task, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Title: %s\n", task.Title)

	time.Sleep(1 * time.Second)
}

func printStatus(task model.Task, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Status: %b\n", task.Status)

	time.Sleep(1 * time.Second)
}

// Worker function
func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {

		go printTitle(task, &wp.wg)
		go printStatus(task, &wp.wg)
	}
}

func (wp *WorkerPool) Run() {
	// Initialize the tasks channel
	wp.tasksChan = make(chan model.Task, len(wp.Tasks))

	// Start workers
	for i := 0; i < wp.ConcurrentWorkers; i++ {
		wp.wg.Add(2)
		go wp.worker()
	}

	// Send tasks to the tasks channel
	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}
	close(wp.tasksChan)

	// Wait for all tasks to finish
	wp.wg.Wait()
}
