package game

import (
	"github.com/jmoiron/sqlx"
)

func Refresh(db *sqlx.DB) error {
	date, letters, centerLetter, words := scrape()
	wordMap := analyzeWords(words)
	_, err := updateDB(db, date, letters, centerLetter, wordMap)
	if err != nil {
		return err
	}
	return nil
}
