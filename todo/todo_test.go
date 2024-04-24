package todo

import (
	"testing"
)

func TestAdd(t *testing.T) {
	list := List{}
	task := "todo"

	list.Add(task)

	if len(list) != 1 {
		t.Errorf("got %d, want 1", len(list))
	}
}

func TestComplete(t *testing.T) {
	list := List{}
	task := "todo"
	taskNumber := 1

	list.Add(task)

	if list[taskNumber-1].Done {
		t.Errorf("Task should not be complete")
	}

	err := list.Complete(taskNumber)

	if err != nil {
		t.Errorf("Error when completing the task: %v", err)
	}

	got := list[0].Done
	want := true

	if got != true {
		t.Errorf("got %v, want %v", got, want)
	}
}
