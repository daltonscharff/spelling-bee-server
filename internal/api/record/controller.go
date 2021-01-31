package record

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Record struct {
	ID         uint `db:"id" json:"id"`
	WordID     uint `db:"word_id" json:"wordId"`
	RoomID     uint `db:"room_id" json:"roomId"`
	PlayerName uint `db:"player_name" json:"playerName"`
	FoundAt    uint `db:"found_at" json:"foundAt"`
}

type Controller struct {
	*sqlx.DB
}

func (c *Controller) Find(id uint64) (Record, error) {
	var f Record
	if err := c.Get(&f, `SELECT * FROM finds WHERE id = $1;`, id); err != nil {
		return Record{}, fmt.Errorf("error getting find: %w", err)
	}
	return f, nil
}

func (c *Controller) Finds() ([]Record, error) {
	var ff []Record
	if err := c.Select(&ff, `SELECT * FROM finds;`); err != nil {
		return []Record{}, fmt.Errorf("error getting finds: %w", err)
	}
	return ff, nil
}

func (c *Controller) CreateFind(f *Record) error {
	if err := c.Get(f, `INSERT INTO finds (word_id, room_id, player_name, found_at) VALUES ($1, $2, $3, $4) RETURNING *`,
		f.WordID,
		f.RoomID,
		f.PlayerName,
		f.FoundAt); err != nil {
		return fmt.Errorf("error creating find: %w", err)
	}
	return nil
}

func (c *Controller) UpdateFind(f *Record) error {
	if err := c.Get(f, `UPDATE finds SET word_id = $1, room_id = $2, player_name = $3, found_at = $4 RETURNING *;`,
		f.WordID,
		f.RoomID,
		f.PlayerName,
		f.FoundAt); err != nil {
		return fmt.Errorf("error updating find: %w", err)
	}
	return nil
}

func (c *Controller) DeleteFind(id uint64) (Record, error) {
	var f Record
	if err := c.Get(f, `DELETE FROM finds WHERE id = $1 RETURNING *;`, id); err != nil {
		return Record{}, fmt.Errorf("error deleting word: %w", err)
	}
	return f, nil
}
