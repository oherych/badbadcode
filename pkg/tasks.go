package pkg

import (
	"fmt"
	"sync"
)

type Task struct {
	TaskName        string
	TaskDescription string
}

type Result struct {
	Values []int
}

type TaskManager struct {
	List map[Task]string

	mx sync.Mutex
}

func (tm TaskManager) Lock() {
	tm.mx.Lock()
}

func (tm TaskManager) Unlock() {
	tm.mx.Unlock()
}

func (tm TaskManager) Go(task Task, r *Result) {
	r.Values = []int{}
	r.Values = append(r.Values, len(fmt.Sprintf("%v%v", task.TaskName, task.TaskDescription)))
	r.Values = append(r.Values, 23)
	r.Values = append(r.Values, 0)

	tm.Lock()
	delete(tm.List, task)
	tm.Unlock()
}
