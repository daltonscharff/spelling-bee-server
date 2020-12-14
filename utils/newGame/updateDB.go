package newGame

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func Update(db *sqlx.DB, date string, letters []string, centerLetter string, words []string) error {

	// clear tables
	db.MustExec("DELETE FROM puzzles;")
	db.MustExec("DELETE FROM words;")
	db.MustExec("DELETE FROM finds;")

	fmt.Println(date)
	fmt.Println(letters)
	fmt.Println(centerLetter)

	res := db.MustExec(`INSERT INTO puzzles (date, letters, center) VALUES ($1, $2, $3);`, date, pq.Array(letters), centerLetter)

	fmt.Println(res)

	return nil
}

// func getDefinitions (word string) (error)
