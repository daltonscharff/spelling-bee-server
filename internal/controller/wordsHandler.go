package controller

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/daltonscharff/spelling-bee-server/internal/database"
)

func viewAllWords(c *fiber.Ctx) error {
	response, err := database.Words.ReadAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(response)
}

func viewWord(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	response, err := database.Words.Read(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(response)
}

func createWord(c *fiber.Ctx) error {
	word := database.Word{}
	if err := c.BodyParser(&word); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	if err := database.Words.Create(&word); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(word)
}

func updateWord(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	word := database.Word{}
	if err := c.BodyParser(&word); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	word.ID = uint(id)
	if err := database.Words.Update(&word); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(word)
}

func deleteAllWords(c *fiber.Ctx) error {
	response, err := database.Words.DeleteAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(response)
}

func deleteWord(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	response, err := database.Words.Delete(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(response)
}
