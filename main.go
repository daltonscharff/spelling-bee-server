package main

import (
	"github.com/daltonscharff/spelling-bee-server/db"
	"github.com/daltonscharff/spelling-bee-server/utils/game"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	// app := fiber.New()
	// app.Get("/", helloWorld)
	// app.Listen(":3000")

	db, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := game.Refresh(db); err != nil {
		panic(err)
	}
}
