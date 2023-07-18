// routers/routers.go
package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sako0/todo-api/app/bootstrap"
	"github.com/sako0/todo-api/app/interfaces/validator"
)

func SetupRoutes(e *echo.Echo, appContainer *bootstrap.AppContainer) {
	e.Use(middleware.CORS())
	e.Validator = validator.NewValidator()
	NewTodoRouter(e, appContainer.TodoHandler)
}

func StartServer(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":5000"))
}
