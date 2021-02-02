package controller

import (
	"github.com/gofiber/fiber/v2"
)

func DefineRoutes(app *fiber.App) error {
	v1 := app.Group("/api/v1")

	words := v1.Group("/words")
	words.Get("/", ViewAllWords)
	words.Get("/:id", ViewWord)
	words.Post("/", CreateWord)
	words.Put("/:id", UpdateWord)
	words.Delete("/", DeleteAllWords)
	words.Delete("/:id", DeleteWord)

	return nil
}
