package main

import (
	"fmt"
	"sync"
	"time"
	"todo/model"
)

func printTasks(tasks ...model.Task) {
	for _, task := range tasks {
		fmt.Printf("Title: %s\n", task.Title)
		fmt.Printf("Status: %v\n", task.Status)
		fmt.Printf("Deleted: %v\n", task.Deleted)

		task.Body.Print()

		fmt.Println("------")
	}
}

func updateValueOdd(data *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i += 2 {
		*data = i
		fmt.Println("Odd data: ", *data)
	}
}

func updateValueEven(data *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 100; i += 2 {
		*data = i
		fmt.Println("Even data: ", *data)
	}
}

func main() {
	// task1 := model.Task{
	// 	Title:   "Buy groceries",
	// 	Body:    model.Note("Milk, Bread, Eggs"),
	// 	Status:  model.Completed,
	// 	Deleted: false,
	// }

	// task2 := model.Task{
	// 	Title: "Morning Routine",
	// 	Body: model.TickBoxList{
	// 		{Description: "Brush teeth", Checked: true},
	// 		{Description: "Have breakfast", Checked: false},
	// 	},
	// 	Status:  model.Active,
	// 	Deleted: false,
	// }

	// printTasks(task1, task2)

	// err := utils.WriteTasksToJSONFile("tasks.json", []model.Task{task1, task2})
	// if err != nil {
	// 	fmt.Println("Error writing tasks to JSON file:", err)
	// }

	// tasks, err := utils.ReadTasksFromJSONFile("tasks.json")
	// if err != nil {
	// 	fmt.Println("Error reading tasks from JSON file:", err)
	// 	return
	// }

	// printTasks(tasks...)

	var data int
	var wg sync.WaitGroup

	wg.Add(2)
	start := time.Now()
	go updateValueOdd(&data, &wg)
	go updateValueEven(&data, &wg)

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("time routines: ", elapsed)
}
