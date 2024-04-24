package todo

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	list := List{}
	task := "todo"

	item := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	list = append(list, item)

	if len(list) != 1 {
		t.Errorf("got %d, want 1", len(list))
	}
}
