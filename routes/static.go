package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func newStaticRoute(r fiber.Router) {
	r.Use("/", filesystem.New(filesystem.Config{
		Root:   http.Dir("./assets"),
		Browse: false,
		MaxAge: 3600,
	}))
}
