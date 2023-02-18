package interfaces

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sako0/todo-api/app/usecase"
)

type TodoHandler interface {
	PostTodo(c echo.Context) error
}

type todoHandler struct {
	todoUsecase usecase.TodoUsecase
}

type postTodoRequest struct {
	Text string `json:"text" validate:"required"`
}

func NewTodoHandler(usecase usecase.TodoUsecase) TodoHandler {
	return &todoHandler{todoUsecase: usecase}
}

func (th todoHandler) PostTodo(c echo.Context) error {
	r := &postTodoRequest{}
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := c.Validate(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.(validator.ValidationErrors).Error())
	}
	err := th.todoUsecase.PostTodo(r.Text)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "success")
}
