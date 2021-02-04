package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Puzzles *puzzlesTable
var Records *recordsTable
var Rooms *roomsTable
var Words *wordsTable

func Connect() error {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	Puzzles = &puzzlesTable{DB: db}
	Records = &recordsTable{DB: db}
	Rooms = &roomsTable{DB: db}
	Words = &wordsTable{DB: db}

	Puzzles.InitTable()
	Rooms.InitTable()
	Words.InitTable()
	Records.InitTable()

	return nil
}
