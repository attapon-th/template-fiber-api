// Package models create model
package models

import (
	"github.com/attapon-th/null"
	"github.com/segmentio/ksuid"
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
