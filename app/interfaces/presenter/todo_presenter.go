package presenter

import (
	"github.com/sako0/todo-api/app/domain/model"
	"github.com/sako0/todo-api/app/usecase"
)

type TodoPresenter interface {
	PostTodoResponse(err error) (interface{}, error)
	GetTodoByIdResponse(todo model.Todo, err error) (interface{}, error)
	ListTodoResponse(todos []model.Todo, err error) (interface{}, error)
	DeleteTodoResponse(err error) (interface{}, error)
	UpdateTodoTextResponse(err error) (interface{}, error)
}

type todoPresenter struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoPresenter(usecase usecase.TodoUsecase) TodoPresenter {
	return &todoPresenter{todoUsecase: usecase}
}

func (tp *todoPresenter) PostTodoResponse(err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}
	return map[string]string{"message": "成功しました！"}, nil
}

func (tp *todoPresenter) GetTodoByIdResponse(todo model.Todo, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}
	return map[string]string{"id": string(rune(todo.ID)), "text": todo.Text}, nil
}

type listTodoResponse struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

func (tp *todoPresenter) ListTodoResponse(todos []model.Todo, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	res := []*listTodoResponse{}
	for _, t := range todos {
		res = append(res, &listTodoResponse{ID: t.ID, Text: t.Text})
	}
	return res, nil
}

func (tp *todoPresenter) DeleteTodoResponse(err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}
	return map[string]string{"message": "成功しました！"}, nil
}

func (tp *todoPresenter) UpdateTodoTextResponse(err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}
	return map[string]string{"message": "成功しました！"}, nil
}
