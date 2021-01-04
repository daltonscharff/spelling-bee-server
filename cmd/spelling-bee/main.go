package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/daltonscharff/spelling-bee-server/internal/api"
	"github.com/daltonscharff/spelling-bee-server/internal/postgres"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	store, err := postgres.NewStore(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatal(err)
	}

	h := api.NewHandler(*store)
	addr := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))

	log.Printf("Server started: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, h))
}
