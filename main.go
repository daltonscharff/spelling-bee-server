package main

import (
	"fmt"

	"github.com/daltonscharff/spelling-bee-server/config"
	"github.com/daltonscharff/spelling-bee-server/db"
	"github.com/daltonscharff/spelling-bee-server/utils/updater"
	"github.com/gofiber/fiber/v2"
)

var date string = "2020-12-01"
var letters []byte = []byte{'b', 'i', 'l', 'm', 'o', 't', 'y'}
var center byte = 't'
var words []string = []string{
	"blot",
	"blotto",
	"bolt",
	"boot",
	"booty",
	"bottom",
	"immobility",
	"itty",
	"lilt",
	"limit",
	"lobotomy",
	"loot",
	"lotto",
	"mitt",
	"mobility",
	"molt",
	"moot",
	"motility",
	"motto",
	"obit",
	"omit",
	"till",
	"tilt",
	"toil",
	"toll",
	"tomb",
	"tomboy",
	"tomtit",
	"tool",
	"toot",
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	// app := fiber.New()
	// app.Get("/", helloWorld)
	// app.Listen(":3000")

	configPtr, err := config.Read("./config.yaml")
	if err != nil {
		panic(err)
	}

	config := *configPtr
	fmt.Println(config)

	db, err := db.Connect(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// date, letters, centerLetter, words := scraper.Scrape()

	updater.Update(db, date, letters, center, words)

}
