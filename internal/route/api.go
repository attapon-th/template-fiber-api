// Package route for project api router
package route

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var ()

// Init initializes fiber router
func Init(r fiber.Router) {
	prefix := viper.GetString("prefix")
	prefix = strings.TrimSuffix(prefix, "/")

	log.Debug().Str("path", prefix).Msg("Router initialized")

	r.Use("/", helmet.New(), newCORS(), newLimit(), newAccessLog(), newCompress())

	// controller.NewController(r.Group(prefix))

}

// func createStaticRoute(r fiber.Router) {
// 	r.Use(prefix+"/public", filesystem.New(filesystem.Config{
// 		Root:   http.Dir("./assets"),
// 		Browse: true,
// 		MaxAge: 3600,
// 	}))
// }
