package request

type PostTodoRequest struct {
	Text string `json:"text" validate:"required"`
}

type UpdateTodoTextRequest struct {
	Text string `json:"text" validate:"required"`
}

type TodoRequest interface {
	PostTodo() (*PostTodoRequest, error)
	UpdateTodoText() (*UpdateTodoTextRequest, error)
}

type todoRequest struct {
}

func NewTodoRequest() TodoRequest {
	return &todoRequest{}
}

func (tr *todoRequest) PostTodo() (*PostTodoRequest, error) {
	return &PostTodoRequest{}, nil
}

func (tr *todoRequest) UpdateTodoText() (*UpdateTodoTextRequest, error) {
	return &UpdateTodoTextRequest{}, nil
}
