package todo_dto

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description,omitempty"`
}

type CreateTodoDto struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description,omitempty"`
}
