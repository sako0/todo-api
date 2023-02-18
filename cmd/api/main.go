package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sako0/todo-api/app/config"
	"github.com/sako0/todo-api/app/infra"
	todo "github.com/sako0/todo-api/app/infra/persistence/todo"
	handler "github.com/sako0/todo-api/app/interfaces/handler"
	validator "github.com/sako0/todo-api/app/interfaces/validator"
	"github.com/sako0/todo-api/app/usecase"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Validator = validator.NewValidator()
	// 設定読み込み
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	// データベース接続
	db, err := infra.NewSQLConnection(cfg.AppInfo.DatabaseURL)
	if err != nil {
		panic(err)
	}

	tr := todo.NewTodoRepository(db)
	tu := usecase.NewTodoUsecase(tr)
	th := handler.NewTodoHandler(*tu)

	// todoリストの更新
	e.POST("/todo", th.PostTodo)
	// todoリストをすべて表示する
	e.GET("/todos", th.ListTodo)

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}
