package models

import (
	"time"
)

type Find struct {
	ID         int       `db:"id" json:"id"`
	WordID     int       `db:"word_id" json:"wordId"`
	RoomID     int       `db:"room_id" json:"roomId"`
	PlayerName string    `db:"player_name" json:"playerName"`
	FoundAt    time.Time `db:"found_at" json:"foundAt"`
}
