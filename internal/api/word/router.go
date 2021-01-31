package word

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ViewAllWords(c *Controller) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		response, err := c.ReadAll()
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		return ctx.JSON(response)
	}
}
