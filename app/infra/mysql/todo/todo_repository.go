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
	tx := tr.Conn.Create(&model.Todo{Text: text})
	return tx.Error
}

func (tr todoRepository) ListTodo() ([]model.Todo, error) {
	var users []model.Todo
	db := tr.Conn
	tx := db.Find(&users)
	return users, tx.Error
}

func (tr todoRepository) DeleteTodo(id uint) error {
	db := tr.Conn
	tx := db.Delete(&model.Todo{}, id)
	return tx.Error
}

func (tr todoRepository) UpdateTodoText(id uint, text string) error {
	db := tr.Conn
	tx := db.Model(&model.Todo{}).Where("id = ?", id).Update("text", text)
	return tx.Error
}

func (tr todoRepository) GetTodoById(id uint) (model.Todo, error) {
	var todo model.Todo
	db := tr.Conn
	tx := db.First(&todo, id)
	return todo, tx.Error
}
