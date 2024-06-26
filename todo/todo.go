package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	item := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, item)
}

func (l *List) Complete(itemNumber int) error {
	list := *l
	if itemNumber <= 0 || itemNumber > len(list) {
		return fmt.Errorf("item %d does not exist", itemNumber)
	}

	list[itemNumber-1].Done = true
	list[itemNumber-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(itemNumber int) error {
	list := *l

	if itemNumber <= 0 || itemNumber > len(list) {
		return fmt.Errorf("item %d does not exist", itemNumber)
	}

	*l = append(list[:itemNumber-1], list[itemNumber:]...)

	return nil
}

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "X "
		}

		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}

	return formatted
}
