package repositories

import (
	"context"

	"gorm.io/gorm"
)

// TableModel get table name interface of gorm
type TableModel interface {
	TableName() string
}

type Repository struct {
	*gorm.DB
	model TableModel
}

// NewRepository create new repository for Table Model
func NewRepository(db *gorm.DB, model TableModel) *Repository {
	return &Repository{
		DB:    db,
		model: model,
	}
}

func (t *Repository) Model() TableModel {
	return t.model
}

func (t *Repository) TableName() string {
	return t.model.TableName()
}

func (t *Repository) Context(ctx context.Context) *Repository {
	t.DB = t.DB.WithContext(ctx)
	return t
}
