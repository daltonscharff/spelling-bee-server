package game

import (
	"github.com/daltonscharff/spelling-bee-server/config"
	"github.com/jmoiron/sqlx"
)

func Refresh(db *sqlx.DB, conf config.Config) {
	date, letters, centerLetter, words := scrape()
	wordMap := analyzeWords(words, conf.RapidAPI.Host, conf.RapidAPI.Key)
	puzzleID, err := updateDB(date, letters, centerLetter, wordMap)
}
