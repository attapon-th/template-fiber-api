// Package entities is mapper for database or etc.
package entities

import (
	"time"

	"github.com/attapon-th/null"
	"github.com/segmentio/ksuid"
	"gorm.io/plugin/soft_delete"
)

// ModelID primary key
type ModelID struct {
	ID       ksuid.KSUID `json:"id"`
	LastUser null.String `json:"last_user"`
}

// NewID create new ID
func (m *ModelID) NewID() ksuid.KSUID {
	m.ID = ksuid.New()
	return m.ID
}

// EntityRecordLog record logger
type EntityRecordLog struct {
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt time.Time
	IsDel     soft_delete.DeletedAt `gorm:"index;softDelete:flag,DeletedAtField:DeletedAt"` // use `1` `0`
}
