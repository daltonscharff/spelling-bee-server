package db

type Word struct {
	ID          int    `json:"id,omitempty"`
	Word        string `json:"word"`
	PuzzleID    int    `json:"puzzle_id"`
	Points      int    `json:"points"`
	Definitions []struct {
		Definition   string `json:"definition"`
		PartOfSpeech string `json:"part_of_speech"`
	} `json:"definitions"`
}
