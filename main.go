package main

import (
	"github.com/daltonscharff/spelling-bee-server/config"
	"github.com/daltonscharff/spelling-bee-server/db"
	"github.com/daltonscharff/spelling-bee-server/utils/game"
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	// app := fiber.New()
	// app.Get("/", helloWorld)
	// app.Listen(":3000")

	config, err := config.Read("./config.yaml")
	if err != nil {
		panic(err)
	}

	db, err := db.Connect(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := game.Refresh(db, config); err != nil {
		panic(err)
	}
}
