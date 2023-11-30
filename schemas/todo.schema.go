package schemas

import "github.com/attapon-th/null"

// Todo Default data todo response
type Todo struct {
	ID        null.String `json:"id" `
	UpdatedAt null.Time   `json:"updated_at" `
	TodoItem
}

// TodoItem Default data todo for Create, Update
type TodoItem struct {
	Name        null.String `json:"name" `
	StatusID    null.Int    `json:"status_id" `
	Comment     null.String `json:"comment"`
	ComplatedAt null.Time   `json:"complated_at"`
	Tags        null.String `json:"tags" `
}

// Todos schemage for list of todo with pagination response
type Todos GetsAPIResponse[Todo]

// NewTodos create new list of todo
func NewTodos() *Todos {
	return &Todos{
		APIResponse: *NewAPIResponse(500, "Internal Server Error"),
		Data:        []Todo{},
		Pagination:  NewPagination(1, 10),
	}
}

// TodoOne schemage for todo response
type TodoOne GetOneAPIResponse[Todo]

// NewTodoOne create new todo one
func NewTodoOne() *TodoOne {
	return &TodoOne{
		APIResponse: *NewAPIResponse(500, "Internal Server Error"),
		Data:        Todo{},
	}
}
