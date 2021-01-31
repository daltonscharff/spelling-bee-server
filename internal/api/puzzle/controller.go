package puzzle

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

type Controller struct {
	*sqlx.DB
}

func (c *Controller) Puzzle(id uint64) (Puzzle, error) {
	var p Puzzle
	if err := c.Get(&p, `SELECT * FROM puzzles WHERE id = $1;`, id); err != nil {
		return Puzzle{}, fmt.Errorf("error getting puzzle: %w", err)
	}
	return p, nil
}

func (c *Controller) Puzzles() ([]Puzzle, error) {
	var pp []Puzzle
	if err := c.Select(&pp, `SELECT * FROM puzzles;`); err != nil {
		return []Puzzle{}, fmt.Errorf("error getting puzzles: %w", err)
	}
	return pp, nil
}

func (c *Controller) CreatePuzzle(p *Puzzle) error {
	if err := c.Get(p, `INSERT INTO puzzles (date, letters, center_letter, max_score) VALUES ($1, $2, $3, $4) RETURNING *`,
		p.Date,
		p.Letters,
		p.CenterLetter,
		p.MaxScore); err != nil {
		return fmt.Errorf("error creating puzzle: %w", err)
	}
	return nil
}

func (c *Controller) UpdatePuzzle(p *Puzzle) error {
	if err := c.Get(p, `UPDATE puzzles SET date = $1, letters = $2, center_letter = $3, max_score = $4 RETURNING *;`,
		p.Date,
		p.Letters,
		p.CenterLetter,
		p.MaxScore); err != nil {
		return fmt.Errorf("error updating puzzle: %w", err)
	}
	return nil
}

func (c *Controller) DeletePuzzle(id uint64) (Puzzle, error) {
	var p Puzzle
	if err := c.Get(p, `DELETE FROM puzzles WHERE id = $1 RETURNING *;`, id); err != nil {
		return Puzzle{}, fmt.Errorf("error deleting puzzle: %w", err)
	}
	return p, nil
}
