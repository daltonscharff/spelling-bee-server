package db

type Puzzle struct {
	ID      int
	Date    string
	Letters []byte
	Center  byte
}

type Find struct {
	ID         int
	WordID     int
	RoomID     int
	PlayerName string
	FoundAt    string
}

type Room struct {
	ID   int
	Code string
}

type Word struct {
	ID          int
	Word        string
	PuzzleID    int
	Points      int
	Definitions []struct {
		Definition   string
		PartOfSpeech string
	}
}
