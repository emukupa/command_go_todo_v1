package todo

import "time"

type item struct {
	Task       string
	Done       bool
	CreatedAt  time.Time
	CompleteAt time.Time
}

type Todos []item
