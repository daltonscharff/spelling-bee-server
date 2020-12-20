package game

import (
	"github.com/daltonscharff/spelling-bee-server/config"
	"github.com/jmoiron/sqlx"
)

func Refresh(db *sqlx.DB, conf config.Config) error {
	date, letters, centerLetter, words := scrape()
	wordMap := analyzeWords(words, conf.RapidAPI.Host, conf.RapidAPI.Key)
	_, err := updateDB(db, date, letters, centerLetter, wordMap)
	if err != nil {
		return err
	}
	return nil
}
