package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/daltonscharff/spelling-bee-server/internal/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	r := router.New()
	addr := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))

	log.Printf("Server started: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
