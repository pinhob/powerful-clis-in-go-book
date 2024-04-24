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
