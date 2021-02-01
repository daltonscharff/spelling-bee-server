package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/daltonscharff/spelling-bee-server/internal/handler"
)

func DefineRoutes(app *fiber.App) error {
	v1 := app.Group("/api/v1")

	words := v1.Group("/words")
	words.Get("/", handler.ViewAllWords)

	return nil
}
