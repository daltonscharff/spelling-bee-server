package controller

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/daltonscharff/spelling-bee-server/internal/database"
)

func ViewAllPuzzles(c *fiber.Ctx) error {
	response, err := database.Puzzles.ReadAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(response)
}

func ViewPuzzle(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	response, err := database.Puzzles.Read(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(response)
}

func CreatePuzzle(c *fiber.Ctx) error {
	puzzle := database.Puzzle{}
	if err := c.BodyParser(&puzzle); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	if err := database.Puzzles.Create(&puzzle); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(puzzle)
}

func UpdatePuzzle(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	puzzle := database.Puzzle{}
	if err := c.BodyParser(&puzzle); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	puzzle.ID = uint(id)
	if err := database.Puzzles.Update(&puzzle); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(puzzle)
}

func DeletePuzzle(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	response, err := database.Puzzles.Delete(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(response)
}
