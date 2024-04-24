package todo_test

import (
	"os"
	"testing"
	"todo"
)

func TestAdd(t *testing.T) {
	list := todo.List{}
	task := "todo"

	list.Add(task)

	if len(list) != 1 {
		t.Errorf("got %d, want 1", len(list))
	}
	if list[0].Task != task {
		t.Errorf("Expected %q, got %q instead", task, list[0].Task)
	}
}

func TestComplete(t *testing.T) {
	list := todo.List{}
	task := "todo"
	taskNumber := 1

	list.Add(task)

	if list[taskNumber-1].Task != task {
		t.Errorf("Expected %q, got %q instead", task, list[0].Task)
	}

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

func TestDelete(t *testing.T) {
	list := todo.List{}
	taskNumber := 1

	tasks := []string{
		"task 1",
		"task 2",
		"task 3",
	}

	for _, task := range tasks {
		list.Add(task)
	}

	if len(list) != 3 {
		t.Errorf("list should have the length of 3, but got length of %d", len(list))
	}

	if list[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead", tasks[0], list[0].Task)
	}

	list.Delete(taskNumber)

	if len(list) != 2 {
		t.Errorf("list should have the length of 0, but got length of %d", len(list))
	}

	if list[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead", tasks[2], list[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}
	taskName := "New task"

	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Want %q, got %q instead.", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task,", l1[0].Task, l2[0].Task)
	}
}
