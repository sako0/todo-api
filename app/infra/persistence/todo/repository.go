package infra

import (
	"github.com/sako0/todo-api/app/domain/model"
	"github.com/sako0/todo-api/app/domain/repository"
	"gorm.io/gorm"
)

type todoRepository struct {
	Conn *gorm.DB
}

func NewTodoRepository(conn *gorm.DB) repository.TodoRepository {
	return &todoRepository{Conn: conn}

}

func (tr todoRepository) PostTodo(text string) error {
	_ = tr.Conn.Create(&model.Todo{Text: text})
	return nil
}

func (tr todoRepository) ListTodo() []model.Todo {
	var users []model.Todo
	db := tr.Conn
	db.Find(&users)
	return users
}

func (tr todoRepository) DeleteTodo(id uint) error {
	db := tr.Conn
	db.Delete(&model.Todo{}, id)
	return nil
}

func (tr todoRepository) UpdateTodoText(id uint, text string) error {
	db := tr.Conn
	db.Model(&model.Todo{}).Where("id = ?", id).Update("text", text)
	return nil
}

func (tr todoRepository) GetTodoById(id uint) model.Todo {
	var todo model.Todo
	db := tr.Conn
	db.First(&todo, id)
	return todo
}
