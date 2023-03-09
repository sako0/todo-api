package rest

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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
	todoUsecase usecase.TodoUsecase
}

type postTodoRequest struct {
	Text string `json:"text" validate:"required"`
}

type updateTodoTextRequest struct {
	Text string `json:"text" validate:"required"`
}

type getTodoByIdResponse struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

type listTodoResponse struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

func NewTodoHandler(usecase usecase.TodoUsecase) TodoHandler {
	return &todoHandler{todoUsecase: usecase}
}

func (th *todoHandler) PostTodo(c echo.Context) error {
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

func (th *todoHandler) GetTodoById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	todo := th.todoUsecase.GetTodoById(uint(id))
	if todo.ID == 0 {
		return c.JSON(http.StatusInternalServerError, "レコードがありません")
	}
	// レスポンスの構造体を定義して返す
	res := getTodoByIdResponse{ID: todo.ID, Text: todo.Text}
	return c.JSON(http.StatusOK, res)
}

func (th *todoHandler) ListTodo(c echo.Context) error {

	todos := th.todoUsecase.ListTodo()
	// レスポンスの構造体を定義して返す
	res := make([]listTodoResponse, len(todos))
	if len(todos) == 0 {
		return c.JSON(http.StatusOK, res)
	}
	for i, todo := range todos {
		res[i] = listTodoResponse{ID: todo.ID, Text: todo.Text}
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
	return c.JSON(http.StatusOK, "success")
}

func (th todoHandler) UpdateTodoText(c echo.Context) error {
	r := &updateTodoTextRequest{}
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
	return c.JSON(http.StatusOK, "success")
}
