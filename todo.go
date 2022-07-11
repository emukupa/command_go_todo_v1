package todo

import (
	"fmt"
	"time"
)

type item struct {
	Task       string
	Done       bool
	CreatedAt  time.Time
	CompleteAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:       task,
		Done:       false,
		CreatedAt:  time.Now(),
		CompleteAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) (bool, error) {
	ls := *t

	if index <= 0 || index >= len(ls) {
		return false, fmt.Errorf("out of bound index (%d)", index)
	}

	// complete
	ls[index].Done = true
	ls[index].CompleteAt = time.Now()

	return true, nil
}
