package controller

import (
	"github.com/gofiber/fiber/v2"
)

func DefineRoutes(app *fiber.App) error {
	v1 := app.Group("/api/v1")

	words := v1.Group("/words")
	words.Get("/", viewAllWords)
	words.Get("/:id", viewWord)
	words.Post("/", createWord)
	words.Put("/:id", updateWord)
	words.Delete("/", deleteAllWords)
	words.Delete("/:id", deleteWord)

	rooms := v1.Group("/rooms")
	rooms.Get("/", viewAllRooms)
	rooms.Get("/:id", viewRoom)
	rooms.Post("/", createRoom)
	rooms.Put("/:id", updateRoom)
	rooms.Delete("/:id", deleteRoom)

	records := v1.Group("/records")
	records.Get("/", viewAllRecords)
	records.Get("/:id", viewRecord)
	records.Post("/", createRecord)
	records.Put("/:id", updateRecord)
	records.Delete("/:id", deleteRecord)

	puzzles := v1.Group("/puzzles")
	puzzles.Get("/", viewAllPuzzles)
	puzzles.Get("/:id", viewPuzzle)
	puzzles.Post("/", createPuzzle)
	puzzles.Put("/:id", updatePuzzle)
	puzzles.Delete("/:id", deletePuzzle)

	return nil
}
