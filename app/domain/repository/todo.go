package repository

type TodoRepository interface {
	PostTodo(text string) error
}
