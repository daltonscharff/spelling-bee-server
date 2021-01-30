package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	Finds   *FindStore
	Puzzles *PuzzleStore
	Rooms   *RoomStore
	Words   *WordStore
}

// type WordStoreInterface interface {
// 	Read(id uint64) (Word{}, error)
// 	ReadAll() ([]Word{}, error)
// 	Create(*Word{}) error
// 	Update(*Word{}) error
// 	Delete(id uint64) (Word{}, error)
// 	DeleteAll() ([]Word{}, error)
// }

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return &Store{
		Finds:   &FindStore{DB: db},
		Puzzles: &PuzzleStore{DB: db},
		Rooms:   &RoomStore{DB: db},
		Words:   &WordStore{DB: db},
	}, nil
}
