package room

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Room struct {
	ID    uint   `db:"id" json:"id"`
	Code  string `db:"code" json:"code"`
	Score uint   `db:"score" json:"score"`
}

type Controller struct {
	*sqlx.DB
}

func (c *Controller) Room(id uint64) (Room, error) {
	var r Room
	if err := c.Get(&r, `SELECT * FROM rooms WHERE id = $1;`, id); err != nil {
		return Room{}, fmt.Errorf("error getting room: %w", err)
	}
	return r, nil
}

func (c *Controller) Rooms() ([]Room, error) {
	var rr []Room
	if err := c.Select(&rr, `SELECT * FROM rooms;`); err != nil {
		return []Room{}, fmt.Errorf("error getting rooms: %w", err)
	}
	return rr, nil
}

func (c *Controller) CreateRoom(r *Room) error {
	if err := c.Get(r, `INSERT INTO rooms (code, score) VALUES ($1, $2) RETURNING *`,
		r.Code,
		r.Score); err != nil {
		return fmt.Errorf("error creating room: %w", err)
	}
	return nil
}

func (c *Controller) UpdateRoom(r *Room) error {
	if err := c.Get(r, `UPDATE rooms SET code = $1, score = $2 RETURNING *;`,
		r.Code,
		r.Score); err != nil {
		return fmt.Errorf("error updating room: %w", err)
	}
	return nil
}

func (c *Controller) DeleteRoom(id uint64) (Room, error) {
	var r Room
	if err := c.Get(r, `DELETE FROM rooms WHERE id = $1 RETURNING *;`, id); err != nil {
		return Room{}, fmt.Errorf("error deleting room: %w", err)
	}
	return r, nil
}
