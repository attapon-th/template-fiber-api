package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var (
	prefix = "/api"
)

// NewRouters initializes fiber router
func NewRouters(r fiber.Router) {
	prefix = viper.GetString("prefix")

	// r.Use("/", middlewares.NewHelmet(),
	// 	middlewares.NewCORS(),
	// 	middlewares.NewLimit(),
	// 	middlewares.NewAccessLog(),
	// 	middlewares.NewCompress(),
	// )

	createStaticRoute(r)
	createRestAPIRouter(r)
	cretateSwagerRoute(r)
}
