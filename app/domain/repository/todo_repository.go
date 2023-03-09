package repository

import "github.com/sako0/todo-api/app/domain/model"

type TodoRepository interface {
	PostTodo(text string) error
	ListTodo() []model.Todo
	DeleteTodo(id uint) error
	UpdateTodoText(id uint, text string) error
	GetTodoById(id uint) model.Todo
}