package todosrv

import (
	"time"

	"github.com/attapon-th/null"
	"github.com/attapon-th/template-fiber-api/models/todomodel"
	"github.com/rs/zerolog/log"
)

func (s *TodoService) Create(data *todomodel.TodoItem) *todomodel.TodoOne {
	result := todomodel.NewTodoOne()

	// convert schema to model
	md := todomodel.TodoEntity{
		Name:        data.Name,
		Comment:     data.Comment,
		StatusID:    data.StatusID,
		ComplatedAt: null.NewTime(time.Time{}, false),
		Tags:        data.Tags,
	}

	err := s.repo.WithContext(s.ctx).Model(md).Create(md).Error
	if err != nil {
		log.Error().Err(err).Msg("create todo error")
		// result.
		result.SetError(err)
		result.Data = nil
		return result
	}

	result = s.GetByID(md.ID.String())
	result.SetResponse(true, "Create Todo Success")
	return result
}
