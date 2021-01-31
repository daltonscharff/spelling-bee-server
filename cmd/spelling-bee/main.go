package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/daltonscharff/spelling-bee-server/internal/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	controller, err := api.CreateController()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	api.DefineRoutes(app, controller)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	addr := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	app.Listen(addr)

}
