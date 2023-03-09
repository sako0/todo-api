package infra

import (
	"testing"

	"github.com/sako0/todo-api/app/config"
	"github.com/sako0/todo-api/app/domain/model"
	"github.com/sako0/todo-api/app/infra"
)

// テストコード
func TestTodoRepository(t *testing.T) {
	// 設定読み込み
	cfg, err := config.LoadTestConfig()
	if err != nil {
		t.Fatalf("failed to load test config: %v", err)
	}

	// データベース接続
	db, err := infra.NewSQLConnection(cfg.AppInfo.DatabaseURL)
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	// dbにすでにテーブルがある場合は削除する
	db.Migrator().DropTable(&model.Todo{})
	// dbにテーブルを作成する
	db.AutoMigrate(&model.Todo{})

	repo := NewTodoRepository(db)

	// テストケース
	t.Run("PostTodo", func(t *testing.T) {

		todoList, err := repo.ListTodo()
		if len(todoList) > 0 {
			for _, todo := range todoList {
				repo.DeleteTodo(todo.ID)
			}
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		err = repo.PostTodo("test")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("ListTodo", func(t *testing.T) {
		todoList, err := repo.ListTodo()
		if len(todoList) != 1 {
			t.Fatalf("unexpected length of todo list: %d", len(todoList))
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("GetTodoById", func(t *testing.T) {
		todo, err := repo.GetTodoById(1)
		if todo.ID != 1 || todo.Text != "test" {
			t.Fatalf("unexpected todo: %+v", todo)
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("UpdateTodoText", func(t *testing.T) {
		err := repo.UpdateTodoText(1, "updated")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		todo, err := repo.GetTodoById(1)
		if todo.Text != "updated" {
			t.Fatalf("unexpected todo text: %s", todo.Text)
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("DeleteTodo", func(t *testing.T) {
		err := repo.DeleteTodo(1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		todoList, err := repo.ListTodo()
		if len(todoList) != 0 {
			t.Fatalf("unexpected length of todo list: %d", len(todoList))
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
