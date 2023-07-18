package bootstrap

import (
	"github.com/sako0/todo-api/app/config"
	"github.com/sako0/todo-api/app/infra"
	todo "github.com/sako0/todo-api/app/infra/mysql/todo"
	"github.com/sako0/todo-api/app/interfaces/handler"
	"github.com/sako0/todo-api/app/interfaces/presenter"
	"github.com/sako0/todo-api/app/interfaces/request"
	"github.com/sako0/todo-api/app/usecase"
)

type AppContainer struct {
	TodoHandler handler.TodoHandler
}

func InitializeApp() (*AppContainer, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	db, err := infra.NewSQLConnection(cfg.AppInfo.DatabaseURL)
	if err != nil {
		return nil, err
	}

	tr := todo.NewTodoRepository(db)
	tu := usecase.NewTodoUsecase(tr)
	tReq := request.NewTodoRequest()
	tp := presenter.NewTodoPresenter(tu)
	th := handler.NewTodoHandler(tu, tp, tReq)

	return &AppContainer{
		TodoHandler: th,
	}, nil
}
