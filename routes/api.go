// Package routes for project api router
package routes

import (
	"github.com/attapon-th/template-fiber-api/controllers/todoctl"
	"github.com/gofiber/fiber/v2"
)

// Init initializes fiber router
func createRestAPIRouter(r fiber.Router) {
	apiV1 := r.Group("/v1")
	todoctl.NewTodoCtl(apiV1.Group("/todos")) // prefix+/api/v1/todos
}
