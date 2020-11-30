package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello from main.go")
	fmt.Println(Scrape().JSON())
}
