package game

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func clearTables(db *sqlx.DB) error {
	if _, err := db.Exec("DELETE FROM puzzles;"); err != nil {
		return err
	}
	if _, err := db.Exec("DELETE FROM words;"); err != nil {
		return err
	}
	if _, err := db.Exec("DELETE FROM finds;"); err != nil {
		return err
	}
	return nil
}

func updateDB(db *sqlx.DB, date string, letters []string, centerLetter string, wordMap map[string]analyzedWord) (int, error) {

	if err := clearTables(db); err != nil {
		return -1, err
	}

	res := db.MustExec(`INSERT INTO puzzles (date, letters, center) VALUES ($1, $2, $3);`, date, pq.Array(letters), centerLetter)

	fmt.Println(res)

	return -1, nil
}

// func getDefinitions (word string) (error)
