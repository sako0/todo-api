package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/sako0/todo-api/app/interfaces/handler"
)

func NewTodoRouter(e *echo.Echo, handler handler.TodoHandler) {
	e.POST("/todo", handler.PostTodo)
	e.GET("/todo/:id", handler.GetTodoById)
	e.GET("/todos", handler.ListTodo)
	e.DELETE("/todo/:id", handler.DeleteTodo)
	e.PUT("/todo/:id", handler.UpdateTodoText)
}
