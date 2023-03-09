package model

import "testing"

func TestTodo(t *testing.T) {
	t.Run("IsDone", func(t *testing.T) {
		todo := Todo{Text: "test"}
		if got := todo.IsDone(); got != true {
			t.Errorf("Todo.IsDone() = %v, want true because Text is not empty", got)
		}
	})
}
