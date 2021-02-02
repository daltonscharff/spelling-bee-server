package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/daltonscharff/spelling-bee-server/internal/handler"
)

func DefineRoutes(app *fiber.App) error {
	v1 := app.Group("/api/v1")

	words := v1.Group("/words")
	words.Get("/", handler.ViewAllWords)
	words.Get("/:id", handler.ViewWord)
	words.Post("/", handler.CreateWord)
	words.Put("/:id", handler.UpdateWord)
	words.Delete("/", handler.DeleteAllWords)
	words.Delete("/:id", handler.DeleteWord)

	return nil
}
