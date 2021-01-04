package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Room struct {
	ID    uint   `db:"id" json:"id"`
	Code  string `db:"code" json:"code"`
	Score uint   `db:"score" json:"score"`
}

type RoomStore struct {
	*sqlx.DB
}

func (s *RoomStore) Room(id uint64) (Room, error) {
	var r Room
	if err := s.Get(&r, `SELECT * FROM rooms WHERE id = $1;`, id); err != nil {
		return Room{}, fmt.Errorf("error getting room: %w", err)
	}
	return r, nil
}

func (s *RoomStore) Rooms() ([]Room, error) {
	var rr []Room
	if err := s.Select(&rr, `SELECT * FROM rooms;`); err != nil {
		return []Room{}, fmt.Errorf("error getting rooms: %w", err)
	}
	return rr, nil
}

func (s *RoomStore) CreateRoom(r *Room) error {
	if err := s.Get(r, `INSERT INTO rooms (code, score) VALUES ($1, $2) RETURNING *`,
		r.Code,
		r.Score); err != nil {
		return fmt.Errorf("error creating room: %w", err)
	}
	return nil
}

func (s *RoomStore) UpdateRoom(r *Room) error {
	if err := s.Get(r, `UPDATE rooms SET code = $1, score = $2 RETURNING *;`,
		r.Code,
		r.Score); err != nil {
		return fmt.Errorf("error updating room: %w", err)
	}
	return nil
}

func (s *RoomStore) DeleteRoom(id uint64) (Room, error) {
	var r Room
	if err := s.Get(r, `DELETE FROM rooms WHERE id = $1 RETURNING *;`, id); err != nil {
		return Room{}, fmt.Errorf("error deleting room: %w", err)
	}
	return r, nil
}
