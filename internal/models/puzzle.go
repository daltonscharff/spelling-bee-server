package models

import (
	"time"
)

type Puzzle struct {
	ID       int       `db:"id" json:"id"`
	Date     time.Time `db:"date" json:"date"`
	Letters  []string  `db:"letters" json:"letters"`
	Center   string    `db:"center" json:"center"`
	MaxScore int       `db:"max_score" json:"maxScore"`
}
