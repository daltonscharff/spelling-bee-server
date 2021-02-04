package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Puzzle struct {
	ID           uint           `db:"id" json:"id"`
	Date         time.Time      `db:"date" json:"date"`
	Letters      pq.StringArray `db:"letters" json:"letters"`
	CenterLetter string         `db:"center_letter" json:"centerLetter"`
	MaxScore     uint           `db:"max_score" json:"maxScore"`
}

type puzzlesTable struct {
	*sqlx.DB
}

func (t puzzlesTable) InitTable() error {
	_, err := t.Exec(`CREATE TABLE IF NOT EXISTS puzzles
		(
				id serial NOT NULL,
				date date NOT NULL,
				letters character(1)[] NOT NULL,
				center_letter character(1) NOT NULL,
				max_score integer NOT NULL,
				CONSTRAINT puzzles_pkey PRIMARY KEY (id)
		);`)
	return err
}

func (t *puzzlesTable) Read(id uint64) (Puzzle, error) {
	var p Puzzle
	if err := t.Get(&p, `SELECT * FROM puzzles WHERE id = $1;`, id); err != nil {
		return Puzzle{}, fmt.Errorf("error getting puzzle: %w", err)
	}
	return p, nil
}

func (t *puzzlesTable) ReadAll() ([]Puzzle, error) {
	var pp []Puzzle
	if err := t.Select(&pp, `SELECT * FROM puzzles;`); err != nil {
		return []Puzzle{}, fmt.Errorf("error getting puzzles: %w", err)
	}
	return pp, nil
}

func (t *puzzlesTable) Create(p *Puzzle) error {
	if err := t.Get(p, `INSERT INTO puzzles (date, letters, center_letter, max_score) VALUES ($1, $2, $3, $4) RETURNING *`,
		p.Date,
		p.Letters,
		p.CenterLetter,
		p.MaxScore); err != nil {
		return fmt.Errorf("error creating puzzle: %w", err)
	}
	return nil
}

func (t *puzzlesTable) Update(p *Puzzle) error {
	if err := t.Get(p, `UPDATE puzzles SET date = $1, letters = $2, center_letter = $3, max_score = $4 RETURNING *;`,
		p.Date,
		p.Letters,
		p.CenterLetter,
		p.MaxScore); err != nil {
		return fmt.Errorf("error updating puzzle: %w", err)
	}
	return nil
}

func (t *puzzlesTable) Delete(id uint64) (Puzzle, error) {
	var p Puzzle
	if err := t.Get(p, `DELETE FROM puzzles WHERE id = $1 RETURNING *;`, id); err != nil {
		return Puzzle{}, fmt.Errorf("error deleting puzzle: %w", err)
	}
	return p, nil
}
