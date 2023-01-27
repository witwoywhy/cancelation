package main

import (
	"cancel/task"
	"fmt"
)

func main() {
	tasks := make([]*task.Task, 5)
	for i := 0; i < 5; i++ {
		tasks[i] = task.NewTask(i + 1)
	}

	ch := make(chan task.Status)
	for _, task := range tasks {
		go task.Do(ch)
	}

	for i, task := range tasks {
		if i%2 == 0 {
			task.Cancel()
		}
	}

	for i := 0; i < 5; i++ {
		s := <-ch
		fmt.Println("Result => Task: ", s.Id, "Is Cancel", s.IsCancel)
	}

}
