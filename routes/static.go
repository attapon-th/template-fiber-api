package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func createStaticRoute(r fiber.Router) {
	r.Use(prefix+"/public", filesystem.New(filesystem.Config{
		Root:   http.Dir("./assets"),
		Browse: false,
		MaxAge: 3600,
	}))
}
