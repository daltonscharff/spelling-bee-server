package updater

import (
	"database/sql"
	"fmt"

	"github.com/daltonscharff/spelling-bee-server/utils/scraper"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func Update(db *sql.DB, gameData *scraper.GameData) error {
	if err := clearTables(db); err != nil {
		panic(err)
	}

	res, err := db.Exec(`INSERT INTO puzzles (date, letters, center) VALUES ($1, $2, $3);`, gameData.Date, pq.Array(gameData.Letters), gameData.CenterLetter)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	return nil
}

func clearTables(db *sql.DB) error {
	// _, err := db.Exec(`DELETE FROM $1;`, table)
	_, err := db.Exec(`DELETE FROM puzzles; DELETE FROM words; DELETE FROM finds;`)
	return err
}
