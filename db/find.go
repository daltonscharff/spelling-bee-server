package db

type Find struct {
	ID         int    `json:"id,omitempty"`
	WordID     int    `json:"word_id"`
	RoomID     int    `json:"room_id"`
	PlayerName string `json:"player_name"`
	FoundAt    string `json:"found_at"`
}
