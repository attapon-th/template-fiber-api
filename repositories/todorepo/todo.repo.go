package todorepo

import (
	"context"

	"github.com/attapon-th/template-fiber-api/models/todomodel"
	"gorm.io/gorm"
)

type TodoRepository struct {
	*gorm.DB
	ctx context.Context
}

func NewTodoRepository(ctx context.Context, db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		ctx: ctx,
		DB:  db,
	}
}

// Initialize Query
func (r *TodoRepository) InitQuery() *TodoRepository {
	return NewTodoRepository(r.ctx, r.DB.WithContext(r.ctx).Model(&todomodel.TodoEntity{}))
}

func (r *TodoRepository) GetByID(result any, id any) error {
	return r.DB.WithContext(r.ctx).Model(&todomodel.TodoEntity{}).Where("id = ?", id).First(result).Error
}
