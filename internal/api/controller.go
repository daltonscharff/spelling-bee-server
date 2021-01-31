package api

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/daltonscharff/spelling-bee-server/internal/api/word"
)

type Controller struct {
	// Puzzles *puzzle.Controller
	// Records *record.Controller
	// Rooms   *room.Controller
	Words *word.Controller
}

func CreateController() (*Controller, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Controller{
		// Puzzles: &puzzle.Controller{DB: db},
		// Records: &record.Controller{DB: db},
		// Rooms:   &room.Controller{DB: db},
		Words: &word.Controller{DB: db},
	}, nil
}
