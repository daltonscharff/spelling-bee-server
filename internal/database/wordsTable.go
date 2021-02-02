package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Word struct {
	ID           uint   `db:"id" json:"id"`
	Word         string `db:"word" json:"word"`
	PuzzleID     uint   `db:"puzzle_id" json:"puzzleId"`
	PointValue   uint   `db:"point_value" json:"pointValue"`
	Definition   string `db:"definition" json:"definition"`
	PartOfSpeech string `db:"part_of_speech" json:"partOfSpeech"`
	Synonym      string `db:"synonym" json:"synonym"`
}

type WordsTable struct {
	*sqlx.DB
}

func (t *WordsTable) Read(id uint64) (Word, error) {
	var w Word
	if err := t.Get(&w, `SELECT * FROM words WHERE id = $1;`, id); err != nil {
		return Word{}, fmt.Errorf("error getting word: %w", err)
	}
	return w, nil
}

func (t *WordsTable) ReadAll() ([]Word, error) {
	var ww []Word
	if err := t.Select(&ww, `SELECT * FROM words;`); err != nil {
		return []Word{}, fmt.Errorf("error getting words: %w", err)
	}
	return ww, nil
}

func (t *WordsTable) Create(w *Word) error {
	if err := t.Get(w, `INSERT INTO words (word, puzzle_id, point_value, definition, part_of_speech, synonym) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`,
		w.Word,
		w.PuzzleID,
		w.PointValue,
		w.Definition,
		w.PartOfSpeech,
		w.Synonym); err != nil {
		return fmt.Errorf("error creating word: %w", err)
	}
	return nil
}

func (t *WordsTable) Update(w *Word) error {
	if err := t.Get(w, `UPDATE words SET word = $1, puzzle_id = $2, point_value = $3, definition = $4, part_of_speech = $5, synonym = $6 WHERE id = $7 RETURNING *;`,
		w.Word,
		w.PuzzleID,
		w.PointValue,
		w.Definition,
		w.PartOfSpeech,
		w.Synonym,
		w.ID); err != nil {
		return fmt.Errorf("error updating word: %w", err)
	}
	return nil
}

func (t *WordsTable) Delete(id uint64) (Word, error) {
	var w Word
	if err := t.Get(&w, `DELETE FROM words WHERE id = $1 RETURNING *;`, id); err != nil {
		return Word{}, fmt.Errorf("error deleting word: %w", err)
	}
	return w, nil
}

func (t *WordsTable) DeleteAll() ([]Word, error) {
	var ww []Word
	if err := t.Select(&ww, `DELETE FROM words RETURNING *;`); err != nil {
		return []Word{}, fmt.Errorf("error deleting word: %w", err)
	}
	return ww, nil
}
