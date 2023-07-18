// main.go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sako0/todo-api/app/bootstrap"
	"github.com/sako0/todo-api/app/routers"
)

func main() {
	e := echo.New()
	appContainer, err := bootstrap.InitializeApp()
	if err != nil {
		panic(err)
	}
	routers.SetupRoutes(e, appContainer)
	routers.StartServer(e)
}
