package game

import (
	"testing"

	"github.com/daltonscharff/spelling-bee-server/config"
	"github.com/daltonscharff/spelling-bee-server/db"
	"github.com/jmoiron/sqlx"
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

func beforeEach() (config.Config, *sqlx.DB) {
	conf, err := config.Read("../../config.yaml")
	if err != nil {
		panic(err)
	}
	db, err := db.Connect(conf)
	if err != nil {
		panic(err)
	}

	return conf, db
}

// func TestClearTables(t *testing.T) {
// 	_, db := beforeEach()
// 	defer db.Close()

// 	db.MustExec("INSERT INTO puzzles VALUES ($1, $2, $3, $4, $5);", 999, "01-02-2020", pq.Array([]string{"a", "b", "c"}), "a", 100)
// 	db.MustExec("INSERT INTO words VALUES ($1, $2, $3, $4, $5);", 999, "testing", 999, 1, "[]")
// 	db.MustExec("INSERT INTO finds VALUES ($1, $2, $3, $4, $5);", 999, 999, 999, "test player", "2016-06-22 20:44:52.134125-07")

// }

func TestUpdateDB(t *testing.T) {
	_, db := beforeEach()
	defer db.Close()

	_, err := updateDB(db, date, letters, center, wordMap)
	if err != nil {
		panic(err)
	}
}
