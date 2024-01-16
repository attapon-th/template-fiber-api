package todosrv

import (
	"github.com/attapon-th/template-fiber-api/models"
	"github.com/attapon-th/template-fiber-api/models/todomodel"
	"github.com/rs/zerolog/log"
)

// func (s *TodoService) filters(filters map[string]string) {
// 	if filters == nil {
// 		return
// 	}
// 	r := s.todoRepo
// 	fields := s.todoRepo.Fields
// 	for key, value := range filters {
// 		f := key
// 		if lo.Contains(fields, f) {
// 			// = or like
// 			if strings.Contains(value, "*") {
// 				v := strings.ReplaceAll(value, "*", "%")
// 				r.Where(f+" LIKE ?", v)
// 			} else {
// 				r.Where(f+" = ?", value)
// 			}
// 		}
// 	}
// }
// func (s *TodoService) Gets(limit, page int64, filters map[string]string) *schemas.Todos {
// 	r := s.todoRepo
// 	result := schemas.NewTodos()
// 	result.Pagination = schemas.NewPagination(page, limit)
// 	offset := int64(0)
// 	limit, offset = result.Pagination.GetLimitOffset()
// 	// data := []models.Todo{}
// 	r.Context(s.ctx).Limit(int(limit)).Offset(int(offset))
// 	s.filters(filters)
// 	r.Find(&result.Data)
// 	if err := r.Err(); err != nil {
// 		log.Error().Err(err).Msg("get todo error")
// 		result.Message = "Get Todo Error"
// 		return result
// 	}
// 	log.Debug().Interface("pagination", result.Pagination).Msg("pagination")
// 	result.Pagination.SetTotalRecord(r.Count())
// 	result.APIResponse = *schemas.NewAPIResponse(200, "OK")
// 	return result
// }

func (s *TodoService) GetByID(id string) *todomodel.TodoOne {
	result := todomodel.NewTodoOne()

	err := s.repo.GetByID(result.Data, id)
	if err != nil {
		log.Error().Err(err).Msg("get todo error, " + err.Error())
		result.APIResponse = models.NewAPIResponseError(err)
		result.Data = nil
	}
	return result
}
