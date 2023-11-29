package repositories

import (
	"strings"

	"github.com/attapon-th/template-fiber-api/models"
	"github.com/attapon-th/template-fiber-api/pkg"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const todoFields string = "id,name,status_id,comment,complated_at,tags"

// TodoRepository todo repository
type TodoRepository struct {
	*Repository[models.Todo]
	Fields []string
}

func NewTodoRepository() *TodoRepository {
	td := &TodoRepository{}
	db, err := pkg.ConnectPostgreSQL(viper.GetString("DB_DSN"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	td.Repository = NewRepository(db, models.Todo{})
	td.Fields = strings.Split(todoFields, ",")
	return td
}
