package fooctl

import "github.com/gofiber/fiber/v2"

func NewFooCtl(r fiber.Router) {
	r.Get("/", getFoo)
}

func getFoo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "foo",
		"ok":      true,
	})
}
