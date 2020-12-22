package game

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func clearTables(db *sqlx.DB) error {
	if _, err := db.Exec("TRUNCATE TABLE finds CASCADE;"); err != nil {
		return err
	}
	if _, err := db.Exec("TRUNCATE TABLE words CASCADE;"); err != nil {
		return err
	}
	if _, err := db.Exec("TRUNCATE TABLE puzzles CASCADE;"); err != nil {
		return err
	}
	return nil
}

func clearRoomScores(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE rooms SET score = 0;")
	return err
}

func calcMaxScore(wordMap map[string]analyzedWord) (maxScore int) {
	for _, word := range wordMap {
		maxScore += word.PointValue
	}
	return maxScore
}

func writePuzzle(db *sqlx.DB, date string, letters []string, centerLetter string, wordMap map[string]analyzedWord) (int64, error) {
	res := db.QueryRow(`INSERT INTO puzzles (date, letters, center, max_score) VALUES ($1, $2, $3, $4) RETURNING id;`, date, pq.Array(letters), centerLetter, calcMaxScore(wordMap))

	var puzzleID int64
	err := res.Scan(&puzzleID)
	if err != nil {
		return puzzleID, err
	}

	return puzzleID, nil
}

func writeWords(db *sqlx.DB, wordMap map[string]analyzedWord, puzzleID int64) error {
	for word, details := range wordMap {
		definitions, err := json.Marshal(details.Definitions)
		if err != nil {
			definitions = []byte("[]")
		}
		_ = db.MustExec(`INSERT INTO words (word, puzzle_id, point_value, definitions) VALUES ($1, $2, $3, $4);`, word, puzzleID, details.PointValue, string(definitions))
	}
	return nil
}

func updateDB(db *sqlx.DB, date string, letters []string, centerLetter string, wordMap map[string]analyzedWord) (int64, error) {
	if err := clearTables(db); err != nil {
		return -1, err
	}

	if err := clearRoomScores(db); err != nil {
		return -1, err
	}

	puzzleID, err := writePuzzle(db, date, letters, centerLetter, wordMap)
	if err != nil {
		return puzzleID, err
	}

	if err := writeWords(db, wordMap, puzzleID); err != nil {
		return puzzleID, err
	}

	return puzzleID, nil
}
