package usecase

import (
	"github.com/sako0/todo-api/app/domain/model"
	"github.com/sako0/todo-api/app/domain/repository"
)

type TodoUsecase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUsecase(tr repository.TodoRepository) TodoUsecase {
	return TodoUsecase{todoRepository: tr}
}

func (tu TodoUsecase) PostTodo(text string) error {
	err := tu.todoRepository.PostTodo(text)
	if err != nil {
		return err
	}
	return nil
}

func (tu TodoUsecase) ListTodo() ([]model.Todo, error) {
	todos, err := tu.todoRepository.ListTodo()

	return todos, err
}

func (tu TodoUsecase) DeleteTodo(id uint) error {
	err := tu.todoRepository.DeleteTodo(id)

	return err
}

func (tu TodoUsecase) UpdateTodoText(id uint, text string) error {
	err := tu.todoRepository.UpdateTodoText(id, text)

	return err
}

func (tu TodoUsecase) GetTodoById(id uint) (model.Todo, error) {
	todo, err := tu.todoRepository.GetTodoById(id)

	return todo, err
}
