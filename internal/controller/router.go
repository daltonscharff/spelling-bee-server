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

	rooms := v1.Group("/rooms")
	rooms.Get("/", ViewAllRooms)
	rooms.Get("/:id", ViewRoom)
	rooms.Post("/", CreateRoom)
	rooms.Put("/:id", UpdateRoom)
	rooms.Delete("/:id", DeleteRoom)

	records := v1.Group("/records")
	records.Get("/", ViewAllRecords)
	records.Get("/:id", ViewRecord)
	records.Post("/", CreateRecord)
	records.Put("/:id", UpdateRecord)
	records.Delete("/:id", DeleteRecord)

	puzzles := v1.Group("/puzzles")
	puzzles.Get("/", ViewAllPuzzles)
	puzzles.Get("/:id", ViewPuzzle)
	puzzles.Post("/", CreatePuzzle)
	puzzles.Put("/:id", UpdatePuzzle)
	puzzles.Delete("/:id", DeletePuzzle)

	return nil
}
