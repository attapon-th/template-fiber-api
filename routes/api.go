// Package routes for project api router
package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Init initializes fiber router
func createRestAPIRouter(r fiber.Router) {
	prefix := viper.GetString("prefix")
	prefix = strings.TrimSuffix(prefix, "/")

	log.Debug().Str("path", prefix).Msg("Router RestAPI initialized")

}
