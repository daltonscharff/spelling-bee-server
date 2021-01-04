package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	*FindStore
	*PuzzleStore
	*RoomStore
	*WordStore
}

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return &Store{
		FindStore:   &FindStore{DB: db},
		PuzzleStore: &PuzzleStore{DB: db},
		RoomStore:   &RoomStore{DB: db},
		WordStore:   &WordStore{DB: db},
	}, nil
}
