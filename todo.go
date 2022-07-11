package todo

import "time"

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
