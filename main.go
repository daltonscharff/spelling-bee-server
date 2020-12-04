package main

import (
	"fmt"

	"github.com/daltonscharff/spelling-bee-server/config"
	"github.com/daltonscharff/spelling-bee-server/db"
	"github.com/daltonscharff/spelling-bee-server/utils/scraper"
	"github.com/daltonscharff/spelling-bee-server/utils/updater"
	"github.com/gofiber/fiber/v2"
)

var gameData = &scraper.GameData{
	Date: "2020-12-04",
	Words: []string{
		"ardor",
		"award",
		"awkward",
		"daddy",
		"dark",
		"daywork",
		"dodo",
		"doodad",
		"door",
		"doorway",
		"dooryard",
		"dork",
		"dorky",
		"dory",
		"dowdy",
		"dowry",
		"draw",
		"dray",
		"dryad",
		"dyad",
		"odor",
		"radar",
		"road",
		"roadway",
		"roadwork",
		"rood",
		"rowdy",
		"ward",
		"wayward",
		"wood",
		"woodwork",
		"woody",
		"word",
		"wordy",
		"workaday",
		"workday",
		"yard",
		"yardwork",
	},
	Letters:      []string{"a", "r", "d", "o", "w", "k", "y"},
	CenterLetter: "d",
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

	// gameData := scraper.Scrape()

	updater.Update(db, gameData)

}
