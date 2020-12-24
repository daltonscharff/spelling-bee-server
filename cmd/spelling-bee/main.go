package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	app := fiber.New()
	app.Get("/", helloWorld)
	app.Listen(":3000")
}
