package main

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var expectedTask Task
var expectedTaskJSON string

func TestMain(m *testing.M) {
	expectedTask = Task{
		ID:          1,
		Description: "Unit test for a task",
		Status:      CompleteStatus,
	}
	expectedTaskJSON = `{"id":1,"description":"Unit test for a task","status":"Complete"}`
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestToJSON(t *testing.T) {
	result, _ := expectedTask.ToJSON()
	if diff := cmp.Diff(expectedTaskJSON, result); diff != "" {
		// t.Error("marshall failed: expected", taskJSON, "got", err)
		t.Error(diff)
	}
}

func TestFromJSON(t *testing.T) {
	var result Task

	// err := (&result).FromJSON(`{"id":100,"description":"Unit test for a task","status":"Complete"}`)
	// if err != nil {
	// 	t.Error("unmarshall failed: expected", task, "got", err)
	// }

	(&result).FromJSON(expectedTaskJSON)

	if diff := cmp.Diff(expectedTask, result); diff != "" {
		t.Error(diff)
	}
}
