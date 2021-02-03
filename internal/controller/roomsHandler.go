package controller

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/daltonscharff/spelling-bee-server/internal/database"
)

func ViewAllRooms(c *fiber.Ctx) error {
	response, err := database.Rooms.ReadAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(response)
}

func ViewRoom(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	response, err := database.Rooms.Read(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(response)
}

func CreateRoom(c *fiber.Ctx) error {
	room := database.Room{}
	if err := c.BodyParser(&room); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	if err := database.Rooms.Create(&room); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(room)
}

func UpdateRoom(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	room := database.Room{}
	if err := c.BodyParser(&room); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	room.ID = uint(id)
	if err := database.Rooms.Update(&room); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(room)
}

func DeleteRoom(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	response, err := database.Rooms.Delete(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(response)
}
