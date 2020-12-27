package models

type Word struct {
	ID          int    `db:"id" json:"id"`
	Word        string `db:"word" json:"word"`
	PuzzleID    int    `db:"puzzle_id" json:"puzzleId"`
	PointValue  int    `db:"point_value" json:"pointValue"`
	Definitions []struct {
		Definition   string `db:"definition" json:"definition"`
		PartOfSpeech string `db:"part_of_speech" json:"partOfSpeech"`
	} `db:"definitions" json:"definitions"`
}
