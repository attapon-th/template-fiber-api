// Package controller create controller for api endpoint
package controller

import (
	"github.com/gofiber/fiber/v2"
)

var ()

// NewController creates a new controller for api endpoint
func NewController(r fiber.Router) {
	r.Get("/ping", ping)
}

func ping(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
