package db

type Puzzle struct {
	ID      int    `json:"id,omitempty"`
	Date    string `json:"date"`
	Letters []byte `json:"letters"`
	Center  byte   `json:"center"`
}
