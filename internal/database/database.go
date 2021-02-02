package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// var Puzzle *PuzzleTable
// var Records *RecordTable
// var Rooms *RoomTable
var Words *WordTable

// func createWordTable()

func Connect() error {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	Words = &WordTable{DB: db}

	return nil
}
