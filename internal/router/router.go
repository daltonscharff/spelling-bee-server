package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/daltonscharff/spelling-bee-server/internal/handlers"
)

func DefineRoutes(app *fiber.App) error {
	v1 := app.Group("/api/v1")

	words := v1.Group("/words")
	words.Get("/", handlers.ViewAllWords)
	words.Get("/:id", handlers.ViewWord)
	words.Post("/", handlers.CreateWord)
	words.Put("/:id", handlers.UpdateWord)
	words.Delete("/", handlers.DeleteAllWords)
	words.Delete("/:id", handlers.DeleteWord)

	return nil
}
