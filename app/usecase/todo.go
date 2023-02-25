package usecase

import (
	"fmt"

	"github.com/sako0/todo-api/app/domain/model"
	"github.com/sako0/todo-api/app/domain/repository"
)

type TodoUsecase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUsecase(tr repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{todoRepository: tr}
}

func (tu TodoUsecase) PostTodo(text string) error {
	err := tu.todoRepository.PostTodo(text)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (tu TodoUsecase) ListTodo() []model.Todo {
	todos := tu.todoRepository.ListTodo()

	return todos
}

func (tu TodoUsecase) DeleteTodo(id uint) error {
	err := tu.todoRepository.DeleteTodo(id)

	return err
}

func (tu TodoUsecase) UpdateTodoText(id uint, text string) error {
	err := tu.todoRepository.UpdateTodoText(id, text)

	return err
}
