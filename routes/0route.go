package routes

import (
	"strings"

	"github.com/attapon-th/template-fiber-api/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var (
	prefix = "/api"
)

// NewRouters initializes fiber router
func NewRouters(r fiber.Router) {
	prefix = viper.GetString("prefix")
	prefix = strings.TrimSuffix(prefix, "/")

	r.Use("/", middlewares.NewHelmet(),
		middlewares.NewCORS(),
		middlewares.NewLimit(),
		middlewares.NewAccessLog(),
		middlewares.NewCompress(),
	)

	log.Debug().Str("path", prefix+"/public").Msg("Router Public")
	createStaticRoute(r.Group(prefix + "/public"))

	log.Debug().Str("path", prefix+"/api").Msg("Router RestAPI")
	createRestAPIRouter(r.Group(prefix + "/api"))

	log.Debug().Str("path", prefix+"/swagger").Msg("Router swagger")
	cretateSwagerRoute(r.Group(prefix + "/swagger"))
}
