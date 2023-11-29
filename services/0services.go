package services

import (
	"github.com/attapon-th/template-fiber-api/pkg"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	// DefaultDB is default database connection for service
	DefaultDB *gorm.DB
)

// NewServices initialize global database and configuration
func NewServices() {

	db, err := pkg.ConnectPostgreSQL(viper.GetString("DB_DSN"))
	if err != nil {
		panic(err)
	}
	DefaultDB = db

	if viper.GetBool("dev") {
		db = db.Debug()
	}
	DefaultDB = db
}
