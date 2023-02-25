package infra

import (
	"fmt"
	"testing"

	"github.com/sako0/todo-api/app/config"
	"github.com/sako0/todo-api/app/domain/model"
	"github.com/sako0/todo-api/app/infra"
	"gorm.io/gorm"
)

// テストコード
func TestTodoRepository(t *testing.T) {
	// 設定読み込み
	cfg, err := config.LoadTestConfig()
	if err != nil {
		fmt.Println(err)
	}
	// データベース接続
	db, err := infra.NewSQLConnection(cfg.AppInfo.DatabaseURL)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Todo{
		Model: gorm.Model{},
		Text:  "",
	})
	repo := NewTodoRepository(db)

	// テストケース
	t.Run("PostTodo", func(t *testing.T) {
		repo.PostTodo("test")
	})
	t.Run("ListTodo", func(t *testing.T) {
		repo.ListTodo()
	})
	t.Run("DeleteTodo", func(t *testing.T) {
		repo.DeleteTodo(1)
	})
	t.Run("UpdateTodoText", func(t *testing.T) {
		repo.UpdateTodoText(1, "test")
	})
}
