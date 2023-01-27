package task

import "testing"

func TestDo(t *testing.T) {
	t.Run("can Cancel", func(t *testing.T) {
		task := NewTask(1)

		ch := make(chan Status)
		go task.Do(ch)

		task.Cancel()

		got := <-ch

		if !got.IsCancel {
			t.Errorf("want true but got %v", got.IsCancel)
		}
	})

	t.Run("can Do", func(t *testing.T) {
		task := NewTask(1)

		ch := make(chan Status)
		go task.Do(ch)

		got := <-ch

		if got.IsCancel {
			t.Errorf("want false but got %v", got.IsCancel)
		}
	})
}
