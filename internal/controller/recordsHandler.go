package controller

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/daltonscharff/spelling-bee-server/internal/database"
)

func viewAllRecords(c *fiber.Ctx) error {
	response, err := database.Records.ReadAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(response)
}

func viewRecord(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	response, err := database.Records.Read(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(response)
}

func createRecord(c *fiber.Ctx) error {
	record := database.Record{}
	if err := c.BodyParser(&record); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	if err := database.Records.Create(&record); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(record)
}

func updateRecord(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	record := database.Record{}
	if err := c.BodyParser(&record); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	record.ID = uint(id)
	if err := database.Records.Update(&record); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(record)
}

func deleteRecord(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	response, err := database.Records.Delete(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(response)
}
