package repository

import "github.com/sako0/todo-api/app/domain/model"

type TodoRepository interface {
	PostTodo(text string) error
	ListTodo() []model.Todo
	DeleteTodo(id uint) error
}
