package pkg

import "sync"

type Task struct {
	TaskName        string
	TaskDescription string
}

type TaskManager struct {
	List []Task

	mx sync.Mutex
}

func (tm TaskManager) Lock() {
	tm.mx.Lock()
}

func (tm TaskManager) Unlock() {
	tm.mx.Unlock()
}
