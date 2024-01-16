package todosrv

import (
	"context"

	"github.com/attapon-th/template-fiber-api/repositories/todorepo"
	"gorm.io/gorm"
)

type TodoService struct {
	ctx  context.Context
	repo *todorepo.TodoRepository
}

func NewTodoService(ctx context.Context, db *gorm.DB) *TodoService {
	return &TodoService{
		ctx:  ctx,
		repo: todorepo.NewTodoRepository(ctx, db),
	}
}

// func (s *TodoService) Update(id string, data *schemas.TodoItem) *schemas.TodoOne {
// 	r := s.todoRepo

// 	// check id if exists and no error or deleted
// 	result := s.GetByID(id)
// 	if result.Code == 404 {
// 		return result
// 	} else if result.Code != 200 {
// 		return result
// 	}

// 	md := models.Todo{
// 		Name:        data.Name,
// 		Comment:     data.Comment,
// 		StatusID:    data.StatusID,
// 		ComplatedAt: data.ComplatedAt,
// 		Tags:        data.Tags,
// 	}
// 	if err := md.ID.Scan(id); err != nil {
// 		result.APIResponse = *schemas.NewAPIResponse(500, "Update Todo Error, id not valid")
// 		return result
// 	}

// 	err := r.Context(s.ctx).Where("id = ?", id).Updates(&md)
// 	if err != nil {
// 		result.APIResponse = *schemas.NewAPIResponse(500, "Update Todo Error, "+err.Error())
// 		return result
// 	}

// 	result = s.GetByID(id)
// 	result.Message = "Update Todo Success"
// 	return result

// }

// func (s *TodoService) Delete(id string) *schemas.TodoOne {
// 	r := s.todoRepo
// 	result := schemas.NewTodoOne()
// 	if len(id) != 27 {
// 		result.APIResponse = *schemas.NewAPIResponse(400, "invalid id")
// 		return result
// 	}
// 	err := r.Context(s.ctx).Where("id = ?", id).Delete()
// 	if err != nil {
// 		log.Error().Err(err).Msg("delete todo error")
// 		result.APIResponse = *schemas.NewAPIResponse(500, "Delete Todo Error")
// 		return result
// 	}
// 	result.Data.ID = null.StringFrom(id)
// 	result.APIResponse = *schemas.NewAPIResponse(200, "Delete Todo Success")
// 	return result
// }
