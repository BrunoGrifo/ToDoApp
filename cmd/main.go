package main

import (
	"database/sql"
	"log"
	"todo/cmd/api"
	"todo/config"
	"todo/db"

	"github.com/go-sql-driver/mysql"
)

func main() {
	var cfg mysql.Config = mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Net:                  "tcp",
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := db.NewMySqlStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)
	defer db.Close()

	var server api.APIServer = api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	log.Println("DB: Connecting...")
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected to todoapp database!")
}

// type DataContainer struct {
// 	value int
// 	mutex sync.Mutex
// }

// func printTasks(tasks ...model.Task) {
// 	for _, task := range tasks {
// 		fmt.Printf("Title: %s\n", task.Title)
// 		fmt.Printf("Status: %v\n", task.Status)
// 		fmt.Printf("Deleted: %v\n", task.Deleted)

// 		task.Body.Print()

// 		fmt.Println("------")
// 	}
// }

// func updateValueEven(data *DataContainer, wg *sync.WaitGroup, ch chan<- string) {
// 	defer wg.Done()

// 	for i := 2; i <= 6; i += 2 {
// 		data.mutex.Lock()
// 		data.value = i
// 		ch <- fmt.Sprintf("Even data: %d", data.value)
// 		data.mutex.Unlock()
// 	}
// }

// func updateValueOdd(data *DataContainer, wg *sync.WaitGroup, ch chan<- string) {
// 	defer wg.Done()

// 	for i := 1; i <= 6; i += 2 {
// 		data.mutex.Lock()
// 		data.value = i
// 		ch <- fmt.Sprintf("Odd data: %d", data.value)
// 		data.mutex.Unlock()
// 	}
// }

// func main() {
// 	var server api.APIServer = api.NewApiServer(":8080", nil)
// 	if err := server.Run(); err != nil {
// 		log.Fatal(err)
// 	}

// tasks := []model.Task{
// 	model.Task{
// 		Title:   "Buy groceries",
// 		Body:    model.Note("Milk, Bread, Eggs"),
// 		Status:  model.Completed,
// 		Deleted: false,
// 	},
// 	model.Task{
// 		Title: "Morning Routine",
// 		Body: model.TickBoxList{
// 			{Description: "Brush teeth", Checked: true},
// 			{Description: "Have breakfast", Checked: false},
// 		},
// 		Status:  model.Active,
// 		Deleted: false,
// 	},
// }

// // Create a worker pool
// wp := concurrency.WorkerPool{
// 	Tasks:             tasks,
// 	ConcurrentWorkers: 2, // Number of workers that can run at a time
// }

// Run the pool
// wp.Run()
// fmt.Println("All tasks have been processed!")
// var data DataContainer
// var wg sync.WaitGroup
// ch := make(chan string)

// wg.Add(2)

// go updateValueOdd(&data, &wg, ch)
// go updateValueEven(&data, &wg, ch)

// go func() {
// 	wg.Wait()
// 	close(ch)
// }()

// fmt.Println("ready to print")

// for val := range ch {
// 	fmt.Printf("Main Goroutine Received: %s\n", val)
// }

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

// }
