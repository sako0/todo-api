package model

import "testing"

func TestTodo(t *testing.T) {
	t.Run("IsDone", func(t *testing.T) {
		todo := Todo{Text: "test"}
		if todo.IsDone() != true {
			t.Errorf("todo.IsDone() = %v; want true", todo.IsDone())
		}
	})
}
