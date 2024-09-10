package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"todo/types"
)

func ParseJson(request *http.Request, obj any) error {
	if request.Body == nil {
		return fmt.Errorf("missing request body")
	}

	// Decode the incoming JSON into the TaskDto struct
	var err error = json.NewDecoder(request.Body).Decode(&obj)
	if err != nil {
		// Handle error if JSON is invalid or doesn't match the TaskDto struct
		return fmt.Errorf("invalid request payload. error decoding request body %s", err)
	}

	return err
}

func WriteTasksToJSONFile(filename string, tasks []types.Task) error {
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

func ReadTasksFromJSONFile(filename string) ([]types.Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var tasks []types.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
