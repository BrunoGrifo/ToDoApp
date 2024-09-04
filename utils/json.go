package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"todo/model"
)

func WriteTasksToJSONFile(filename string, tasks []model.Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	fmt.Println("Tasks have been written to", filename)
	return nil
}

func ReadTasksFromJSONFile(filename string) ([]model.Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
