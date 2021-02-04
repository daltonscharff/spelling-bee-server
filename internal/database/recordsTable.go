package database

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

type RecordsTable struct {
	*sqlx.DB
}

func (t RecordsTable) initTable() error {
	_, err := t.Exec(`CREATE TABLE IF NOT EXISTS records
		(
				id serial NOT NULL,
				word_id integer NOT NULL,
				room_id integer NOT NULL,
				player_name character varying(64) NOT NULL,
				found_at timestamp with time zone NOT NULL,
				CONSTRAINT records_pkey PRIMARY KEY (id),
				CONSTRAINT records_word_id_room_id_key UNIQUE (word_id, room_id),
				CONSTRAINT records_room_id_fkey FOREIGN KEY (room_id)
						REFERENCES rooms (id) MATCH SIMPLE
						ON UPDATE NO ACTION
						ON DELETE NO ACTION,
				CONSTRAINT records_word_id_fkey FOREIGN KEY (word_id)
						REFERENCES words (id) MATCH SIMPLE
						ON UPDATE NO ACTION
						ON DELETE NO ACTION
		);`)
	return err
}

func (t *RecordsTable) Read(id uint64) (Record, error) {
	var r Record
	if err := t.Get(&r, `SELECT * FROM finds WHERE id = $1;`, id); err != nil {
		return Record{}, fmt.Errorf("error getting record: %w", err)
	}
	return r, nil
}

func (t *RecordsTable) ReadAll() ([]Record, error) {
	var rr []Record
	if err := t.Select(&rr, `SELECT * FROM finds;`); err != nil {
		return []Record{}, fmt.Errorf("error getting records: %w", err)
	}
	return rr, nil
}

func (t *RecordsTable) Create(r *Record) error {
	if err := t.Get(r, `INSERT INTO finds (word_id, room_id, player_name, found_at) VALUES ($1, $2, $3, $4) RETURNING *`,
		r.WordID,
		r.RoomID,
		r.PlayerName,
		r.FoundAt); err != nil {
		return fmt.Errorf("error creating record: %w", err)
	}
	return nil
}

func (t *RecordsTable) Update(r *Record) error {
	if err := t.Get(r, `UPDATE finds SET word_id = $1, room_id = $2, player_name = $3, found_at = $4 RETURNING *;`,
		r.WordID,
		r.RoomID,
		r.PlayerName,
		r.FoundAt); err != nil {
		return fmt.Errorf("error updating record: %w", err)
	}
	return nil
}

func (t *RecordsTable) Delete(id uint64) (Record, error) {
	var r Record
	if err := t.Get(r, `DELETE FROM finds WHERE id = $1 RETURNING *;`, id); err != nil {
		return Record{}, fmt.Errorf("error deleting record: %w", err)
	}
	return r, nil
}
