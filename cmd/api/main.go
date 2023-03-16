package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sako0/todo-api/app/config"
	"github.com/sako0/todo-api/app/infra"
	todo "github.com/sako0/todo-api/app/infra/mysql/todo"
	handler "github.com/sako0/todo-api/app/interfaces/rest"
	"github.com/sako0/todo-api/app/interfaces/validator"
	"github.com/sako0/todo-api/app/usecase"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Use(middleware.CORS())
	e.Validator = validator.NewValidator()
	// 設定読み込み
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	// データベース接続
	db, err := infra.NewSQLConnection(cfg.AppInfo.DatabaseURL)
	if err != nil {
		panic(err)
	}

	tr := todo.NewTodoRepository(db)
	tu := usecase.NewTodoUsecase(tr)
	th := handler.NewTodoHandler(tu)

	// todoリストの更新
	e.POST("/todo", th.PostTodo)
	// todoリストを1つ表示する
	e.GET("/todo/:id", th.GetTodoById)
	// todoリストをすべて表示する
	e.GET("/todos", th.ListTodo)
	// todoリストを削除する
	e.DELETE("/todo/:id", th.DeleteTodo)
	// todoリストのテキストを更新する
	e.PUT("/todo/:id", th.UpdateTodoText)

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}
