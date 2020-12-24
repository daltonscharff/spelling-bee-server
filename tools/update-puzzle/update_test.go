package main

import (
	"testing"

	"github.com/daltonscharff/spelling-bee-server/internal/db"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var date string = "2020-12-01"
var letters []string = []string{"b", "i", "l", "m", "o", "t", "y"}
var center string = "t"
var wordMap map[string]analyzedWord = map[string]analyzedWord{
	"immobility": {
		PointValue: 14,
		Definitions: []definition{
			{
				Definition:   "remaining in place",
				PartOfSpeech: "noun",
			},
			{
				Definition:   "the quality of not moving",
				PartOfSpeech: "noun",
			},
		},
	},
}

func beforeEach() *sqlx.DB {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}

	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	return db
}

// func TestClearTables(t *testing.T) {
// 	db := beforeEach()
// 	defer db.Close()

// 	db.MustExec("INSERT INTO puzzles VALUES ($1, $2, $3, $4, $5);", 999, "01-02-2020", pq.Array([]string{"a", "b", "c"}), "a", 100)
// 	db.MustExec("INSERT INTO words VALUES ($1, $2, $3, $4, $5);", 999, "testing", 999, 1, "[]")
// 	db.MustExec("INSERT INTO finds VALUES ($1, $2, $3, $4, $5);", 999, 999, 999, "test player", "2016-06-22 20:44:52.134125-07")

// }

func TestUpdateDB(t *testing.T) {
	db := beforeEach()
	defer db.Close()

	_, err := updateDB(db, date, letters, center, wordMap)
	if err != nil {
		panic(err)
	}
}
