package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

func (t *Todos) Delete(index int) (bool, error) {
	ls := *t

	if index <= 0 || index >= len(ls) {
		return false, fmt.Errorf("out of bound index (%d)", index)
	}

	// delete task, everything except the index
	*t = append(ls[:index], ls[index+1:]...)
	return true, nil
}

func (t *Todos) Load(filename string) (bool, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		// check if file exists, like the first time running the program
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}

	if len(data) == 0 {
		return false, fmt.Errorf("%d data loaded", len(data))
	}

	err = json.Unmarshal(data, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
