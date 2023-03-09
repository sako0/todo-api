package repository

import "github.com/sako0/todo-api/app/domain/model"

type TodoRepository interface {
	PostTodo(text string) error
	ListTodo() ([]model.Todo, error)
	DeleteTodo(id uint) error
	UpdateTodoText(id uint, text string) error
	GetTodoById(id uint) (model.Todo, error)
}

// モックデータ作成のコマンドは下記
// $ mockgen -source=todo_repository.go -destination=mock_todo_repository/mock_todo_repository.go -package=mock_todo_repository
