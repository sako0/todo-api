package usecase

import (
	"fmt"

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
