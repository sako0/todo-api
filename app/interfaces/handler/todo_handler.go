package handler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sako0/todo-api/app/interfaces/presenter"
	"github.com/sako0/todo-api/app/interfaces/request"
	"github.com/sako0/todo-api/app/usecase"
)

type TodoHandler interface {
	PostTodo(c echo.Context) error
	GetTodoById(c echo.Context) error
	ListTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
	UpdateTodoText(c echo.Context) error
}

type todoHandler struct {
	todoUsecase   usecase.TodoUsecase
	todoPresenter presenter.TodoPresenter
	todoRequest   request.TodoRequest
}

func NewTodoHandler(usecase usecase.TodoUsecase, presenter presenter.TodoPresenter, request request.TodoRequest) TodoHandler {
	return &todoHandler{todoUsecase: usecase, todoPresenter: presenter, todoRequest: request}
}

func (th *todoHandler) PostTodo(c echo.Context) error {
	r := &request.PostTodoRequest{}
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
	return c.JSON(http.StatusOK, map[string]string{"message": "成功しました！"})
}

func (th *todoHandler) GetTodoById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	todo, err := th.todoUsecase.GetTodoById(uint(id))
	if todo.ID == 0 {
		return c.JSON(http.StatusNotFound, "レコードがありません")
	}
	res, err := th.todoPresenter.GetTodoByIdResponse(todo, err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (th *todoHandler) ListTodo(c echo.Context) error {
	todos, err := th.todoUsecase.ListTodo()
	res, err := th.todoPresenter.ListTodoResponse(todos, err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (th *todoHandler) DeleteTodo(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = th.todoUsecase.DeleteTodo(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}

func (th todoHandler) UpdateTodoText(c echo.Context) error {
	r := &request.UpdateTodoTextRequest{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := c.Validate(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.(validator.ValidationErrors).Error())
	}
	err = th.todoUsecase.UpdateTodoText(uint(id), r.Text)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
