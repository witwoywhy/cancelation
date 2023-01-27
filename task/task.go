package task

import (
	"context"
	"time"
)

type Status struct {
	Id       int
	IsCancel bool
}

type Task struct {
	Id     int
	ctx    context.Context
	Cancel context.CancelFunc
}

func (t *Task) Do(ch chan Status) {
	s := Status{
		Id: t.Id,
	}

	select {
	case <-t.ctx.Done():
		s.IsCancel = true
		// fmt.Println("Task", t.Id, "Cancel")
		ch <- s
		return
	case <-time.After(3 * time.Second):
		// fmt.Println("Task", t.Id, "Do")
		ch <- s
		return
	}
}

func NewTask(id int) *Task {
	ctx, cancel := context.WithCancel(context.Background())

	return &Task{
		Id:     id,
		ctx:    ctx,
		Cancel: cancel,
	}
}
