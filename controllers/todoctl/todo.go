package todoctl

import (
	"strconv"

	"github.com/attapon-th/template-fiber-api/services/todosrv"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

var (
	log zerolog.Logger
)

// NewTodoCtl creates a new todo controller
func NewTodoCtl(r fiber.Router) {
	log = zlog.With().Str("ctl", "todo").Logger()

	r.Get("/", gets)         // get list data
	r.Get("/:id", getByID)   // get by id
	r.Post("/", create)      // add data
	r.Put("/:id", update)    // update all data
	r.Patch("/:id", patch)   // update some data
	r.Delete("/:id", delete) // soft delete by default
	log.Info().Msg("New Todo Contoller initialized")
}

func gets(c *fiber.Ctx) error {

	sv := todosrv.NewTodoService(c.UserContext())

	q := c.Queries()
	limit, _ := strconv.ParseInt(c.Query("_limit", "10"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("_page", "1"), 10, 64)

	result := sv.Gets(limit, page, q)
	return c.Status(result.Code).JSON(result)

}
func getByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) != 27 {
		return fiber.NewError(400, "invalid id")
	}

	sv := todosrv.NewTodoService(c.UserContext())
	result := sv.GetByID(id)
	return c.JSON(result)
}

func create(c *fiber.Ctx) error {
	// TODO: implement me
	return nil
}

func update(c *fiber.Ctx) error {
	// TODO: implement me
	return nil
}

func patch(c *fiber.Ctx) error {
	// TODO: implement me
	return nil
}

func delete(c *fiber.Ctx) error {
	// TODO: implement me
	return nil
}
