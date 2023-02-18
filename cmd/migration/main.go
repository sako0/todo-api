package main

import (
	"fmt"

	"github.com/sako0/todo-api/app/config"
	"github.com/sako0/todo-api/app/domain/model"
	"github.com/sako0/todo-api/app/infra"
)

func main() {
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

	db.AutoMigrate(
		&model.Todo{},
	)
	todo := []model.Todo{{Text: "test"}}

	db.Create(&todo)
}
