package repositories

import "gorm.io/gorm"

type repo struct {
	*gorm.DB
}

// NewRepository create new repository
func NewRepository() {

}
