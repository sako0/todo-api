package main

import (
	"fmt"
	"os"

	"github.com/sako0/todo-api/app/config"
	"github.com/sako0/todo-api/app/domain/model"
	"github.com/sako0/todo-api/app/infra"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(os.Args) > 1 && os.Args[1] == "down" {
		err = down(cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	if err := migrate(cfg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg, err = config.LoadTestConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := migrate(cfg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func migrate(cfg *config.AppConfig) error {
	db, err := infra.NewSQLConnection(cfg.AppInfo.DatabaseURL)
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.Todo{})
	if err != nil {
		return err
	}

	todo := model.Todo{Text: "test"}

	err = db.Create(&todo).Error
	if err != nil {
		return err
	}

	return nil
}

func down(cfg *config.AppConfig) error {
	db, err := infra.NewSQLConnection(cfg.AppInfo.DatabaseURL)
	if err != nil {
		return err
	}

	err = db.Migrator().DropTable(&model.Todo{})
	if err != nil {
		return err
	}

	return nil
}
