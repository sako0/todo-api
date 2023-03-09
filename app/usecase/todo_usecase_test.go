package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sako0/todo-api/app/domain/model"
	mock "github.com/sako0/todo-api/app/testdata"
)

func TestTodoUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRepository := mock.NewMockTodoRepository(ctrl)
	mockTodoRepository.EXPECT().PostTodo("test").Return(nil)
	mockTodoRepository.EXPECT().ListTodo().Return([]model.Todo{}, nil)
	mockTodoRepository.EXPECT().GetTodoById(uint(1)).Return(model.Todo{}, nil)
	mockTodoRepository.EXPECT().UpdateTodoText(uint(1), "test").Return(nil)
	mockTodoRepository.EXPECT().DeleteTodo(uint(1)).Return(nil)

	tu := NewTodoUsecase(mockTodoRepository)

	t.Run("PostTodo", func(t *testing.T) {
		err := tu.PostTodo("test")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
	t.Run("ListTodo", func(t *testing.T) {
		todoList, err := tu.ListTodo()
		if len(todoList) != 0 {
			t.Errorf("unexpected length of todo list: %d", len(todoList))
		}
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
	t.Run("GetTodoById", func(t *testing.T) {
		// intをuintに変更
		todo, err := tu.GetTodoById(uint(1))
		if todo.ID != 0 {
			t.Errorf("unexpected todo: %+v", todo)
		}
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
	t.Run("UpdateTodoText", func(t *testing.T) {
		err := tu.UpdateTodoText(uint(1), "test")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
	t.Run("DeleteTodo", func(t *testing.T) {
		err := tu.DeleteTodo(uint(1))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
