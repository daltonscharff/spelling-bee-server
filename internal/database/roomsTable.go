package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Room struct {
	ID    uint   `db:"id" json:"id"`
	Code  string `db:"code" json:"code"`
	Score uint   `db:"score" json:"score"`
}

type RoomsTable struct {
	*sqlx.DB
}

func (t RoomsTable) initTable() error {
	_, err := t.Exec(`CREATE TABLE IF NOT EXISTS rooms
		(
				id serial NOT NULL,
				code character(8) NOT NULL,
				score integer NOT NULL,
				CONSTRAINT rooms_pkey PRIMARY KEY (id),
				CONSTRAINT rooms_code_key UNIQUE (code)
		);`)
	return err
}

func (t *RoomsTable) Read(id uint64) (Room, error) {
	var r Room
	if err := t.Get(&r, `SELECT * FROM rooms WHERE id = $1;`, id); err != nil {
		return Room{}, fmt.Errorf("error getting room: %w", err)
	}
	return r, nil
}

func (t *RoomsTable) ReadAll() ([]Room, error) {
	var rr []Room
	if err := t.Select(&rr, `SELECT * FROM rooms;`); err != nil {
		return []Room{}, fmt.Errorf("error getting rooms: %w", err)
	}
	return rr, nil
}

func (t *RoomsTable) Create(r *Room) error {
	if err := t.Get(r, `INSERT INTO rooms (code, score) VALUES ($1, $2) RETURNING *`,
		r.Code,
		r.Score); err != nil {
		return fmt.Errorf("error creating room: %w", err)
	}
	return nil
}

func (t *RoomsTable) Update(r *Room) error {
	if err := t.Get(r, `UPDATE rooms SET code = $1, score = $2 RETURNING *;`,
		r.Code,
		r.Score); err != nil {
		return fmt.Errorf("error updating room: %w", err)
	}
	return nil
}

func (t *RoomsTable) Delete(id uint64) (Room, error) {
	var r Room
	if err := t.Get(r, `DELETE FROM rooms WHERE id = $1 RETURNING *;`, id); err != nil {
		return Room{}, fmt.Errorf("error deleting room: %w", err)
	}
	return r, nil
}
