package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/daltonscharff/spelling-bee-server/internal/api/word"
)

func DefineRoutes(app *fiber.App, c *Controller) error {

	v1 := app.Group("/api/v1")

	words := v1.Group("/words")
	words.Get("/", word.ViewAllWords(c.Words))

	return nil
}
