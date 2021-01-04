package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Find struct {
	ID         uint `db:"id" json:"id"`
	WordID     uint `db:"word_id" json:"wordId"`
	RoomID     uint `db:"room_id" json:"roomId"`
	PlayerName uint `db:"player_name" json:"playerName"`
	FoundAt    uint `db:"found_at" json:"foundAt"`
}

type FindStore struct {
	*sqlx.DB
}

func (s *FindStore) Find(id uint64) (Find, error) {
	var f Find
	if err := s.Get(&f, `SELECT * FROM finds WHERE id = $1;`, id); err != nil {
		return Find{}, fmt.Errorf("error getting find: %w", err)
	}
	return f, nil
}

func (s *FindStore) Finds() ([]Find, error) {
	var ff []Find
	if err := s.Select(&ff, `SELECT * FROM finds;`); err != nil {
		return []Find{}, fmt.Errorf("error getting finds: %w", err)
	}
	return ff, nil
}

func (s *FindStore) CreateFind(f *Find) error {
	if err := s.Get(f, `INSERT INTO finds (word_id, room_id, player_name, found_at) VALUES ($1, $2, $3, $4) RETURNING *`,
		f.WordID,
		f.RoomID,
		f.PlayerName,
		f.FoundAt); err != nil {
		return fmt.Errorf("error creating find: %w", err)
	}
	return nil
}

func (s *FindStore) UpdateFind(f *Find) error {
	if err := s.Get(f, `UPDATE finds SET word_id = $1, room_id = $2, player_name = $3, found_at = $4 RETURNING *;`,
		f.WordID,
		f.RoomID,
		f.PlayerName,
		f.FoundAt); err != nil {
		return fmt.Errorf("error updating find: %w", err)
	}
	return nil
}

func (s *FindStore) DeleteFind(id uint64) (Find, error) {
	var f Find
	if err := s.Get(f, `DELETE FROM finds WHERE id = $1 RETURNING *;`, id); err != nil {
		return Find{}, fmt.Errorf("error deleting word: %w", err)
	}
	return f, nil
}
