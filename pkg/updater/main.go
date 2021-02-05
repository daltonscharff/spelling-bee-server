package main

import (
	"github.com/daltonscharff/spelling-bee-server/internal/db"
	"github.com/jmoiron/sqlx"

	"github.com/joho/godotenv"
)

func updatePuzzle(db *sqlx.DB) error {
	date, letters, centerLetter, words := scrape()
	wordMap := analyzeWords(words)
	_, err := updateDB(db, date, letters, centerLetter, wordMap)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := updatePuzzle(db); err != nil {
		panic(err)
	}
}
