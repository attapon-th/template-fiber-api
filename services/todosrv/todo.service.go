package todosrv

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type TodoService struct {
	log zerolog.Logger
	id  string

	filters []string
}

func NewTodoService() *TodoService {
	return &TodoService{
		log: log.Logger,
	}
}

func (s *TodoService) Filter(key string, value any, condition ...string) {

}
func (s *TodoService) Gets() {

}
