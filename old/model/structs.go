package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type Note string
type Status int
type TickBoxList []TickBox

const (
	Active Status = iota
	Completed
)

type BodyPrinter interface {
	Print()
}

type Task struct {
	Title   string      `json:"title"`
	Body    BodyPrinter `json:"body"`
	Status  Status      `json:"status"`
	Deleted bool        `json:"deleted"`
}

type TickBox struct {
	Description string `json:"description"`
	Checked     bool   `json:"checked"`
}

func (n Note) Print() {
	fmt.Println("Body (Note):", string(n))
}

func (tbl TickBoxList) Print() {
	fmt.Println("Body (TickBoxes):")
	for i, tick := range tbl {
		fmt.Printf("  %d. %s [Checked: %v]\n", i+1, tick.Description, tick.Checked)
	}
}

// Custom unmarshaling logic for Task
func (t *Task) UnmarshalJSON(data []byte) error {
	// Define an intermediate struct to unmarshal common fields
	type Alias Task
	aux := &struct {
		Body json.RawMessage `json:"body"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	// Unmarshal into the intermediate struct
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Inspect the raw JSON to determine the type of Body
	var testType interface{}
	if err := json.Unmarshal(aux.Body, &testType); err != nil {
		return err
	}

	// Determine the actual type of Body and unmarshal accordingly
	switch reflect.TypeOf(testType).Kind() {
	case reflect.String:
		var note Note
		if err := json.Unmarshal(aux.Body, &note); err != nil {
			return err
		}
		t.Body = note
	case reflect.Slice:
		var tickBoxes TickBoxList
		if err := json.Unmarshal(aux.Body, &tickBoxes); err != nil {
			return err
		}
		t.Body = tickBoxes
	default:
		return errors.New("unknown body type")
	}

	return nil
}
