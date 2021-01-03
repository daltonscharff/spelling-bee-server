package router

import "time"

type Find struct {
	ID         uint      `json:"-"`
	Word       Word      `json:"word"`
	PlayerName string    `json:"playerName"`
	FoundAt    time.Time `json:"foundAt"`
}
