package db

type Find struct {
	ID         int    `json:"id,omitempty"`
	WordID     int    `json:"word_id"`
	RoomID     int    `json:"room_id"`
	PlayerName string `json:"player_name"`
	FoundAt    string `json:"found_at"`
}

type Puzzle struct {
	ID      int      `json:"id,omitempty"`
	Date    string   `json:"date"`
	Letters []string `json:"letters"`
	Center  string   `json:"center"`
}

type Room struct {
	ID   int    `json:"id,omitempty"`
	Code string `json:"code"`
}

type Word struct {
	ID          int    `json:"id,omitempty"`
	Word        string `json:"word"`
	PuzzleID    int    `json:"puzzle_id"`
	PointValue  int    `json:"point_value"`
	Definitions []struct {
		Definition   string `json:"definition"`
		PartOfSpeech string `json:"part_of_speech"`
	} `json:"definitions"`
}
