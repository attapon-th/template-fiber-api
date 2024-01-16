package routes

import (
	"strings"

	"github.com/attapon-th/template-fiber-api/controllers/pingctl"
	"github.com/attapon-th/template-fiber-api/controllers/todoctl"
	"github.com/attapon-th/template-fiber-api/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// NewRouters initializes fiber router
func NewRouters(r fiber.Router) {
	prefix := viper.GetString("prefix")
	prefix = strings.TrimSuffix(prefix, "/")

	defaultMiddleware := []any{
		"/",
		middlewares.NewHelmet(),
		middlewares.NewCORS(),
		middlewares.NewLimit(),
		middlewares.NewAccessLog(),
		middlewares.NewCompress(),
	}

	r.Use(defaultMiddleware...)

	// Group Route
	routeBaseGroup := r.Group(prefix).Group
	apiPathPrefix := prefix + "/api/v1"
	apiRouteGroup := r.Group(apiPathPrefix).Group

	// Static route
	newStaticRoute(routeBaseGroup("/public"))

	// Swagger
	newSwagerRoute(routeBaseGroup("/swagger"), prefix)

	// API Endpoint
	pingctl.NewPings(routeBaseGroup("/ping"))
	todoctl.NewTodos(apiRouteGroup("/todos"))
}
