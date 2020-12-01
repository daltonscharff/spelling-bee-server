package main

import (
	"fmt"

	"github.com/daltonscharff/spelling-bee-server/utils/scraper/scraper"
)

func main() {
	fmt.Println(scraper.Scrape().JSON())
}
