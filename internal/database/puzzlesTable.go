package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Puzzle struct {
	ID           uint      `db:"id" json:"id"`
	Date         time.Time `db:"date" json:"date"`
	Letters      []string  `db:"letters" json:"letters"`
	CenterLetter string    `db:"center_letter" json:"centerLetter"`
	MaxScore     uint      `db:"max_score" json:"maxScore"`
}

type PuzzlesTable struct {
	*sqlx.DB
}

func (t *PuzzlesTable) Read(id uint64) (Puzzle, error) {
	var p Puzzle
	if err := t.Get(&p, `SELECT * FROM puzzles WHERE id = $1;`, id); err != nil {
		return Puzzle{}, fmt.Errorf("error getting puzzle: %w", err)
	}
	return p, nil
}

func (t *PuzzlesTable) ReadAll() ([]Puzzle, error) {
	var pp []Puzzle
	if err := t.Select(&pp, `SELECT * FROM puzzles;`); err != nil {
		return []Puzzle{}, fmt.Errorf("error getting puzzles: %w", err)
	}
	return pp, nil
}

func (t *PuzzlesTable) Create(p *Puzzle) error {
	if err := t.Get(p, `INSERT INTO puzzles (date, letters, center_letter, max_score) VALUES ($1, $2, $3, $4) RETURNING *`,
		p.Date,
		p.Letters,
		p.CenterLetter,
		p.MaxScore); err != nil {
		return fmt.Errorf("error creating puzzle: %w", err)
	}
	return nil
}

func (t *PuzzlesTable) Update(p *Puzzle) error {
	if err := t.Get(p, `UPDATE puzzles SET date = $1, letters = $2, center_letter = $3, max_score = $4 RETURNING *;`,
		p.Date,
		p.Letters,
		p.CenterLetter,
		p.MaxScore); err != nil {
		return fmt.Errorf("error updating puzzle: %w", err)
	}
	return nil
}

func (t *PuzzlesTable) Delete(id uint64) (Puzzle, error) {
	var p Puzzle
	if err := t.Get(p, `DELETE FROM puzzles WHERE id = $1 RETURNING *;`, id); err != nil {
		return Puzzle{}, fmt.Errorf("error deleting puzzle: %w", err)
	}
	return p, nil
}
