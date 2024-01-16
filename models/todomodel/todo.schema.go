package todomodel

import (
	"github.com/attapon-th/null"
	"github.com/attapon-th/template-fiber-api/models"
)

// Todo Default data todo response
type TodoResponse struct {
	ID          null.String `json:"id" `
	Name        null.String `json:"name" `
	StatusID    null.Int    `json:"status_id" `
	Comment     null.String `json:"comment"`
	ComplatedAt null.Time   `json:"complated_at"`
	Tags        null.String `json:"tags" `
	UpdatedAt   null.Time   `json:"updated_at" `
} // @name	Todo

// TodoItem Default data todo for Create, Update
type TodoItem struct {
	Name        null.String `json:"name" `
	StatusID    null.Int    `json:"status_id" `
	Comment     null.String `json:"comment"`
	ComplatedAt null.Time   `json:"complated_at"`
	Tags        null.String `json:"tags" `
} // @name TodoItem

type Todos models.APIResponsePagination[TodoResponse]

// NewAPIResponseTodos create new list of todo
func NewTodos() *Todos {
	return &Todos{
		APIResponse: models.NewAPIResponse(true, "Success"),
		Data:        []TodoResponse{},
		Pagination:  models.NewPagination(1, 10),
	}
}

type TodoOne models.APIResponseOne[TodoResponse]

func NewTodoOne() *TodoOne {
	return &TodoOne{
		APIResponse: models.NewAPIResponse(true, "Success"),
		Data:        &TodoResponse{},
	}
}
